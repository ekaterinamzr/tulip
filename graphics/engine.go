package graphics

import (
	"image"
	"tulip/scene"
)

type GraphicsEngine interface {
	RenderModel(cnv Canvas, scene scene.PolygonialModel)

	RenderScene(cnv Canvas, scn scene.Scene)

	Image() image.Image
}
