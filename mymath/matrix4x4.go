package mymath

import "math"

// import "fmt"

type Matrix4x4 [4][4]float64

func MakeIdentityM() Matrix4x4 {
	var m Matrix4x4

	m[0][0] = 1.0
	m[1][1] = 1.0
	m[2][2] = 1.0
	m[3][3] = 1.0

	return m
}

func MulMatrices(m1, m2 Matrix4x4) Matrix4x4 {
	var m Matrix4x4

	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			m[r][c] = m1[r][0]*m2[0][c] + m1[r][1]*m2[1][c] + m1[r][2]*m2[2][c] + m1[r][3]*m2[3][c]
		}
	}

	return m
}

func (m Matrix4x4) MulNumber(k float64) Matrix4x4 {
	var res Matrix4x4

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			res[i][j] = m[i][j] * k
		}
	}

	return res
}

func (m Matrix4x4) transpose() Matrix4x4 {
	var res Matrix4x4

	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			res[j][i] = m[i][j]
		}
	}

	return res
}

func (m Matrix4x4) submatrixDet(r, c int) float64 {

	for i := r; i < 3; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = m[i+1][j]
		}
	}

	for i := 0; i < 4; i++ {
		for j := c; j < 3; j++ {
			m[i][j] = m[i][j+1]
		}
	}

	return m[0][0]*m[1][1]*m[2][2] +
		m[0][1]*m[1][2]*m[2][0] +
		m[0][2]*m[1][0]*m[2][1] -
		m[0][2]*m[1][1]*m[2][0] -
		m[0][1]*m[1][0]*m[2][2] -
		m[0][0]*m[1][2]*m[2][1]

}

func (m Matrix4x4) Determinant() float64 {
	det := m[0][0]*m.submatrixDet(0, 0) -
		m[1][0]*m.submatrixDet(1, 0) +
		m[2][0]*m.submatrixDet(2, 0) -
		m[3][0]*m.submatrixDet(3, 0)

	return det
}

func (m Matrix4x4) Inverse() (Matrix4x4, bool) {
	var res, adjugate Matrix4x4

	det := m.Determinant()

	if det == 0 {
		return res, false
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			adjugate[i][j] = m.submatrixDet(j, i)

			if (i+j)%2 != 0 {
				adjugate[i][j] *= -1
			}
		}
	}

	res = adjugate.MulNumber(1.0 / det)

	// fmt.Println(MulMatrices(m, res))

	return res, true
}


func MakeFovProjectionM(fov, ar, n, f float64) Matrix4x4 {
	var proj Matrix4x4

	fovRad := fov * math.Pi / 180.0
	w := 1.0 / math.Tan(fovRad/2.0)
	h := w * ar

	proj[0][0] = h
	proj[1][1] = w
	proj[2][2] = f / (f - n)
	proj[3][2] = -n * f / (f - n)
	proj[2][3] = 1.0

	return proj
}

func MakeTranslationM(x, y, z float64) Matrix4x4 {
	tr := MakeIdentityM()
	tr[3][0] = x
	tr[3][1] = y
	tr[3][2] = z

	return tr
}


func MakePointAtM(pos, target, up Vec3) Matrix4x4 {
	newForward := Vec3Diff(target, pos)
	newForward.Normalize()

	a := Vec3Mul(newForward, up.DotProduct(newForward))
	newUp := Vec3Diff(up, a)
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
