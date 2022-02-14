package mymath

import (
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func MakeVec3(x, y, z float64) Vec3 {
	var vec Vec3

	vec.X, vec.Y, vec.Z = x, y, z

	return vec
}


func (a Vec3) DotProduct(b Vec3) float64 {
	res := a.X*b.X + a.Y*b.Y + a.Z*b.Z
	return res
}

func (a Vec3) CrossProduct(b Vec3) Vec3 {
	var vec Vec3

	vec.X = (a.Y)*(b.Z) - (a.Z)*(b.Y)
	vec.Y = (a.Z)*(b.X) - (a.X)*(b.Z)
	vec.Z = (a.X)*(b.Y) - (a.Y)*(b.X)

	return vec
}

func (vec Vec3) Length() float64 {
	length := math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
	return length
}

func (vec *Vec3) Normalize() {
	vec.Div(vec.Length())
}

func CosAlpha(a, b Vec3) float64 {
	if a.Length() == 0 || b.Length() == 0 {
		return 0
	}

	return a.DotProduct(b) / (a.Length() * b.Length())
}

func (a *Vec3) Add(b Vec3) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}

func (a *Vec3) Sub(b Vec3) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
}

func (a *Vec3) Mul(k float64) {
	a.X *= k
	a.Y *= k
	a.Z *= k
}

func (a *Vec3) Div(k float64) {
	if k != 0 {
		a.X /= k
		a.Y /= k
		a.Z /= k
	}
}

func Vec3Sum(a, b Vec3) Vec3 {
	a.Add(b)
	return a
}

func Vec3Diff(a, b Vec3) Vec3 {
	a.Sub(b)
	return a
}

func Vec3Mul(a Vec3, k float64) Vec3 {
	a.Mul(k)
	return a
}

func Vec3Div(a Vec3, k float64) Vec3 {
	a.Div(k)
	return a
}

//

func (vec *Vec3) Scale(center Vec3, k float64) {
	vec.Sub(center)
	vec.Mul(k)
	vec.Add(center)
}

func Vec3Scale(vec, center Vec3, k float64) Vec3 {
	vec.Scale(center, k)
	return vec
}

func (vec *Vec3) Move(delta Vec3) {
	vec.Add(delta)
}

func Vec3Move(vec, delta Vec3) Vec3 {
	vec.Move(delta)
	return vec
}

// Reflection over (0, 0, 0)
func (vec *Vec3) Reflect(x, y, z bool) {
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

func Vec3Reflect(vec Vec3, x, y, z bool) Vec3 {
	vec.Reflect(x, y, z)
	return vec
}

func toRadians(angle float64) float64 {
	return angle * 3.14 / 180
}

func (vec *Vec3) rotateZ(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*vec.X + sin*vec.Y
	newY := cos*vec.Y - sin*vec.X

	vec.X = newX
	vec.Y = newY
}

func (vec *Vec3) rotateX(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newZ := cos*vec.Z - sin*vec.Y
	newY := sin*vec.Z + cos*vec.Y

	vec.Z = newZ
	vec.Y = newY
}

func (vec *Vec3) rotateY(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*vec.X + sin*vec.Z
	newZ := cos*vec.Z - sin*vec.X

	vec.X = newX
	vec.Z = newZ
}

// angles in degrees
func (vec *Vec3) Rotate(center, angles Vec3) {
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

func Vec3Rotate(vec, center, angles Vec3) Vec3 {
	vec.Rotate(center, angles)
	return vec
}

func (vec3 Vec3) ToVec4(w... float64) Vec4 {
	return Vec3ToVec4(vec3, w...)
}
