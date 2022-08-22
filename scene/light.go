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

func (l Light) NormalIntensity(n mymath.Vec3) float64 {
	if l.Intensity < 0 {
		return 0
	}
	
	i := l.Intensity * (1) * (float64(mymath.CosAlpha(l.Direction, n)))

	if i > 1 {
		i = 1
	}

	if i < 0 {
		i = 0
	}

	return i
}
