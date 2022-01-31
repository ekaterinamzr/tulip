package object

import (
	"image/color"
	"tulip/mymath"
)

type Light struct {
	Intensity float64
	Pos       mymath.Vector3d
	Direction mymath.Vector3d
}

func NewLight(intensity float64, pos, direction mymath.Vector3d) *Light {
	light := new(Light)

	light.Intensity = intensity
	light.Pos = pos
	light.Direction = direction

	return light
}

func VertexIntensity(v Vertex, l Light) float64 {
	return l.Intensity * (1) * (float64(mymath.CosAlpha(l.Direction, v.Normal)))
}

func VectorIntensity(v mymath.Vector3d, l Light) float64 {
	intensity := l.Intensity * (1) * (float64(mymath.CosAlpha(l.Direction, v)))

	if intensity > 1 {
		intensity = 1
	}

	if intensity < 0 {
		intensity = 0
	}

	return intensity
}

func Lightness(clr color.NRGBA, intensity float64) color.NRGBA {
	if intensity > 1 || intensity < 0 {
		return clr
	}

	r := clr.R + uint8(float64((255-clr.R))*intensity)
	g := clr.G + uint8(float64((255-clr.G))*intensity)
	b := clr.B + uint8(float64((255-clr.B))*intensity)

	return color.NRGBA{r, g, b, 255}
}
