package object

import (
	"image/color"
)

type Scene struct {
	Objects     []PolygonialModel
	LightSource Light

	BackGround color.Color
}

func (scn *Scene) AddObject(obj PolygonialModel) {
	scn.Objects = append(scn.Objects, obj)
}
