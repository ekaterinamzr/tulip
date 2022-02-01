package mymath

import "math"

type Matrix4x4 [4][4]float64

func MulVectorMatrix(vec Vector3d, m Matrix4x4) Vector3d {
	var res Vector3d

	res.X = vec.X*m[0][0] + vec.Y*m[1][0] + vec.Z*m[2][0] + m[3][0]
	res.Y = vec.X*m[0][1] + vec.Y*m[1][1] + vec.Z*m[2][1] + m[3][1]
	res.Z = vec.X*m[0][2] + vec.Y*m[1][2] + vec.Z*m[2][2] + m[3][2]
	res.w = vec.X*m[0][3] + vec.Y*m[1][3] + vec.Z*m[2][3] + m[3][3]

	// if k != 0.0 {
	// 	res.Div(k)
	// }

	return res
}

func MakeIdentityM() Matrix4x4 {
	var m Matrix4x4

	m[0][0] = 1.0
	m[1][1] = 1.0
	m[2][2] = 1.0
	m[3][3] = 1.0

	return m
}

func MakeTranslationM(x, y, z float64) Matrix4x4 {
	var m Matrix4x4

	m[0][0] = 1.0
	m[1][1] = 1.0
	m[2][2] = 1.0
	m[3][3] = 1.0
	m[3][0] = x
	m[3][1] = y
	m[3][2] = z

	return m
}

func MakePointAtM(pos, target, up Vector3d) Matrix4x4 {
	newForward := Vector3dDiff(target, pos)
	newForward.Normalize()

	a := Vector3dMul(newForward, up.DotProduct(newForward))
	newUp := Vector3dDiff(up, a)
	newUp.Normalize()

	newRight := newUp.CrossProduct(newForward)

	var m Matrix4x4
	m[0][0] = newRight.X
	m[0][1] = newRight.Y
	m[0][2] = newRight.Z
	m[0][3] = 0.0

	m[1][0] = newUp.X
	m[1][1] = newUp.Y
	m[1][2] = newUp.Z
	m[1][3] = 0.0

	m[2][0] = newForward.X
	m[2][1] = newForward.Y
	m[2][2] = newForward.Z
	m[2][3] = 0.0

	m[3][0] = pos.X
	m[3][1] = pos.Y
	m[3][2] = pos.Z
	m[3][3] = 1.0

	return m
}

func InverseTranslationM(m Matrix4x4) Matrix4x4 {
	var res Matrix4x4

	res[0][0] = m[0][0]
	res[0][1] = m[1][0]
	res[0][2] = m[2][0]
	res[0][3] = 0.0

	res[1][0] = m[0][1]
	res[1][1] = m[1][1]
	res[1][2] = m[2][1]
	res[1][3] = 0.0

	res[2][0] = m[0][2]
	res[2][1] = m[1][2]
	res[2][2] = m[2][2]
	res[2][3] = 0.0

	res[3][0] = -(m[3][0]*res[0][0] + m[3][1]*res[1][0] + m[3][2]*res[2][0])
	res[3][1] = -(m[3][0]*res[0][1] + m[3][1]*res[1][1] + m[3][2]*res[2][1])
	res[3][2] = -(m[3][0]*res[0][2] + m[3][1]*res[1][2] + m[3][2]*res[2][2])
	res[3][3] = 1.0

	return res
}

func MulMatrix(m1, m2 Matrix4x4) Matrix4x4 {
	var m Matrix4x4

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			m[r][c] = m1[r][0]*m2[0][c] + m1[r][1]*m2[1][c] + m1[r][2]*m2[2][c] + m1[r][3]*m2[3][c]
		}
	}

	return m
}

func MakeRotationXM(fAngleRad float64) Matrix4x4 {
	var m Matrix4x4
	m[0][0] = 1.0
	m[1][1] = math.Cos(fAngleRad)
	m[1][2] = math.Sin(fAngleRad)
	m[2][1] = -math.Sin(fAngleRad)
	m[2][2] = math.Cos(fAngleRad)
	m[3][3] = 1.0
	return m
}

func MakeRotationYM(fAngleRad float64) Matrix4x4 {
	var m Matrix4x4
	m[0][0] = math.Cos(fAngleRad)
	m[2][2] = math.Cos(fAngleRad)
	m[0][2] = math.Sin(fAngleRad)
	m[2][0] = -math.Sin(fAngleRad)
	m[1][1] = 1.0
	m[3][3] = 1.0
	return m
}

func MakeRotationZM(fAngleRad float64) Matrix4x4 {
	var m Matrix4x4
	m[2][2] = 1.0
	m[0][0] = math.Cos(fAngleRad)
	m[0][1] = math.Sin(fAngleRad)
	m[1][0] = -math.Sin(fAngleRad)
	m[1][1] = math.Cos(fAngleRad)
	m[3][3] = 1.0
	return m
}
