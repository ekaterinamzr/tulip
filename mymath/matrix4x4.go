package mymath

type Matrix4x4 [4][4] float64

func MulVectorMatrix(vec Vector3d, m Matrix4x4) Vector3d{
	var res Vector3d

	res.X = vec.X * m[0][0] + vec.Y * m[1][0] + vec.Z * m[2][0] + m[3][0]
	res.Y = vec.X * m[0][1] + vec.Y * m[1][1] + vec.Z * m[2][1] + m[3][1]
	res.Z = vec.X * m[0][2] + vec.Y * m[1][2] + vec.Z * m[2][2] + m[3][2]
	k := vec.X * m[0][3] + vec.Y * m[1][3] + vec.Z * m[2][3] + m[3][3]
	
	if (k != 0.0) {
		res.Div(k)
	}

	return res
}