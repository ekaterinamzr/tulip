package bezier

import (
	"tulip/object"
)

type BezierCurve [4]object.Vertex

type BicubicBezierSurface [4]BezierCurve

func (curve BezierCurve) GetPoint(t float64) object.Vertex {
	a := (1 - t) * (1 - t) * (1 - t)
    b := 3 * (1 - t) * (1 - t) * t
    c := 3 * (1 - t) * t * t
    d := t * t * t

    x := a * curve[0].X + b * curve[1].X + c * curve[2].X + d * curve[3].X
    y := a * curve[0].Y + b * curve[1].Y + c * curve[2].Y + d * curve[3].Y
    z := a * curve[0].Z + b * curve[1].Z + c * curve[2].Z + d * curve[3].Z

	return object.Vertex{x, y, z}
}

func (surface BicubicBezierSurface) GetPoint(u, v float64) object.Vertex {
	var curve BezierCurve

	curve[0] = surface[0].GetPoint(u)
	curve[1] = surface[1].GetPoint(u)
	curve[2] = surface[2].GetPoint(u)
	curve[3] = surface[3].GetPoint(u)

	return curve.GetPoint(v)
}

