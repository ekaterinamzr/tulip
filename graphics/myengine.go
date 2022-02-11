package graphics

import (
	"image/color"
	"tulip/mymath"
	"tulip/scene"
)

type MyGrEngine struct {
	cnv Canvas

	projection mymath.Matrix4x4
	viewport   mymath.Matrix4x4

	zBuf  []float64
	zBack float64

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

	return engine
}

func (engine *MyGrEngine) initZBuf() {
	for i := range engine.zBuf {
		engine.zBuf[i] = engine.zBack
	}
}

// Rendering scene
func (engine MyGrEngine) RenderScene(scn *scene.Scene) {

	// setting up the engine
	// setting shader
	engine.shader.light = scn.LightSource

	// rendering
	engine.cnv.fill(scn.Background)
	engine.initZBuf()

	for i := range scn.Objects {
		engine.renderModel(scn.Objects[i])
	}
}

func (engine MyGrEngine) renderModel(m scene.PolygonialModel) {
	m.IterateOverPolygons(engine.processTriangle)
}

// creating a tringle -> then rendering
func (engine MyGrEngine) processTriangle(v0, v1, v2 scene.Vertex, clr color.NRGBA) {
	engine.shader.vs(&v0)
	engine.shader.vs(&v1)
	engine.shader.vs(&v2)

	t := makeTriangle(v0, v1, v2, clr)

	engine.renderTriangle(t)
}

// applying projection, viewport -> then rasterizing
func (engine MyGrEngine) renderTriangle(t triangle) {
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
	p0, p1, p2 := t.v0.Point.Vec3, t.v1.Point.Vec3, t.v2.Point.Vec3
	i0, i1, i2 := t.v0.Intensity, t.v1.Intensity, t.v2.Intensity

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

		if a.X > b.X {
			a, b = b, a
			ia, ib = ib, ia
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   scene.Vertex
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
					pixelX, pixelY, pixelClr := engine.shader.ps(p, t.clr)
					pixelY = engine.cnv.height() - pixelY
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
				p   scene.Vertex
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
					pixelX, pixelY, pixelClr := engine.shader.ps(p, t.clr)
					pixelY = engine.cnv.height() - pixelY
					engine.cnv.setPixel(pixelX, pixelY, pixelClr)
				}
			}
		}
	}
}
