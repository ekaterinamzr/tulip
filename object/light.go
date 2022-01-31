package object

import (
	"image/color"
	//"tulip/vector3d"
)

type Light struct {
	//clr color.Color
	Intensity float64
	Pos       Point
	Direction Vector3d
}

func NewLight(intensity float64, pos Point) *Light {
	light := new(Light)

	//light.clr = clr
	light.Intensity = intensity
	light.Pos = pos
	light.Direction = Make(-1.0, 0.0, 0.0)

	return light
}

func CalculateIntensity(v Vertex, l Light) float64 {
	//lightDir := MakeTwoPoints(v.Point, l.pos)
	//fmt.Println(v.Normal)
	return l.Intensity * (1) * (float64(cosAlpha(l.Direction, v.Normal)))
}

func CalculateIntensityVector(v Vector3d, l Light) float64 {
	//lightDir := MakeTwoPoints(v.Point, l.pos)
	//fmt.Println(v.Normal)
	intensity := l.Intensity * (1) * (float64(cosAlpha(l.Direction, v)))

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
