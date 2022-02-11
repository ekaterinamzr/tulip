package graphics

import (
	"image"
	"image/color"
	"tulip/scene"
)

type GraphicsEngine interface {
	//RenderVertex(cnv Canvas, vertex scene.Vertex)
	RenderPolygon(cnv Canvas, v1, v2, v3 scene.Vertex, clr color.NRGBA)
	RenderModel(cnv Canvas, scene scene.PolygonialModel)

	RenderScene(cnv Canvas, scn scene.Scene)

	Image() image.Image
}
