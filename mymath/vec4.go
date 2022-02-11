package mymath

type Vec4 struct {
	Vec3
	w float64
}

func MakeVec4(x, y, z float64, w... float64) Vec4 {
	var vec4 Vec4
	vec4.Vec3 = MakeVec3(x, y, z)

	if len(w) != 0 {
		vec4.w = w[0]
	}

	return vec4
}

func Vec3ToVec4(vec3 Vec3, w... float64) Vec4 {
	var vec4 Vec4
	vec4.Vec3 = vec3

	if len(w) != 0 {
		vec4.w = w[0]
	}

	return vec4
}

func (vec *Vec4) DivW() {
	if vec.w != 0 {
		vec.Div(vec.w)
	}
}

func mulVecMat(vec Vec4, m Matrix4x4) Vec4 {
	var res Vec4

	res.X = vec.X*m[0][0] + vec.Y*m[1][0] + vec.Z*m[2][0] + m[3][0]
	res.Y = vec.X*m[0][1] + vec.Y*m[1][1] + vec.Z*m[2][1] + m[3][1]
	res.Z = vec.X*m[0][2] + vec.Y*m[1][2] + vec.Z*m[2][2] + m[3][2]
	res.w = vec.X*m[0][3] + vec.Y*m[1][3] + vec.Z*m[2][3] + m[3][3]

	// res.X = vec.X*m[0][0] + vec.Y*m[0][1] + vec.Z*m[0][2] + m[0][3]
	// res.Y = vec.X*m[1][0] + vec.Y*m[1][1] + vec.Z*m[1][2] + m[1][3]
	// res.Z = vec.X*m[2][0] + vec.Y*m[2][1] + vec.Z*m[2][2] + m[2][3]
	// res.w = vec.X*m[3][0] + vec.Y*m[3][1] + vec.Z*m[3][2] + m[3][3]

	return res
}
