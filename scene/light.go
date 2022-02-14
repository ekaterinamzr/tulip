package scene

import (
	"tulip/mymath"
)

type Light struct {
	Intensity float64
	Pos       mymath.Vec3
	Direction mymath.Vec3
}

func NewLight(intensity float64, pos, direction mymath.Vec3) *Light {
	light := new(Light)

	light.Intensity = intensity
	light.Pos = pos
	light.Direction = direction

	return light
}
