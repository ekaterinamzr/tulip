package graphics

import (
	//"fmt"
	"image"
	"image/color"
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

func projection(h int, v object.Point) (int, int) {
	return int(v.X), h - int(v.Y)
}

func (engine ZBufferGraphicsEngine) RenderPolygon(v1, v2, v3 object.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv

	if v1.Y > v2.Y {
		v1, v2 = v2, v1
	}
	if v1.Y > v3.Y {
		v1, v3 = v3, v1
	}
	if v2.Y > v3.Y {
		v2, v3 = v3, v2
	}

	total_dy := v3.Y - v1.Y

	light := engine.light

	v1_intensity := object.CalculateIntensity(v1, light)
	v2_intensity := object.CalculateIntensity(v2, light)
	v3_intensity := object.CalculateIntensity(v3, light)

	//fmt.Println(v1_intensity, v2_intensity, v3_intensity)

	//fmt.Println("rendering")

	for y := v1.Y; y <= v2.Y; y++ {
		segment_dy := v2.Y - v1.Y + 1
		alpha := float64((y - v1.Y) / total_dy)
		beta := float64((y - v1.Y) / segment_dy)

		var a, b object.Point

		a = v3.Sub(v1.Point)
		a = a.Mul(alpha)
		a = a.Add(v1.Point)

		b = v2.Sub(v1.Point)
		b = b.Mul(beta)
		b = b.Add(v1.Point)

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
				p   object.Point
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
					// cnv.setPixel(pixel_x, pixel_y, clr)
				}
			}
		}
	}

	for y := v2.Y; y <= v3.Y; y++ {
		segment_dy := v3.Y - v2.Y + 1
		alpha := float64((y - v1.Y) / total_dy)
		beta := float64((y - v2.Y) / segment_dy)

		var a, b object.Point

		a = v3.Sub(v1.Point)
		a = a.Mul(alpha)
		a = a.Add(v1.Point)

		b = v3.Sub(v2.Point)
		b = b.Mul(beta)
		b = b.Add(v2.Point)

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
				p   object.Point
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

			//fmt.Println(i1, i2, phi)

			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.zbuf[int(x)][int(y)] < p.Z {
					engine.zbuf[int(x)][int(y)] = p.Z
					pixel_x, pixel_y := projection(cnv.height(), p)
					cnv.setPixel(pixel_x, pixel_y, object.Lightness(clr, intensity))
					// cnv.setPixel(pixel_x, pixel_y, clr)
				}

			}
		}
	}

	//cnv.drawLine(int(v1.X), cnv.height()-int(v1.Y), int(v1.X+v1.Normal.X), cnv.height()-int(v1.Y+v1.Normal.Y), color.Black)
	//cnv.drawLine(int(v2.X), cnv.height()-int(v2.Y), int(v2.X+v2.Normal.X), cnv.height()-int(v2.Y+v2.Normal.Y), color.Black)
	//cnv.drawLine(int(v3.X), cnv.height()-int(v3.Y), int(v3.X+v3.Normal.X), cnv.height()-int(v3.Y+v3.Normal.Y), color.Black)
}

func (engine ZBufferGraphicsEngine) RenderWire(v1, v2, v3 object.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv
	cnv.drawLine(int(v1.X), cnv.height()-int(v1.Y), int(v2.X), cnv.height()-int(v2.Y), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v1.X), cnv.height()-int(v1.Y), int(v3.X), cnv.height()-int(v3.Y), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v2.X), cnv.height()-int(v2.Y), int(v3.X), cnv.height()-int(v3.Y), color.RGBA{0, 0, 0, 255})
}

func (engine ZBufferGraphicsEngine) RenderNormals(v1, v2, v3 object.Vertex, clr color.NRGBA) {
	cnv := engine.Cnv
	cnv.drawLine(int(v1.X), cnv.height()-int(v1.Y), int(v1.X+v1.Normal.X), cnv.height()-int(v1.Y+v1.Normal.Y), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v2.X), cnv.height()-int(v2.Y), int(v2.X+v2.Normal.X), cnv.height()-int(v2.Y+v2.Normal.Y), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v3.X), cnv.height()-int(v3.Y), int(v3.X+v3.Normal.X), cnv.height()-int(v3.Y+v3.Normal.Y), color.RGBA{0, 0, 0, 255})
}

func (engine ZBufferGraphicsEngine) RenderModel(obj object.PolygonialModel) {

	//obj.SortVertices()
	obj.IterateOverPolygons(engine.RenderPolygon)
	//obj.IterateOverPolygons(engine.RenderWire)
	//obj.IterateOverPolygons(engine.RenderNormals)
	//DrawObject(cnv, obj)
}

// TODO NewGraphicsEngine

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
