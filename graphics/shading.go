package graphics

import (
	"image/color"
	"math"
	"tulip/mymath"
	"tulip/scene"
)

type shaderInterface interface {
	vs(v scene.Vertex) Vertex
	ps(v Vertex, clr color.NRGBA) (int, int, color.NRGBA)
}

type Vertex struct {
	Point      mymath.Vec4
	Normal     mymath.Vec4
	worldPoint mymath.Vec4
	Intensity  float64
	clr        color.NRGBA
}

type gouraudShader struct {
	light scene.Light

	projection    mymath.Matrix4x4
	view          mymath.Matrix4x4
	viewProj      mymath.Matrix4x4
	lightViewProj mymath.Matrix4x4
	pst           psTransformer

	sBuf [][]float64
}

func makeGouraudShader(view, proj, lViewProj mymath.Matrix4x4, sBuf [][]float64, light scene.Light, pst psTransformer) gouraudShader {
	var g gouraudShader

	g.projection = proj
	g.view = view
	g.viewProj = mymath.MulMatrices(g.view, g.projection)
	g.light = light
	g.lightViewProj = lViewProj
	g.sBuf = sBuf
	g.pst = pst

	return g
}

func intensity(n mymath.Vec3, l scene.Light) float64 {
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

// Vertex shader
func (g gouraudShader) vs(v scene.Vertex) Vertex {
	var vS Vertex
	vS.Point = mymath.MulVecMat(v.Point, g.viewProj)
	vS.Normal = v.Normal
	// vS.Normal = mymath.MulVecMat(v.Normal, g.projection)
	vS.clr = v.Clr
	vS.worldPoint = v.Point
	vS.Intensity = intensity(v.Normal.Vec3, g.light)

	return vS
}

func Tint(clr color.NRGBA, intensity float64) color.NRGBA {
	if intensity > 1 || intensity < 0 {
		return clr
	}

	r := clr.R + uint8(float64((255-clr.R))*intensity)
	g := clr.G + uint8(float64((255-clr.G))*intensity)
	b := clr.B + uint8(float64((255-clr.B))*intensity)

	return color.NRGBA{r, g, b, 255}
}

func Shade(clr color.NRGBA, intensity float64) color.NRGBA {
	if intensity > 1 || intensity < 0 {
		return clr
	}

	r := uint8(float64(clr.R) * (1.0-intensity))
	g := uint8(float64(clr.G) * (1.0-intensity))
	b := uint8(float64(clr.B) * (1.0-intensity))

	return color.NRGBA{r, g, b, 255}
}

// Pixel shader
func (g gouraudShader) ps(v Vertex, clr color.NRGBA) (int, int, color.NRGBA) {
	px := int(math.Round(v.Point.X))
	py := int(math.Round(v.Point.Y))

	lightPoint := mymath.MulVecMat(v.worldPoint, g.lightViewProj)
	g.pst.transformVec4(&lightPoint)
	//idx := int((math.Round)(lightPoint.X)) + int((math.Round)(lightPoint.Y))*500
	// lpx := int(math.Round(lightPoint.X))
	// lpy := int(math.Round(lightPoint.Y))

	clr = Shade(clr, 0.5 - g.light.Intensity)
	pixelClr := Tint(clr, v.Intensity)

	// if idx < len(g.sBuf) && idx >= 0 && lightPoint.Z > g.sBuf[idx]+0.0006 {
	// 	fmt.Println(lightPoint.Z, g.sBuf[idx])
	// 	pixelClr = color.NRGBA{0, 0, 0, 255}
	// }

	// if lpx >= 0 && lpy >= 0 && lpx < 500.0 && lpy < 500.0 && lightPoint.Z > g.sBuf[lpx][lpy] {
	// 	// fmt.Println(lightPoint.Z, g.sBuf[lpx][lpy])
	// 	pixelClr = color.NRGBA{0, 0, 0, 255}
	// }

	return px, py, pixelClr
}
