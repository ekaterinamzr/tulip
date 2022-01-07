package object

import (
	"image/color"
	"tulip/vector"
)

type Light struct {
	clr color.Color
	location Vertex
	direction vector.Vector
}

func NewLight (clr color.Color, location Vertex, direction vector.Vector) *Light {
	light := new(Light)
	
	light.clr = clr
	light.location = location
	light.direction = direction

	return light
}

func Lightness (clr color.NRGBA, intensity float32) color.NRGBA {
	if intensity > 1 || intensity < 0 {
		return clr
	}
	
	r := clr.R + uint8(float32((255 - clr.R)) * intensity)
	g := clr.G + uint8(float32((255 - clr.G)) * intensity)
	b := clr.B + uint8(float32((255 - clr.B)) * intensity)

	return color.NRGBA{r, g, b, 255}
}