package graphics

import (
	"image"
	"image/color"
	"tulip/object"
)

type GraphicsEngine interface {
	RenderVertex(cnv Canvas, vertex object.Vertex)
	RenderPolygon(cnv Canvas, v1, v2, v3 object.Vertex, clr color.Color)
	RenderModel(cnv Canvas, object object.Model)

	RenderScene(cnv Canvas, scn object.Scene)

	Image() image.Image
}

type ZBufferGraphicsEngine struct {
	Cnv  Canvas
	zbuf [][]float64
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

func projection(h int, v object.Vertex) (int, int) {
	return int(v.X), h - int(v.Y)
}

func (engine ZBufferGraphicsEngine) RenderPolygon(v1, v2, v3 object.Vertex, clr color.Color) {
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

	//fmt.Println("rendering")

	for y := v1.Y; y <= v2.Y; y++ {
		segment_dy := v2.Y - v1.Y + 1
		alpha := float64((y - v1.Y) / total_dy)
		beta := float64((y - v1.Y) / segment_dy)

		var a, b object.Vertex

		a = v3.Sub(v1)
		a = a.Mul(alpha)
		a = a.Add(v1)

		b = v2.Sub(v1)
		b = b.Mul(beta)
		b = b.Add(v1)

		// a.X = v1.X + (v3.X-v1.X)*alpha
		// a.Y = v1.Y + (v3.Y-v1.Y)*alpha
		// b.X = v1.X + (v2.X-v1.X)*beta
		// b.Y = v1.Y + (v2.Y-v1.Y)*beta

		if a.X > b.X {
			a, b = b, a
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   object.Vertex
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}
			// p := b.Sub(a)
			// p = p.Mul(phi)
			// p.Add(a)

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			// transform coordinates
			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.zbuf[int(x)][int(y)] < p.Z {
					engine.zbuf[int(x)][int(y)] = p.Z
					pixel_x, pixel_y := projection(cnv.height(), p)
					cnv.setPixel(pixel_x, pixel_y, clr)
				}
			}

			//fmt.Println(a.X, b.X, x)
		}

		//fmt.Println(y, v2.Y)
	}

	for y := v2.Y; y <= v3.Y; y++ {
		segment_dy := v3.Y - v2.Y + 1
		alpha := float64((y - v1.Y) / total_dy)
		beta := float64((y - v2.Y) / segment_dy)

		var a, b object.Vertex

		a = v3.Sub(v1)
		a = a.Mul(alpha)
		a = a.Add(v1)

		b = v3.Sub(v2)
		b = b.Mul(beta)
		b = b.Add(v2)

		// a.X = v1.X + (v3.X-v1.X)*alpha
		// a.Y = v1.Y + (v3.Y-v1.Y)*alpha
		// b.X = v2.X + (v3.X-v2.X)*beta
		// b.Y = v2.Y + (v3.Y-v2.Y)*beta

		if a.X > b.X {
			a, b = b, a
		}

		for x := a.X; x <= b.X; x++ {
			var (
				phi float64
				p   object.Vertex
			)

			if a.X == b.X {
				phi = float64(1)
			} else {
				phi = (x - a.X) / (b.X - a.X)
			}
			// p := b.Sub(a)
			// p = p.Mul(phi)
			// p.Add(a)

			p.Z = a.Z + (b.Z-a.Z)*phi

			p.X = x
			p.Y = y

			if x >= 0 && y >= 0 && x < float64(engine.Cnv.width()) && y < float64(engine.Cnv.height()) {
				if engine.zbuf[int(x)][int(y)] < p.Z {
					engine.zbuf[int(x)][int(y)] = p.Z
					pixel_x, pixel_y := projection(cnv.height(), p)
					cnv.setPixel(pixel_x, pixel_y, clr)
				}

			}
		}
	}

	//DrawPolygon(cnv, v1, v2, v3, color.NRGBA{0, 0, 255, 255})
}

func (engine ZBufferGraphicsEngine) RenderWire(v1, v2, v3 object.Vertex, clr color.Color) {
	cnv := engine.Cnv
	cnv.drawLine(int(v1.X), cnv.height()-int(v1.Y), int(v2.X), cnv.height()-int(v2.Y), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v1.X), cnv.height()-int(v1.Y), int(v3.X), cnv.height()-int(v3.Y), color.RGBA{0, 0, 0, 255})
	cnv.drawLine(int(v2.X), cnv.height()-int(v2.Y), int(v3.X), cnv.height()-int(v3.Y), color.RGBA{0, 0, 0, 255})
}

func (engine ZBufferGraphicsEngine) RenderModel(obj object.PolygonialModel) {

	//obj.SortVertices()
	obj.IterateOverPolygons(engine.RenderPolygon)
	//obj.IterateOverPolygons(engine.RenderWire)
	//DrawObject(cnv, obj)
}

// TODO NewGraphicsEngine

func (engine ZBufferGraphicsEngine) RenderScene(scn object.Scene) {
	cnv := engine.Cnv
	zBack := float64(-10000)
	engine.initBuf(cnv.height(), cnv.width(), zBack)

	cnv.fill(scn.BackGround)

	for i := range scn.Objects {
		engine.RenderModel(scn.Objects[i])
	}

}

func (engine ZBufferGraphicsEngine) Image() image.Image {
	return engine.Cnv.Image()
}
