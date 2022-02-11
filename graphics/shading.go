package graphics

import (
	"image/color"
	"tulip/mymath"
	"tulip/scene"
)

type gouraudShader struct {
	light scene.Light
}

func intensity(n mymath.Vec3, l scene.Light) float64 {
	i := l.Intensity * (1) * (float64(mymath.CosAlpha(l.Direction, n)))

	if i > 1 {
		i = 1
	}

	if i < 0 {
		i = 0
	}

	return i
}

// Vertex shader
func (g gouraudShader) vs(v *scene.Vertex) {
	v.Intensity = intensity(v.Normal.Vec3, g.light)
}

func lightness(clr color.NRGBA, intensity float64) color.NRGBA {
	if intensity > 1 || intensity < 0 {
		return clr
	}

	r := clr.R + uint8(float64((255-clr.R))*intensity)
	g := clr.G + uint8(float64((255-clr.G))*intensity)
	b := clr.B + uint8(float64((255-clr.B))*intensity)

	return color.NRGBA{r, g, b, 255}
}

// Pixel shader

func (g gouraudShader) ps(v scene.Vertex, clr color.NRGBA) (int, int, color.NRGBA) {
	pixelX := int(v.Point.X)
	pixelY := int(v.Point.Y)
	pixelClr := lightness(clr, v.Intensity)

	return pixelX, pixelY, pixelClr
}
