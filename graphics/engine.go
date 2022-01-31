package graphics

import (
	"image"
	"image/color"
	"tulip/mymath"
	"tulip/object"
)

type GraphicsEngine interface {
	RenderVertex(cnv Canvas, vertex object.Vertex)
	RenderPolygon(cnv Canvas, v1, v2, v3 object.Vertex, clr color.NRGBA)
	RenderModel(cnv Canvas, object object.Model)

	RenderScene(cnv Canvas, scn object.Scene)

	Image() image.Image
}

type ZBufferGraphicsEngine struct {
	Cnv   Canvas
	zbuf  [][]float64
	light object.Light
}

func (engine *ZBufferGraphicsEngine) initBuf(h, w int, value float64) {
	engine.zbuf = make([][]float64, h)
	for i := range engine.zbuf {
		engine.zbuf[i] = make([]float64, w)
		for j := range engine.zbuf[i] {
			engine.zbuf[i][j] = value
		}
	}
}

func projection(h int, v mymath.Vector3d) (int, int) {
	return int(v.X), h - int(v.Y)
}

func (engine ZBufferGraphicsEngine) RenderPolygon(v1, v2, v3 object.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv

	if v1.Point.Y > v2.Point.Y {
		v1, v2 = v2, v1
	}
	if v1.Point.Y > v3.Point.Y {
		v1, v3 = v3, v1
	}
	if v2.Point.Y > v3.Point.Y {
		v2, v3 = v3, v2
	}

	total_dy := v3.Point.Y - v1.Point.Y

	light := engine.light

	v1_intensity := object.VertexIntensity(v1, light)
	v2_intensity := object.VertexIntensity(v2, light)
	v3_intensity := object.VertexIntensity(v3, light)

	for y := v1.Point.Y; y <= v2.Point.Y; y++ {
		segment_dy := v2.Point.Y - v1.Point.Y + 1
		alpha := float64((y - v1.Point.Y) / total_dy)
		beta := float64((y - v1.Point.Y) / segment_dy)

		var a, b mymath.Vector3d

		a = mymath.Vector3dDiff(v3.Point, v1.Point)
		a.Mul(alpha)
		a.Add(v1.Point)

		b = mymath.Vector3dDiff(v2.Point, v1.Point)
		b.Mul(beta)
		b.Add(v1.Point)

		if a.X > b.X {
			a, b = b, a
		}

		var i1, i2 float64

		i1 = v1_intensity + (v3_intensity-v1_intensity)*alpha
		i2 = v1_intensity + (v2_intensity-v1_intensity)*beta

		if i1 > i2 {
			i1, i2 = i2, i1
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vector3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			intensity := i1 + (i2-i1)*phi

			// transform coordinates
			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.zbuf[int(x)][int(y)] < p.Z {
					engine.zbuf[int(x)][int(y)] = p.Z
					pixel_x, pixel_y := projection(cnv.height(), p)

					cnv.setPixel(pixel_x, pixel_y, object.Lightness(clr, intensity))
				}
			}
		}
	}

	for y := v2.Point.Y; y <= v3.Point.Y; y++ {
		segment_dy := v3.Point.Y - v2.Point.Y + 1
		alpha := float64((y - v1.Point.Y) / total_dy)
		beta := float64((y - v2.Point.Y) / segment_dy)

		var a, b mymath.Vector3d

		a = mymath.Vector3dDiff(v3.Point, v1.Point)
		a.Mul(alpha)
		a.Add(v1.Point)

		b = mymath.Vector3dDiff(v3.Point, v2.Point)
		b.Mul(beta)
		b.Add(v2.Point)

		if a.X > b.X {
			a, b = b, a
		}

		var i1, i2 float64

		i1 = v1_intensity + (v3_intensity-v1_intensity)*alpha
		i2 = v1_intensity + (v2_intensity-v1_intensity)*beta

		if i1 > i2 {
			i1, i2 = i2, i1
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   mymath.Vector3d
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			intensity := i1 + (i2-i1)*phi

			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.zbuf[int(x)][int(y)] < p.Z {
					engine.zbuf[int(x)][int(y)] = p.Z
					pixel_x, pixel_y := projection(cnv.height(), p)
					cnv.setPixel(pixel_x, pixel_y, object.Lightness(clr, intensity))
				}
			}
		}
	}
}

func (engine ZBufferGraphicsEngine) RenderWire(v1, v2, v3 object.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv
	cnv.drawLine(int(v1.Point.X), cnv.height()-int(v1.Point.X), int(v2.Point.X), cnv.height()-int(v2.Point.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v1.Point.X), cnv.height()-int(v1.Point.X), int(v3.Point.X), cnv.height()-int(v3.Point.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v2.Point.X), cnv.height()-int(v2.Point.X), int(v3.Point.X), cnv.height()-int(v3.Point.X), color.RGBA{0, 0, 0, 255})
}

func (engine ZBufferGraphicsEngine) RenderNormals(v1, v2, v3 object.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv
	cnv.drawLine(int(v1.Point.X), cnv.height()-int(v1.Point.X), int(v1.Point.X+v1.Normal.X), cnv.height()-int(v1.Point.X+v1.Normal.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v2.Point.X), cnv.height()-int(v2.Point.X), int(v2.Point.X+v2.Normal.X), cnv.height()-int(v2.Point.X+v2.Normal.X), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v3.Point.X), cnv.height()-int(v3.Point.X), int(v3.Point.X+v3.Normal.X), cnv.height()-int(v3.Point.X+v3.Normal.X), color.RGBA{0, 0, 0, 255})
}

func (engine ZBufferGraphicsEngine) RenderModel(obj object.PolygonialModel) {
	obj.IterateOverPolygons(engine.RenderPolygon)
}

func (engine ZBufferGraphicsEngine) RenderScene(scn object.Scene) {
	engine.light = scn.LightSource

	cnv := engine.Cnv
	zBack := float64(-10000)
	engine.initBuf(cnv.height(), cnv.width(), zBack)

	cnv.fill(scn.Background)

	for i := range scn.Objects {
		engine.RenderModel(scn.Objects[i])
	}

	pixel_x, pixel_y := projection(cnv.height(), scn.LightSource.Pos)
	cnv.setPixel(pixel_x, pixel_y, color.White)

}

func (engine ZBufferGraphicsEngine) Image() image.Image {
	return engine.Cnv.Image()
}
