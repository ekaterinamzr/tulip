package mymath

import (
	"math"
)

type Vector3d struct {
	X float64
	Y float64
	Z float64
}

func MakeVector3d(x, y, z float64) Vector3d {
	var vec Vector3d

	vec.X, vec.Y, vec.Z = x, y, z

	return vec
}

func (a Vector3d) DotProduct(b Vector3d) float64 {
	res := a.X*b.X + a.Y*b.Y + a.Z*b.Z
	return res
}

func (a Vector3d) CrossProduct(b Vector3d) Vector3d {
	var vec Vector3d

	vec.X = (a.Y)*(b.Z) - (a.Z)*(b.Y)
	vec.Y = (a.Z)*(b.X) - (a.X)*(b.Z)
	vec.Z = (a.X)*(b.Y) - (a.Y)*(b.X)

	return vec
}

// a, b, c are position Vector3ds
func MakeNormal(a, b, c Vector3d) Vector3d {
	vec1 := Vector3dDiff(b, a)
	vec2 := Vector3dDiff(c, a)

	return vec1.CrossProduct(vec2)
}

func (vec Vector3d) Length() float64 {
	length := math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
	return length
}

func CosAlpha(a, b Vector3d) float64 {
	if a.Length() == 0 || b.Length() == 0 {
		return 0
	}

	return a.DotProduct(b) / (a.Length() * b.Length())
}

func (a *Vector3d) Add(b Vector3d) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
}

func (a *Vector3d) Sub(b Vector3d) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
}

func (a *Vector3d) Mul(k float64) {
	a.X *= k
	a.Y *= k
	a.Z *= k
}

func (a *Vector3d) Div(k float64) {
	if k != 0 {
		a.X /= k
		a.Y /= k
		a.Z /= k
	}
}

func Vector3dSum(a, b Vector3d) Vector3d {
	a.Add(b)
	return a
}

func Vector3dDiff(a, b Vector3d) Vector3d {
	a.Sub(b)
	return a
}

func Vector3dMul(a Vector3d, k float64) Vector3d {
	a.Mul(k)
	return a
}

func Vector3dDiv(a Vector3d, k float64) Vector3d {
	a.Div(k)
	return a
}

//

func (vec *Vector3d) Scale(center Vector3d, k float64) {
	vec.Sub(center)
	vec.Mul(k)
	vec.Add(center)
}

func Vector3dScale(vec, center Vector3d, k float64) Vector3d {
	vec.Scale(center, k)
	return vec
}

func (vec *Vector3d) Move(delta Vector3d) {
	vec.Add(delta)
}

func Vector3dMove(vec, delta Vector3d) Vector3d {
	vec.Move(delta)
	return vec
}

// Reflection over (0, 0, 0)
func (vec *Vector3d) Reflect(x, y, z bool) {
	if x {
		vec.X = -vec.X
	}
	if y {
		vec.X = -vec.X
	}
	if z {
		vec.Z = -vec.Z
	}
}

func Vector3dReflect(vec Vector3d, x, y, z bool) Vector3d {
	vec.Reflect(x, y, z)
	return vec
}

func toRadians(angle float64) float64 {
	return angle * 3.14 / 180
}

func (vec *Vector3d) rotateZ(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*vec.X + sin*vec.Y
	newY := cos*vec.Y - sin*vec.X

	vec.X = newX
	vec.Y = newY
}

func (vec *Vector3d) rotateX(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newZ := cos*vec.Z - sin*vec.Y
	newY := sin*vec.Z + cos*vec.Y

	vec.Z = newZ
	vec.Y = newY
}

func (vec *Vector3d) rotateY(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*vec.X + sin*vec.Z
	newZ := cos*vec.Z - sin*vec.X

	vec.X = newX
	vec.Z = newZ
}

// angles in degrees
func (vec *Vector3d) Rotate(center, angles Vector3d) {
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

func Vector3dRotate(vec, center, angles Vector3d) Vector3d {
	vec.Rotate(center, angles)
	return vec
}

func (vec *Vector3d) MulMatrix(m Matrix4x4) {
	*vec = MulVectorMatrix(*vec, m)
}
