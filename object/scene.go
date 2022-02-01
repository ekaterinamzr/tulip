package object

import (
	"image/color"
	"tulip/mymath"
)

type Scene struct {
	Objects     []PolygonialModel
	LightSource Light

	Camera Camera

	Background color.NRGBA
	Ground     PolygonialModel
	GroundClr  color.NRGBA
}

func (scn *Scene) AddObject(obj PolygonialModel) {
	scn.Objects = append(scn.Objects, obj)
}

func (scn *Scene) SetLight(intensity float64, pos, dir mymath.Vector3d) {
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

func (scn *Scene) SetGround(g mymath.Vector3d) {
	var (
		ground         Model
		v1, v2, v3, v4 Vertex
	)

	v1.Point = g
	v1.Normal = mymath.MakeVector3d(0, 1, 0)

	v2.Point = mymath.MakeVector3d(-g.X, g.Y, g.Z)
	v2.Normal = mymath.MakeVector3d(0, 1, 0)

	v3.Point = mymath.MakeVector3d(-g.X, g.Y, -g.Z)
	v3.Normal = mymath.MakeVector3d(0, 1, 0)

	v4.Point = mymath.MakeVector3d(g.X, g.Y, -g.Z)
	v4.Normal = mymath.MakeVector3d(0, 1, 0)

	ground.Vertices = append(ground.Vertices, v1, v2, v3, v4)
	ground.AddPolygon(0, 1, 2, scn.GroundClr)
	ground.AddPolygon(2, 3, 0, scn.GroundClr)

	//ground.Rotate(mymath.MakeVector3d(0, 0, 0), mymath.MakeVector3d(1, 0, 0))

	ground.Move(mymath.MakeVector3d(0, 0, 200))

	scn.Ground = &ground

	// scn.AddObject(&ground)
}
