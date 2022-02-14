package mymath

type BezierCurve [4]Vec3

type BicubicBezierSurface [4]BezierCurve

func (curve BezierCurve) GetPoint(t float64) Vec3 {
	a := (1 - t) * (1 - t) * (1 - t)
	b := 3 * (1 - t) * (1 - t) * t
	c := 3 * (1 - t) * t * t
	d := t * t * t

	x := a*curve[0].X + b*curve[1].X + c*curve[2].X + d*curve[3].X
	y := a*curve[0].Y + b*curve[1].Y + c*curve[2].Y + d*curve[3].Y
	z := a*curve[0].Z + b*curve[1].Z + c*curve[2].Z + d*curve[3].Z

	return Vec3{x, y, z}
}

func (surface BicubicBezierSurface) GetPoint(u, v float64) Vec3 {
	var curve BezierCurve

	curve[0] = surface[0].GetPoint(u)
	curve[1] = surface[1].GetPoint(u)
	curve[2] = surface[2].GetPoint(u)
	curve[3] = surface[3].GetPoint(u)

	return curve.GetPoint(v)
}

func (surface BicubicBezierSurface) DUBezier(u, v float64) Vec3 {
	var (
		curve, vCurve BezierCurve
		vec           Vec3
	)

	for i := 0; i < 4; i++ {
		curve[0] = surface[0][i]
		curve[1] = surface[1][i]
		curve[2] = surface[2][i]
		curve[3] = surface[3][i]

		vCurve[i] = curve.GetPoint(v)
	}

	vec = Vec3Mul(vCurve[0], -3*(1-u)*(1-u))
	vec.Add(Vec3Mul(vCurve[1], 3*(1-u)*(1-u)-6*u*(1-u)))
	vec.Add(Vec3Mul(vCurve[2], 6*u*(1-u)-3*u*u))
	vec.Add(Vec3Mul(vCurve[3], 3*u*u))

	return vec
}

func (surface BicubicBezierSurface) DVBezier(u, v float64) Vec3 {
	var (
		uCurve BezierCurve
		vec    Vec3
	)

	for i := 0; i < 4; i++ {
		uCurve[i] = surface[i].GetPoint(u)
	}

	vec = Vec3Mul(uCurve[0], -3*(1-v)*(1-v))
	vec.Add(Vec3Mul(uCurve[1], 3*(1-v)*(1-v)-6*v*(1-v)))
	vec.Add(Vec3Mul(uCurve[2], 6*v*(1-v)-3*v*v))
	vec.Add(Vec3Mul(uCurve[3], 3*v*v))

	return vec
}
