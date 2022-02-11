package scene

import (
	"image/color"
	"tulip/mymath"
)

type Vertex struct {
	Point mymath.Vec3d
	Normal mymath.Vec3d
}

type Polygon struct {
	V1 int
	V2 int
	V3 int

	Clr color.NRGBA
}


func (v *Vertex) Scale(center mymath.Vec3d, k float64) {
	v.Point.Scale(center, k)
}

func (v *Vertex) Move(delta mymath.Vec3d) {
	v.Point.Move(delta)
}

func (v *Vertex) Reflect(x, y, z bool) {
	v.Point.Reflect(x, y, z)
	v.Normal.Reflect(x, y, z)
}

func (v *Vertex) Rotate(center, angles mymath.Vec3d) {
	v.Point.Rotate(center, angles)
	v.Normal.Rotate(mymath.MakeVec3d(0, 0, 0), angles)
}
