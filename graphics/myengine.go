package graphics

import (
	"fmt"
	"image/color"
	"math"
	"tulip/mymath"
	"tulip/scene"
)

type MyGrEngine struct {
	cnv Canvas

	zBuf  [][]float64
	zBack float64
	sBuf  [][]float64

	//sBuf []float64

	pst psTransformer
	// shader gouraudShader
	shader shaderInterface
}

// Creating new engine
func NewMyGrEngine(cnv Canvas) *MyGrEngine {
	engine := new(MyGrEngine)

	// setting canvas
	engine.cnv = cnv

	// setting z-buffer
	engine.zBack = 10000.0
	engine.zBuf = make([][]float64, engine.cnv.height())
	for i := range engine.zBuf {
		engine.zBuf[i] = make([]float64, engine.cnv.width())
	}
	engine.sBuf = make([][]float64, engine.cnv.height())
	for i := range engine.sBuf {
		engine.sBuf[i] = make([]float64, engine.cnv.width())
	}

	engine.pst = makePST(cnv.width(), cnv.height())

	return engine
}

func (engine *MyGrEngine) initZBuf() {
	for i := range engine.zBuf {
		for j := range engine.zBuf[i] {
			engine.zBuf[i][j] = engine.zBack
		}
	}
}

// Rendering scene
func (engine MyGrEngine) RenderScene(scn *scene.Scene) {
	// shader settings
	// projection matrix
	proj := mymath.MakeFovProjectionM(90.0, float64(engine.cnv.height())/float64(engine.cnv.width()), 1.0, 100.0)

	// view matrix
	vUp := mymath.MakeVec3(0, 1, 0)
	t := mymath.MakeVec4(0, 0, 1, 0)
	scn.Camera.VLookDir = mymath.MulVecMat(t, mymath.MakeRotationYM(scn.Camera.FYaw))
	scn.Camera.VTarget = mymath.Vec3Sum(scn.Camera.VCamera, scn.Camera.VLookDir.Vec3)
	scn.Camera.VForward = mymath.Vec3Mul(scn.Camera.VLookDir.Vec3, 1.0)
	mCamera := mymath.MakePointAtM(scn.Camera.VCamera, scn.Camera.VTarget, vUp)
	view := mymath.InverseTranslationM(mCamera)

	// shader
	shadow := makeShadowMap(scn.LightSource, proj)
	// engine.shader = shadow

	// rendering
	// engine.cnv.fill(color.Black)
	// engine.initZBuf()

	// for i := range scn.Objects {
	// 	engine.renderModel(scn.Objects[i])
	// }

	// for i := range engine.zBuf {
	// 	for j := range engine.zBuf[i] {
	// 		engine.sBuf[i][j] = engine.zBuf[i][j]
	// 	}
	// }

	lightViewProj := shadow.ViewProj
	gouraud := makeGouraudShader(view, proj, lightViewProj, engine.sBuf, scn.LightSource, engine.pst)
	engine.shader = gouraud

	// rendering
	engine.cnv.fill(scn.Background)
	engine.initZBuf()

	for i := range scn.Objects {
		engine.renderModel(scn.Objects[i])
	}
}

// extracting vertices and indices in a slice form
// TODO: change the model structure
func (engine MyGrEngine) renderModel(m scene.PolygonialModel) {
	vertices, indices := m.GetVertices()
	engine.processVertices(vertices, indices)
}

// applying vertex shader, world -> viewProj
// scene.Vertex -> graphics.Vertex struct
func (engine MyGrEngine) processVertices(vertices []scene.Vertex, indices []int) {
	processed := make([]Vertex, len(vertices))
	for i := range vertices {
		processed[i] = engine.shader.vs(vertices[i])
	}
	engine.assembleTriangles(processed, indices)
}

// extracting triangles out of the slice
func (engine MyGrEngine) assembleTriangles(processed []Vertex, indices []int) {
	end := len(indices) / 3

	fmt.Println(len(processed))
	for i := 0; i < end; i++ {
		v0 := processed[indices[i*3]]
		v1 := processed[indices[i*3+1]]
		v2 := processed[indices[i*3+2]]

		engine.processTriangle(v0, v1, v2)
	}
}

// creating a tringle -> then rendering
func (engine MyGrEngine) processTriangle(v0, v1, v2 Vertex) {
	t := makeTriangle(v0, v1, v2)

	engine.renderTriangle(t)
}

// perspective division, viewport -> then rasterizing
func (engine MyGrEngine) renderTriangle(t triangle) {
	engine.pst.transform(&t.v0)
	engine.pst.transform(&t.v1)
	engine.pst.transform(&t.v2)

	engine.rasterizeTriangle(t, engine.zBuf)
	// engine.rasterizeWire(t)
	// engine.rasterizeNormals(t)
}

// z-buffer algorithm
// TODO: refactoring
func (engine MyGrEngine) rasterizeTriangle(t triangle, buf [][]float64) {
	clr := t.v0.clr
	p0, p1, p2 := t.v0.Point, t.v1.Point, t.v2.Point
	i0, i1, i2 := t.v0.Intensity, t.v1.Intensity, t.v2.Intensity
	w0, w1, w2 := t.v0.worldPoint, t.v1.worldPoint, t.v2.worldPoint

	print()
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
		dySegment := p1.Y - p0.Y //+ 1
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p0.Y) / dySegment)

		var a, b mymath.Vec4

		a = mymath.Vec4Diff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mymath.Vec4Diff(p1, p0)
		b.Mul(beta)
		b.Add(p0)

		var ia, ib float64
		ia = i0 + (i2-i0)*alpha
		ib = i0 + (i1-i0)*beta

		var wa, wb mymath.Vec4
		wa = mymath.Vec4Diff(w2, w0)
		wa.Mul(alpha)
		wa.Add(w0)

		wb = mymath.Vec4Diff(w1, w0)
		wb.Mul(beta)
		wb.Add(w0)

		if a.X > b.X {
			a, b = b, a
			ia, ib = ib, ia
			wa, wb = wb, wa
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   Vertex
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Point.Z = a.Z + (b.Z-a.Z)*phi
			p.Point.W = a.W + (b.W-a.W)*phi

			p.Point.X = x
			p.Point.Y = y

			p.Intensity = ia + (ib-ia)*phi

			p.worldPoint.X = wa.X + (wb.X-wa.X)*phi
			p.worldPoint.Y = wa.Y + (wb.Y-wa.Y)*phi
			p.worldPoint.Z = wa.Z + (wb.Z-wa.Z)*phi

			// w := 1.0 / p.Point.W
			// p.worldPoint.Mul(w)

			//idx := int((math.Round)(p.Point.X)) + int((math.Round)(p.Point.Y))*engine.cnv.width()
			px := int(math.Round(p.Point.X))
			py := int(math.Round(p.Point.Y))
			if px >= 0 && py >= 0 && px < engine.cnv.width() && py < engine.cnv.height() {
				if p.Point.Z < buf[px][py] {
					buf[px][py] = p.Point.Z
					_, _, pixelClr := engine.shader.ps(p, clr)
					engine.cnv.setPixel(px, py, pixelClr)
				}
			}
		}
	}

	for y := p1.Y; y <= p2.Y; y++ {
		dySegment := p2.Y - p1.Y //+ 1
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p1.Y) / dySegment)

		var a, b mymath.Vec4

		a = mymath.Vec4Diff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mymath.Vec4Diff(p2, p1)
		b.Mul(beta)
		b.Add(p1)

		var ia, ib float64
		ia = i0 + (i2-i0)*alpha
		ib = i1 + (i2-i1)*beta

		var wa, wb mymath.Vec4
		wa = mymath.Vec4Diff(w2, w0)
		wa.Mul(alpha)
		wa.Add(w0)

		wb = mymath.Vec4Diff(w2, w1)
		wb.Mul(beta)
		wb.Add(w1)

		if a.X > b.X {
			a, b = b, a
			ia, ib = ib, ia
			wa, wb = wb, wa
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   Vertex
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Point.Z = a.Z + (b.Z-a.Z)*phi
			p.Point.W = a.W + (b.W-a.W)*phi
			p.Point.X = x
			p.Point.Y = y

			p.Intensity = ia + (ib-ia)*phi

			p.worldPoint.X = wa.X + (wb.X-wa.X)*phi
			p.worldPoint.Y = wa.Y + (wb.Y-wa.Y)*phi
			p.worldPoint.Z = wa.Z + (wb.Z-wa.Z)*phi

			// w := 1.0 / p.Point.W
			// p.worldPoint.Mul(w)

			px := int(math.Round(p.Point.X))
			py := int(math.Round(p.Point.Y))
			if px >= 0 && py >= 0 && px < engine.cnv.width() && py < engine.cnv.height() {
				if p.Point.Z < buf[px][py] {
					buf[px][py] = p.Point.Z
					_, _, pixelClr := engine.shader.ps(p, clr)
					engine.cnv.setPixel(px, py, pixelClr)
				}
			}
		}
	}
}

// func makeProjection(w, h, n, f float64) mymath.Matrix4x4 {
// 	var proj mymath.Matrix4x4

// 	proj[0][0] = 2.0 * n / w
// 	proj[1][1] = 2.0 * n / h
// 	proj[2][2] = f / (f - n)
// 	proj[3][2] = -n * f / (f - n)
// 	proj[2][3] = 1.0

// 	return proj
// }

func (engine MyGrEngine) rasterizeWire(t triangle) {
	h := engine.cnv.height()
	x0, y0 := point2pixel(h, t.v0.Point.Vec3)
	x1, y1 := point2pixel(h, t.v1.Point.Vec3)
	x2, y2 := point2pixel(h, t.v2.Point.Vec3)

	engine.cnv.drawLine(x0, y0, x1, y1, color.Black)
	engine.cnv.drawLine(x1, y1, x2, y2, color.Black)
	engine.cnv.drawLine(x0, y0, x2, y2, color.Black)
}

func (engine MyGrEngine) rasterizeNormals(t triangle) {
	h := engine.cnv.height()

	end0 := mymath.Vec3Sum(t.v0.Point.Vec3, t.v0.Normal.Vec3)
	end0.Scale(t.v0.Point.Vec3, 10)

	end1 := mymath.Vec3Sum(t.v1.Point.Vec3, t.v1.Normal.Vec3)
	end1.Scale(t.v1.Point.Vec3, 10)

	end2 := mymath.Vec3Sum(t.v2.Point.Vec3, t.v2.Normal.Vec3)
	end2.Scale(t.v2.Point.Vec3, 10)

	x00, y00 := point2pixel(h, end0)
	x10, y10 := point2pixel(h, end1)
	x20, y20 := point2pixel(h, end2)

	x01, y01 := point2pixel(h, t.v0.Point.Vec3)
	x11, y11 := point2pixel(h, t.v1.Point.Vec3)
	x21, y21 := point2pixel(h, t.v2.Point.Vec3)

	engine.cnv.drawLine(x00, y00, x01, y01, color.Black)
	engine.cnv.drawLine(x10, y10, x11, y11, color.Black)
	engine.cnv.drawLine(x20, y20, x21, y21, color.Black)
}

func point2pixel(h int, v mymath.Vec3) (int, int) {
	return int(v.X), h - int(v.Y)
}
