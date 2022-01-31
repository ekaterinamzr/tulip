package object

import (
	"math"
)

type Vector3d struct {
	Point
}

func Make(x, y, z float64) Vector3d {
	var v Vector3d

	v.X = x
	v.Y = y
	v.Z = z

	return v
}

// a, b, c are position Vector3ds
func MakeNormal(a, b, c Point) Vector3d {
	var v Vector3d

	// v.X = (b.X-a.X)*(c.Z-a.Z) - (c.Y-a.Y)*(b.Z-a.Z)
	// v.Y = (b.X-a.X)*(c.Z-a.Z) - (c.X-a.X)*(b.Z-a.Z)
	// v.Z = (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)

	v.X = (b.Y-a.Y)*(c.Z-a.Z) - (b.Z-a.Z)*(c.Y-a.Y)
	v.Y = (b.Z-a.Z)*(c.X-a.X) - (b.X-a.X)*(c.Z-a.Z)
	v.Z = (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)

	// k := 1.0 / v.Length()
	// v.Mul(k)

	return v
}

func (a Vector3d) Cross(b Vector3d) Vector3d {
	var v Vector3d

	// v.X = (b.X-a.X)*(c.Z-a.Z) - (c.Y-a.Y)*(b.Z-a.Z)
	// v.Y = (b.X-a.X)*(c.Z-a.Z) - (c.X-a.X)*(b.Z-a.Z)
	// v.Z = (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)

	v.X = (a.Y)*(b.Z) - (a.Z)*(b.Y)
	v.Y = (a.Z)*(b.X) - (a.X)*(b.Z)
	v.Z = (a.X)*(b.Y) - (a.Y)*(b.X)

	// k := 1.0 / v.Length()
	// v.Mul(k)

	return v
}

func MakeTwoPoints(a, b Point) Vector3d {
	var v Vector3d

	v.X = b.X - a.X
	v.Y = b.Y - a.Y
	v.Z = b.Z - a.Z

	return v
}

func (v1 *Vector3d) Add(v2 Vector3d) {
	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z
}

func (v1 *Vector3d) Mul(k float64) {
	v1.X *= k
	v1.Y *= k
	v1.Z *= k
}

func (v Vector3d) Length() float64 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	//fmt.Println(length)
	return length
}

func cosAlpha(v1, v2 Vector3d) float64 {
	if v1.Length() == 0 || v2.Length() == 0 {
		return 0
	}
	scalar := v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
	return scalar / (v1.Length() * v2.Length())
}
