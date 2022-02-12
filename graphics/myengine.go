package graphics

import (
	"image/color"
	"math"
	"tulip/mymath"
	"tulip/scene"
)

type MyGrEngine struct {
	cnv Canvas

	projection mymath.Matrix4x4
	viewport   mymath.Matrix4x4

	zBuf  []float64
	zBack float64

	pst    psTransformer
	shader gouraudShader
}

// Creating new engine
func NewMyGrEngine(cnv Canvas) *MyGrEngine {
	engine := new(MyGrEngine)

	// setting canvas
	engine.cnv = cnv

	// setting z-buffer
	engine.zBack = 10000.0
	engine.zBuf = make([]float64, engine.cnv.height()*engine.cnv.width())

	engine.pst = makePST(cnv.width(), cnv.height())

	return engine
}

func (engine *MyGrEngine) initZBuf() {
	for i := range engine.zBuf {
		engine.zBuf[i] = engine.zBack
	}
}

// Rendering scene
func (engine MyGrEngine) RenderScene(scn *scene.Scene) {
	// proj := makeProjection(2.0, 2.0, 1.0, 10.0)
	proj := makeFovProjection(scn.Camera.Hfov, scn.Camera.Aspect_ratio, 1.0, 100.0)
	view := makeTranslation(-scn.Camera.Pos.X, -scn.Camera.Pos.Y, -scn.Camera.Pos.Z)
	// setting up the engine
	// setting shader
	engine.shader.makeShader(proj, view, scn.LightSource)
	//engine.shader.light = scn.LightSource
	//engine.shader.projection = makeProjection(2.0, 2.0, 1.0, 10.0)
	//engine.shader.cam = scn.Camera

	// rendering
	engine.cnv.fill(scn.Background)
	engine.initZBuf()

	for i := range scn.Objects {
		engine.renderModel(scn.Objects[i])
	}
}

func (engine MyGrEngine) renderModel(m scene.PolygonialModel) {
	//m.IterateOverPolygons(engine.processTriangle)
	vertices, indices := m.GetVertices()
	engine.processVertices(vertices, indices)
}

func (engine MyGrEngine) processVertices(vertices []scene.Vertex, indices []int) {
	processed := make([]Vertex, len(vertices))
	for i := range vertices {
		processed[i] = engine.shader.vs(vertices[i])
	}
	engine.assembleTriangles(processed, indices)
}

func (engine MyGrEngine) assembleTriangles(processed []Vertex, indices []int) {
	end := len(indices) / 3
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

func makeProjection(w, h, n, f float64) mymath.Matrix4x4 {
	var proj mymath.Matrix4x4

	proj[0][0] = 2.0 * n / w
	proj[1][1] = 2.0 * n / h
	proj[2][2] = f / (f - n)
	proj[3][2] = -n * f / (f - n)
	proj[2][3] = 1.0

	return proj
}

func makeFovProjection(fov, ar, n, f float64) mymath.Matrix4x4 {
	var proj mymath.Matrix4x4

	fovRad := fov * math.Pi / 180.0
	w := 1.0 / math.Tan(fovRad/2.0)
	h := w * ar

	proj[0][0] = w
	proj[1][1] = h
	proj[2][2] = f / (f - n)
	proj[3][2] = -n * f / (f - n)
	proj[2][3] = 1.0

	return proj
}

func makeTranslation(x, y, z float64) mymath.Matrix4x4 {
	tr := mymath.MakeIdentityM()
	tr[3][0] = x
	tr[3][1] = y
	tr[3][2] = z

	return tr
}

// applying projection, viewport -> then rasterizing
func (engine MyGrEngine) renderTriangle(t triangle) {
	engine.pst.transform(&t.v0)
	engine.pst.transform(&t.v1)
	engine.pst.transform(&t.v2)

	// fmt.Println(t.v0.Point)

	engine.rasterizeTriangle(t, engine.zBuf)
	// engine.rasterizeWire(t)
	// engine.rasterizeNormals(t)
}

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

func (engine MyGrEngine) rasterizeTriangle(t triangle, buf []float64) {
	// clr := color.NRGBA{134, 29, 250, 255}
	clr := t.v0.clr
	p0, p1, p2 := t.v0.Point.Vec3, t.v1.Point.Vec3, t.v2.Point.Vec3
	i0, i1, i2 := t.v0.Intensity, t.v1.Intensity, t.v2.Intensity
	// c0, c1, c2 := t.v0.clr, t.v1.clr, t.v2.clr

	// c0 := mymath.MakeVec3(float64(t.v0.clr.R), float64(t.v0.clr.G), float64(t.v0.clr.B))
	// c1 := mymath.MakeVec3(float64(t.v1.clr.R), float64(t.v1.clr.G), float64(t.v1.clr.B))
	// c2 := mymath.MakeVec3(float64(t.v2.clr.R), float64(t.v2.clr.G), float64(t.v2.clr.B))

	print()
	if p0.Y > p1.Y {
		p0, p1 = p1, p0
		i0, i1 = i1, i0
		// c0, c1 = c1, c0
	}
	if p0.Y > p2.Y {
		p0, p2 = p2, p0
		i0, i2 = i2, i0
		// c0, c2 = c2, c0
	}
	if p1.Y > p2.Y {
		p1, p2 = p2, p1
		i1, i2 = i2, i1
		// c1, c2 = c2, c1
	}

	dyTotal := p2.Y - p0.Y

	for y := p0.Y; y <= p1.Y; y++ {
		dySegment := p1.Y - p0.Y + 1
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p0.Y) / dySegment)

		var a, b mymath.Vec3

		a = mymath.Vec3Diff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mymath.Vec3Diff(p1, p0)
		b.Mul(beta)
		b.Add(p0)

		var ia, ib float64
		ia = i0 + (i2-i0)*alpha
		ib = i0 + (i1-i0)*beta
		// var ca, cb mymath.Vec3

		// ca = mymath.Vec3Diff(c2, c0)
		// ca.Mul(alpha)
		// ca.Add(p0)

		// cb = mymath.Vec3Diff(c1, c0)
		// cb.Mul(beta)
		// cb.Add(c0)

		if a.X > b.X {
			a, b = b, a
			ia, ib = ib, ia
			// ca, cb = cb, ca
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

			p.Point.X = x
			p.Point.Y = y

			// cp := mymath.Vec3Mul(mymath.Vec3Diff(cb, ca), phi)
			// cp.Add()
			p.Intensity = ia + (ib-ia)*phi

			idx := int(p.Point.X) + int(p.Point.Y)*engine.cnv.width()
			if x >= 0 && y >= 0 && x < float64(engine.cnv.width()) && y < float64(engine.cnv.height()) {
				if p.Point.Z < buf[idx] {
					buf[idx] = p.Point.Z
					pixelX, pixelY, pixelClr := engine.shader.ps(p, clr)
					// pixelY = engine.cnv.height() - pixelY
					engine.cnv.setPixel(pixelX, pixelY, pixelClr)
				}
			}
		}
	}

	for y := p1.Y; y <= p2.Y; y++ {
		dySegment := p2.Y - p1.Y + 1
		alpha := float64((y - p0.Y) / dyTotal)
		beta := float64((y - p1.Y) / dySegment)

		var a, b mymath.Vec3

		a = mymath.Vec3Diff(p2, p0)
		a.Mul(alpha)
		a.Add(p0)

		b = mymath.Vec3Diff(p2, p1)
		b.Mul(beta)
		b.Add(p1)

		var ia, ib float64
		ia = i0 + (i2-i0)*alpha
		ib = i1 + (i2-i1)*beta

		if a.X > b.X {
			a, b = b, a
			ia, ib = ib, ia
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

			p.Point.X = x
			p.Point.Y = y

			p.Intensity = ia + (ib-ia)*phi

			idx := int(p.Point.X) + int(p.Point.Y)*engine.cnv.width()
			if x >= 0 && y >= 0 && x < float64(engine.cnv.width()) && y < float64(engine.cnv.height()) {
				if p.Point.Z < buf[idx] {
					buf[idx] = p.Point.Z
					pixelX, pixelY, pixelClr := engine.shader.ps(p, clr)
					// pixelY = engine.cnv.height() - pixelY
					engine.cnv.setPixel(pixelX, pixelY, pixelClr)
				}
			}
		}
	}
}
