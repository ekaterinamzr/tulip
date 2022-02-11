package graphics

import (
	// "fmt"
	"image"
	"image/color"

	// "math"
	"tulip/mymath"
	"tulip/scene"
)

type MyGraphicsEngine struct {
	cnv Canvas

	zbuf  []float64
	sbuf  []float64
	zback float64

	Projection mymath.Matrix4x4
	viewport   mymath.Matrix4x4
	view       mymath.Matrix4x4

	world2Screen mymath.Matrix4x4
	sun2Screen   mymath.Matrix4x4

	worldS2sunS mymath.Matrix4x4

	scn *scene.Scene
}

func NewEngine(h, w int) *MyGraphicsEngine {
	engine := new(MyGraphicsEngine)

	engine.cnv = MakeImageCanvas(h, w)
	engine.zback = -1000
	engine.zbuf = make([]float64, h*w)
	engine.sbuf = make([]float64, h*w)

	//engine.setViewport(float64(w)/8.0, float64(h)/8.0, float64(w)*3.0/4.0, float64(h)*3.0/4.0)
	engine.viewport = engine.makeViewport(0, 0, float64(w), float64(h))
	//engine.setProjection()

	return engine
}

func (engine MyGraphicsEngine) makeViewport(x, y, w, h float64) mymath.Matrix4x4 {
	viewport := mymath.MakeIdentityM()
	depth := 500.0

	viewport[0][3] = x + w/2.0
	viewport[1][3] = y + h/2.0
	viewport[2][3] = depth / 2.0

	viewport[0][0] = w / 2.0
	viewport[1][1] = h / 2.0
	viewport[2][2] = depth / 2.0

	return viewport
}

func makeProjection(eye, center mymath.Vec3d) (mymath.Matrix4x4, bool) {
	var proj mymath.Matrix4x4

	dist := mymath.Vec3dDiff(eye, center).Length()

	if dist == 0 {
		return proj, false
	}

	proj = mymath.MakeIdentityM()
	proj[3][2] = -1.0 / dist

	return proj, true
}

func makeProjectionK(k float64) mymath.Matrix4x4 {
	proj := mymath.MakeIdentityM()
	proj[3][2] = k

	return proj
}

func (engine *MyGraphicsEngine) initZBuffer() {
	for i := range engine.zbuf {
		engine.zbuf[i] = engine.zback
		engine.sbuf[i] = engine.zback
	}
}

// Матрица view переводит координаты в новый базис
func (engine MyGraphicsEngine) lookAt(eye, center, up mymath.Vec3d) mymath.Matrix4x4 {
	var view mymath.Matrix4x4

	z := mymath.Vec3dDiff(eye, center)
	z.Normalize()

	x := up.CrossProduct(z)
	x.Normalize()

	y := z.CrossProduct(x)
	y.Normalize()

	minv := mymath.MakeIdentityM()
	tr := mymath.MakeIdentityM()

	// for i := 0; i < 3; i++ {
	// 	minv[0][1] = x[i]
	// }

	minv[0][0] = x.X
	minv[1][0] = y.X
	minv[2][0] = z.X
	tr[0][3] = -center.X

	minv[0][1] = x.Y
	minv[1][1] = y.Y
	minv[2][1] = z.Y
	tr[1][3] = -center.Y

	minv[0][2] = x.Z
	minv[1][2] = y.Z
	minv[2][2] = z.Z
	tr[2][3] = -center.Z

	view = mymath.MulMatrices(minv, tr)

	return view
}

func (engine MyGraphicsEngine) RenderScene(scn *scene.Scene, onlyShadowMap, shadows bool) {
	camPos := scn.Camera.Pos
	center := scn.Camera.Center
	engine.scn = scn

	lightPos := scn.LightSource.Pos

	sunView := engine.lookAt(lightPos, center, mymath.MakeVec3d(0, 1, 0))
	sunProj := makeProjectionK(0)

	engine.view = engine.lookAt(camPos, center, mymath.MakeVec3d(0, 1, 0))

	ok := true
	engine.Projection, ok = makeProjection(camPos, center)
	// engine.Projection = makeProjectionK(0)

	if !ok {
		// fmt.Println("Could not make projection")
		return
	}

	engine.world2Screen = mymath.MulMatrices(engine.viewport, engine.Projection)
	engine.world2Screen = mymath.MulMatrices(engine.view, engine.world2Screen)

	engine.sun2Screen = mymath.MulMatrices(engine.viewport, sunProj)
	engine.sun2Screen = mymath.MulMatrices(sunView, engine.sun2Screen)

	mInv, _ := engine.world2Screen.Inverse()

	engine.worldS2sunS = mymath.MulMatrices(engine.sun2Screen, mInv)

	// engine.cnv.fill(scn.Background)

	engine.initZBuffer()

	//engine.RenderModel(scn.Ground)

	if onlyShadowMap || shadows {
		engine.cnv.fill(color.NRGBA{100, 100, 200, 255})
		for i := 1; i < len(scn.Objects); i++ {
			//fmt.Println("Rendering object ", i)
			engine.RenderModelShadow(scn.Objects[i])
		}
	}

	if !onlyShadowMap {
		engine.cnv.fill(scn.Background)
		for i := range scn.Objects {
			//fmt.Println("Rendering object ", i)
			engine.RenderModel(scn.Objects[i])
		}
	}

}

func (engine MyGraphicsEngine) inShadow(p mymath.Vec3d) bool {
	pSun := mymath.MulVectorMatrix(p, engine.worldS2sunS)
	pSun.DivW()
	idx := int(p.X) + int(p.Y)*engine.cnv.width()

	if engine.sbuf[idx] == engine.zback {
		return false
	}

	if engine.sbuf[idx] > pSun.Z {
		return true
	}

	return false
}

func (engine MyGraphicsEngine) RenderModel(obj scene.PolygonialModel) {
	obj.IterateOverPolygons(engine.RenderPolygon)
}

func (engine MyGraphicsEngine) RenderPolygon(v0, v1, v2 scene.Vertex, clr color.NRGBA) {
	//fmt.Println("Rendering polygon ", v0, v1, v2)
	var (
		screenCoords [3]mymath.Vec3d
		intensity    [3]float64
	)
	// world2Screen := mymath.MulMatrices(engine.viewport, engine.Projection)
	// world2Screen = mymath.MulMatrices(engine.view, world2Screen)

	screenCoords[0] = mymath.MulVectorMatrix(v0.Point, engine.world2Screen)
	screenCoords[1] = mymath.MulVectorMatrix(v1.Point, engine.world2Screen)
	screenCoords[2] = mymath.MulVectorMatrix(v2.Point, engine.world2Screen)

	// fmt.Println("Screen polygon ", screenCoords[0], screenCoords[1], screenCoords[2])
	// fmt.Println()

	screenCoords[0].DivW()
	screenCoords[1].DivW()
	screenCoords[2].DivW()

	// fmt.Println("Screen polygon ", screenCoords[0], screenCoords[1], screenCoords[2])
	// fmt.Println()

	intensity[0] = scene.VectorIntensity(v0.Normal, engine.scn.LightSource)
	intensity[1] = scene.VectorIntensity(v1.Normal, engine.scn.LightSource)
	intensity[2] = scene.VectorIntensity(v2.Normal, engine.scn.LightSource)

	engine.rasterizePolygon(screenCoords[0], screenCoords[1], screenCoords[2],
		intensity[0], intensity[1], intensity[2], clr, engine.zbuf, true)

	//engine.rasterizeWire(screenCoords[0], screenCoords[1], screenCoords[2], color.NRGBA{0, 0, 0, 255})
}

func (engine MyGraphicsEngine) RenderModelShadow(obj scene.PolygonialModel) {
	obj.IterateOverPolygons(engine.RenderPolygonShadow)
}

func (engine MyGraphicsEngine) RenderPolygonShadow(v0, v1, v2 scene.Vertex, clr color.NRGBA) {
	//fmt.Println("Rendering polygon ", v0, v1, v2)
	var (
		screenCoords [3]mymath.Vec3d
		intensity    [3]float64
	)

	screenCoords[0] = mymath.MulVectorMatrix(v0.Point, engine.sun2Screen)
	screenCoords[1] = mymath.MulVectorMatrix(v1.Point, engine.sun2Screen)
	screenCoords[2] = mymath.MulVectorMatrix(v2.Point, engine.sun2Screen)

	// fmt.Println("Screen polygon ", screenCoords[0], screenCoords[1], screenCoords[2])
	// fmt.Println()

	screenCoords[0].DivW()
	screenCoords[1].DivW()
	screenCoords[2].DivW()

	// fmt.Println("Screen polygon ", screenCoords[0], screenCoords[1], screenCoords[2])
	// fmt.Println()

	intensity[0] = scene.VectorIntensity(v0.Normal, engine.scn.LightSource)
	intensity[1] = scene.VectorIntensity(v1.Normal, engine.scn.LightSource)
	intensity[2] = scene.VectorIntensity(v2.Normal, engine.scn.LightSource)

	engine.rasterizePolygon(screenCoords[0], screenCoords[1], screenCoords[2],
		intensity[0], intensity[1], intensity[2], color.NRGBA{255, 255, 255, 255}, engine.sbuf, false)

	//engine.rasterizeWire(screenCoords[0], screenCoords[1], screenCoords[2], color.NRGBA{0, 0, 0, 255})
}

func point2pixel(h int, v mymath.Vec3d) (int, int) {
	return int(v.X), h - int(v.Y)
}

func (engine MyGraphicsEngine) rasterizeWire(p0, p1, p2 mymath.Vec3d, clr color.NRGBA) {
	h := engine.cnv.height()
	x0, y0 := point2pixel(h, p0)
	x1, y1 := point2pixel(h, p1)
	x2, y2 := point2pixel(h, p2)

	engine.cnv.drawLine(x0, y0, x1, y1, clr)
	engine.cnv.drawLine(x1, y1, x2, y2, clr)
	engine.cnv.drawLine(x0, y0, x2, y2, clr)
}

func (engine MyGraphicsEngine) rasterizePolygon(p0, p1, p2 mymath.Vec3d, i0, i1, i2 float64, clr color.NRGBA, buf []float64, draw bool) {
	if p0.Y > p1.Y {
		p0, p1 = p1, p0
		i0, i1 = i1, i0
	}
	if p0.Y > p2.Y {
		p0, p2 = p2, p0
		i0, i2 = i2, i0
	}
	if p1.Y > p2.Y {
		p1, p2 = p2, p1
		i1, i2 = i2, i1
	}

	dyTotal := p2.Y - p0.Y

	for y := p0.Y; y <= p1.Y; y++ {
		dySegment := p1.Y - p0.Y + 1
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p0.Y) / dySegment)

		var a, b mymath.Vec3d

		a = mymath.Vec3dDiff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mymath.Vec3dDiff(p1, p0)
		b.Mul(beta)
		b.Add(p0)

		var ia, ib float64

		ia = i0 + (i2-i0)*alpha
		ib = i0 + (i1-i0)*beta

		if a.X > b.X {
			a, b = b, a
			ia, ib = ib, ia
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vec3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			intensity := ia + (ib-ia)*phi

			idx := int(p.X) + int(p.Y)*engine.cnv.width()
			if x >= 0 && y >= 0 && x < float64(engine.cnv.width()) && y < float64(engine.cnv.height()) {
				if buf[idx] < p.Z {
					buf[idx] = p.Z

					if draw {
						pixelX, pixelY := point2pixel(engine.cnv.height(), p)
						if engine.inShadow(p) {
							engine.cnv.setPixel(pixelX, pixelY, color.Black)
						} else {
							engine.cnv.setPixel(pixelX, pixelY, scene.Lightness(clr, intensity))
						}

					}

				}
			}
		}
	}

	for y := p1.Y; y <= p2.Y; y++ {
		dySegment := p2.Y - p1.Y + 1
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p1.Y) / dySegment)

		var a, b mymath.Vec3d

		a = mymath.Vec3dDiff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mymath.Vec3dDiff(p2, p1)
		b.Mul(beta)
		b.Add(p1)

		var ia, ib float64

		ia = i0 + (i2-i0)*alpha
		ib = i0 + (i1-i0)*beta

		if a.X > b.X {
			a, b = b, a
			ia, ib = ib, ia
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vec3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			intensity := ia + (ib-ia)*phi

			idx := int(p.X) + int(p.Y)*engine.cnv.width()
			if x >= 0 && y >= 0 && x < float64(engine.cnv.width()) && y < float64(engine.cnv.height()) {
				if buf[idx] < p.Z {
					buf[idx] = p.Z

					if draw {
						pixelX, pixelY := point2pixel(engine.cnv.height(), p)
						if engine.inShadow(p) {
							engine.cnv.setPixel(pixelX, pixelY, color.Black)
						} else {
							engine.cnv.setPixel(pixelX, pixelY, scene.Lightness(clr, intensity))
						}

					} else {
						pixelX, pixelY := point2pixel(engine.cnv.height(), p)
						engine.cnv.setPixel(pixelX, pixelY, scene.Lightness(color.NRGBA{0, 0, 0, 255}, intensity))
					}
				}
			}
		}
	}
}

func (engine MyGraphicsEngine) Image() image.Image {
	return engine.cnv.Image()
}
