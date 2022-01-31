package object

import (
	"image/color"
	"tulip/mymath"
)

type Scene struct {
	Objects     []PolygonialModel
	LightSource Light

	Background color.Color
}

func (scn *Scene) AddObject(obj PolygonialModel) {
	scn.Objects = append(scn.Objects, obj)
}

func (scn *Scene) SetLight(intensity float64, pos, dir mymath.Vector3d) {
	light := NewLight(intensity, pos, dir)
	scn.LightSource = *light
}

func (scn *Scene) SetBackground(clr color.Color) {
	scn.Background = clr
}
