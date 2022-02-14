package graphics

import (
	"image/color"
	"tulip/mymath"
	"tulip/scene"
	"math"
)

type shadowMap struct {
	light      scene.Light
	projection mymath.Matrix4x4
	view       mymath.Matrix4x4
	ViewProj   mymath.Matrix4x4
}

func makeShadowMap(light scene.Light, proj mymath.Matrix4x4) shadowMap {
	var sm shadowMap
	// for i := range buf {
	// 	buf[i] = 10000
	// }

	// sun view matrix
	vUp := mymath.MakeVec3(0, 1, 0)
	vTarget := mymath.Vec3Sum(light.Pos, light.Direction)
	mCamera := mymath.MakePointAtM(light.Pos, vTarget, vUp)
	sm.view = mymath.InverseTranslationM(mCamera)
	sm.projection = proj

	sm.ViewProj = mymath.MulMatrices(sm.view, sm.projection)

	sm.light = light

	return sm

}

func (sm shadowMap) vs(v scene.Vertex) Vertex {
	var vS Vertex
	vS.Point = mymath.MulVecMat(v.Point, sm.ViewProj)
	vS.Normal = v.Normal
	// vS.Normal = mymath.MulVecMat(v.Normal, g.projection)
	vS.clr = color.NRGBA{255, 255, 255, 255}
	vS.worldPoint = v.Point

	return vS
}

func (sm shadowMap) ps(v Vertex, clr color.NRGBA) (int, int, color.NRGBA) {
	pixelX := int(math.Round(v.Point.X))
	pixelY := int(math.Round(v.Point.Y))
	pixelClr := color.NRGBA{255, 255, 255, 255}

	return pixelX, pixelY, pixelClr
}
