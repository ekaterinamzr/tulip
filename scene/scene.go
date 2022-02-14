package scene

import (
	"image/color"
	"tulip/mymath"
)

type Scene struct {
	Objects     []Model
	LightSource Light

	Camera Camera

	Background color.NRGBA
	Ground     Model
	GroundClr  color.NRGBA
}

func (scn *Scene) Add(m... Model) {
	for i := range(m) {
		scn.Objects = append(scn.Objects, m[i])
	}
	
}

func (scn *Scene) SetLight(intensity float64, pos, dir mymath.Vec3) {
	light := NewLight(intensity, pos, dir)
	scn.LightSource = *light
}

func (scn *Scene) SetCamera(Cam Camera) {
	scn.Camera = Cam
}

func (scn *Scene) SetBackground(clr color.NRGBA) {
	scn.Background = clr
}

func (scn *Scene) SetGroundClr(clr color.NRGBA) {
	scn.GroundClr = clr
}

// func (scn *Scene) SetSquareGround(g float64) {
// 	ground := primitives.NewBlock(g, g, 1, mymath.MakeVec3(0, -g/2, 0), scn.GroundClr)
// 	scn.Ground = ground
// }
func (scn *Scene) SetGround(g mymath.Vec3) {
	var (
		ground         Model
		v1, v2, v3, v4 Vertex
	)

	v1.Point = mymath.Vec3ToVec4(g)
	v1.Normal = mymath.MakeVec4(0, -1, 0)

	v2.Point = mymath.MakeVec4(-g.X, g.Y, g.Z)
	v2.Normal = mymath.MakeVec4(0, -1, 0)

	v3.Point = mymath.MakeVec4(-g.X, g.Y, -g.Z)
	v3.Normal = mymath.MakeVec4(0, -1, 0)

	v4.Point = mymath.MakeVec4(g.X, g.Y, -g.Z)
	v4.Normal = mymath.MakeVec4(0, -1, 0)

	ground.Vertices = append(ground.Vertices, v1, v2, v3, v4)
	ground.AddPolygon(0, 1, 2, scn.GroundClr)
	ground.AddPolygon(2, 3, 0, scn.GroundClr)

	ground.Move(mymath.MakeVec3(0, 0, 200))

	scn.Ground = ground
}
