package object

import (
	"image/color"
	"math"
)

type Vertex struct {
	Point

	Normal Vector3d

	cnt int
}

type Point struct {
	X float64
	Y float64
	Z float64

	//N Vector3d
}

func MakePoint(x, y, z float64) Point {
	var p Point

	p.X, p.Y, p.Z = x, y, z

	return p
}

type Polygon struct {
	V1 int
	V2 int
	V3 int

	Clr color.NRGBA
}

func PolygonNormal(v1, v2, v3 Vertex) Vector3d {
	return MakeNormal(v1.Point, v2.Point, v3.Point)
}

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Point) Sub(b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Point) Mul(k float64) Point {
	return Point{a.X * k, a.Y * k, a.Z * k}
}

func (v *Point) Scale(center Point, k float64) {
	v.X = center.X + k*(v.X-center.X)
	v.Y = center.Y + k*(v.Y-center.Y)
	v.Z = center.Z + k*(v.Z-center.Z)
}

func (v *Point) Move(delta Point) {
	v.X += delta.X
	v.Y += delta.Y
	v.Z += delta.Z
}

func (v *Point) Flip(x, y, z bool) {
	if x {
		v.X = -v.X
	}
	if y {
		v.X = -v.X
	}
	if z {
		v.Z = -v.Z
	}
}

func toRadians(angle float64) float64 {
	return angle * 3.14 / 180
}

func (v *Point) rotateZ(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*v.X + sin*v.Y
	newY := cos*v.Y - sin*v.X

	v.X = newX
	v.Y = newY
}

func (v *Point) rotateX(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newZ := cos*v.Z - sin*v.Y
	newY := sin*v.Z + cos*v.Y

	v.Z = newZ
	v.Y = newY
}

func (v *Point) rotateY(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*v.X + sin*v.Z
	newZ := cos*v.Z - sin*v.X

	v.X = newX
	v.Z = newZ
}

func (v *Point) Rotate(center, angles Point) {
	angleX := toRadians(angles.X)
	angleY := toRadians(angles.Y)
	angleZ := toRadians(angles.Z)

	v.Move(Point{-center.X, -center.Y, -center.Z})

	if angleX != 0 {
		v.rotateX(angleX)
	}

	if angleY != 0 {
		v.rotateY(angleY)
	}

	if angleZ != 0 {
		v.rotateZ(angleZ)
	}

	v.Move(center)
}

func (v *Vertex) Scale(center Point, k float64) {
	v.Point.Scale(center, k)
	//v.Normal.Scale(center, k)
}

func (v *Vertex) Move(delta Point) {
	v.Point.Move(delta)
	//v.Normal.Move(delta)
}

func (v *Vertex) Flip(x, y, z bool) {
	v.Point.Flip(x, y, z)
	v.Normal.Flip(x, y, z)
}

func (v *Vertex) Rotate(center, angles Point) {
	v.Point.Rotate(center, angles)
	v.Normal.Rotate(MakePoint(0, 0, 0), angles)
}
