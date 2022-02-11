package graphics

import (
	"tulip/scene"
	"image/color"
)

type triangle struct {
	v0, v1, v2 scene.Vertex
	clr color.NRGBA
}

func makeTriangle(v0, v1, v2 scene.Vertex, clr color.NRGBA) triangle{
	var t triangle

	t.v0, t.v1, t.v2 = v0, v1, v2
	t.clr = clr
	
	return t
}