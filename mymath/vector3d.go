package mymath

import (
	"math"
)

type Vec3d struct {
	X float64
	Y float64
	Z float64

	w float64
}

func MakeVec3d(x, y, z float64) Vec3d {
	var vec Vec3d

	vec.X, vec.Y, vec.Z = x, y, z
	vec.w = 1.0

	return vec
}

func (vec *Vec3d) DivW() {
	if vec.w != 0 {
		vec.Div(vec.w)
	}
}

func (a Vec3d) DotProduct(b Vec3d) float64 {
	res := a.X*b.X + a.Y*b.Y + a.Z*b.Z
	return res
}

func (a Vec3d) CrossProduct(b Vec3d) Vec3d {
	var vec Vec3d

	vec.X = (a.Y)*(b.Z) - (a.Z)*(b.Y)
	vec.Y = (a.Z)*(b.X) - (a.X)*(b.Z)
	vec.Z = (a.X)*(b.Y) - (a.Y)*(b.X)

	return vec
}

func (vec Vec3d) Length() float64 {
	length := math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
	return length
}

func (vec *Vec3d) Normalize() {
	vec.Div(vec.Length())
}

func CosAlpha(a, b Vec3d) float64 {
	if a.Length() == 0 || b.Length() == 0 {
		return 0
	}

	return a.DotProduct(b) / (a.Length() * b.Length())
}

func (a *Vec3d) Add(b Vec3d) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}

func (a *Vec3d) Sub(b Vec3d) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
}

func (a *Vec3d) Mul(k float64) {
	a.X *= k
	a.Y *= k
	a.Z *= k
}

func (a *Vec3d) Div(k float64) {
	if k != 0 {
		a.X /= k
		a.Y /= k
		a.Z /= k
	}
}

func Vec3dSum(a, b Vec3d) Vec3d {
	a.Add(b)
	return a
}

func Vec3dDiff(a, b Vec3d) Vec3d {
	a.Sub(b)
	return a
}

func Vec3dMul(a Vec3d, k float64) Vec3d {
	a.Mul(k)
	return a
}

func Vec3dDiv(a Vec3d, k float64) Vec3d {
	a.Div(k)
	return a
}

//

func (vec *Vec3d) Scale(center Vec3d, k float64) {
	vec.Sub(center)
	vec.Mul(k)
	vec.Add(center)
}

func Vec3dScale(vec, center Vec3d, k float64) Vec3d {
	vec.Scale(center, k)
	return vec
}

func (vec *Vec3d) Move(delta Vec3d) {
	vec.Add(delta)
}

func Vec3dMove(vec, delta Vec3d) Vec3d {
	vec.Move(delta)
	return vec
}

// Reflection over (0, 0, 0)
func (vec *Vec3d) Reflect(x, y, z bool) {
	if x {
		vec.X = -vec.X
	}
	if y {
		vec.Y = -vec.Y
	}
	if z {
		vec.Z = -vec.Z
	}
}

func Vec3dReflect(vec Vec3d, x, y, z bool) Vec3d {
	vec.Reflect(x, y, z)
	return vec
}

func toRadians(angle float64) float64 {
	return angle * 3.14 / 180
}

func (vec *Vec3d) rotateZ(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*vec.X + sin*vec.Y
	newY := cos*vec.Y - sin*vec.X

	vec.X = newX
	vec.Y = newY
}

func (vec *Vec3d) rotateX(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newZ := cos*vec.Z - sin*vec.Y
	newY := sin*vec.Z + cos*vec.Y

	vec.Z = newZ
	vec.Y = newY
}

func (vec *Vec3d) rotateY(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*vec.X + sin*vec.Z
	newZ := cos*vec.Z - sin*vec.X

	vec.X = newX
	vec.Z = newZ
}

// angles in degrees
func (vec *Vec3d) Rotate(center, angles Vec3d) {
	angleX := toRadians(angles.X)
	angleY := toRadians(angles.Y)
	angleZ := toRadians(angles.Z)

	vec.Sub(center)

	if angleX != 0 {
		vec.rotateX(angleX)
	}

	if angleY != 0 {
		vec.rotateY(angleY)
	}

	if angleZ != 0 {
		vec.rotateZ(angleZ)
	}

	vec.Add(center)
}

func Vec3dRotate(vec, center, angles Vec3d) Vec3d {
	vec.Rotate(center, angles)
	return vec
}

func (vec *Vec3d) MulMatrix(m Matrix4x4) {
	*vec = MulVectorMatrix(*vec, m)
}
