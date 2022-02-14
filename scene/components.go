package scene

import (
	"image/color"
	"tulip/mymath"
)

type Vertex struct {
	Point  mymath.Vec4
	Normal mymath.Vec4

	Clr color.NRGBA

	// Intensity float64
}

type Polygon struct {
	V1 int
	V2 int
	V3 int

	Clr color.NRGBA
}

func (v *Vertex) Scale(center mymath.Vec3, k float64) {
	v.Point.Vec3.Scale(center, k)
}

func (v *Vertex) Move(delta mymath.Vec3) {
	v.Point.Vec3.Move(delta)
}

func (v *Vertex) Reflect(x, y, z bool) {
	v.Point.Vec3.Reflect(x, y, z)
	v.Normal.Vec3.Reflect(x, y, z)
}

func (v *Vertex) Rotate(center, angles mymath.Vec3) {
	v.Point.Vec3.Rotate(center, angles)
	v.Normal.Vec3.Rotate(mymath.MakeVec3(0, 0, 0), angles)
}
