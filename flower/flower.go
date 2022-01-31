package flower

import (
	"image/color"
	"math"
	"tulip/mymath"
	"tulip/object"
)

type Tulip struct {
	object.CompositeModel

	Petals [6]int
	Leaves [2]int
	Stem   int

	clr color.NRGBA

	stemLen float64

	pos mymath.Vector3d

	stage1 mymath.BezierCurve
	stage2 mymath.BezierCurve
}

type TulipOld struct {
	Petals [6]object.Model
	Leafs  [2]object.Model
	Stem   object.Model

	clr color.NRGBA

	stage1 mymath.BezierCurve
	stage2 mymath.BezierCurve
}

func TestPetal() object.Model {
	var curve mymath.BezierCurve
	curve[0] = mymath.Vector3d{0, 0, 0}
	curve[1] = mymath.Vector3d{6, 0, 0}
	curve[2] = mymath.Vector3d{4, 10, 0}
	curve[3] = mymath.Vector3d{0, 10, 0}

	for i := range curve {
		curve[i].Scale(mymath.Vector3d{0, 0, 0}, 10)
	}

	petal := MakePetal(curve, 5, 5, color.NRGBA{255, 0, 0, 255})

	petal.Move(mymath.Vector3d{100, 100, 0})

	return petal
}

func MakePetal(curve mymath.BezierCurve, m, n int, clr color.NRGBA) object.Model {
	var (
		petal   object.Model
		surface mymath.BicubicBezierSurface
	)

	surface[0] = curve
	surface[1] = curve
	surface[2] = curve
	surface[3] = curve

	surface[1][1].Rotate(mymath.Vector3d{0, surface[1][1].Y, 0},
		mymath.Vector3d{0, 20, 0})
	surface[1][2].Rotate(mymath.Vector3d{0, surface[1][2].Y, 0},
		mymath.Vector3d{0, 20, 0})

	surface[2][1].Rotate(mymath.Vector3d{0, surface[2][1].Y, 0},
		mymath.Vector3d{0, 40, 0})
	surface[2][2].Rotate(mymath.Vector3d{0, surface[2][2].Y, 0},
		mymath.Vector3d{0, 40, 0})

	surface[3][1].Rotate(mymath.Vector3d{0, surface[3][1].Y, 0},
		mymath.Vector3d{0, 60, 0})
	surface[3][2].Rotate(mymath.Vector3d{0, surface[3][2].Y, 0},
		mymath.Vector3d{0, 60, 0})

	for j := 0; j <= n; j++ {
		v := float64(j) / float64(n)
		for i := 0; i <= m; i++ {
			u := float64(i) / float64(m)
			point := surface.GetPoint(float64(i)/float64(m), float64(j)/float64(n))

			var vertex object.Vertex
			vertex.Point = point

			dU := surface.DUBezier(u, v)
			dV := surface.DVBezier(u, v)

			vertex.Normal = dU.CrossProduct(dV)

			petal.Vertices = append(petal.Vertices, vertex)
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			v1 := (m+1)*j + i
			v2 := (m+1)*(j+1) + i
			v3 := (m+1)*(j+1) + i + 1
			v4 := (m+1)*j + i + i

			petal.AddPolygon(v1, v2, v3, clr)
			petal.AddPolygon(v3, v4, v1, clr)
		}
	}

	for i := range petal.Polygons {
		v1 := petal.Vertices[petal.Polygons[i].V1]
		v2 := petal.Vertices[petal.Polygons[i].V2]
		v3 := petal.Vertices[petal.Polygons[i].V3]

		v1.Reflect(false, false, true)
		v2.Reflect(false, false, true)
		v3.Reflect(false, false, true)

		n := len(petal.Vertices)

		petal.Vertices = append(petal.Vertices, v1, v2, v3)
		petal.Polygons = append(petal.Polygons, object.Polygon{n, n + 1, n + 2, clr})
	}

	return petal
}

func MakeStem(h, r float64, n, k int, clr color.NRGBA) object.Model {
	var stem object.Model

	dy := h / float64(n)
	df := math.Pi * 2 / float64(k)

	for i := 0; i < n+1; i++ {
		y := float64(i) * dy
		x := r
		z := 0.0

		for f := math.Pi / 2; f < 5*math.Pi/2; f += df {
			x = math.Sin(f) * r
			z = math.Cos(f) * r

			p := mymath.Vector3d{x, y, z}
			var vertex object.Vertex
			vertex.Point = p

			center := mymath.MakeVector3d(0, y, 0)
			vertex.Normal = mymath.Vector3dDiff(center, p)

			stem.Vertices = append(stem.Vertices, vertex)
		}
	}

	for i := 0; i < n*k; i++ {
		p1 := object.Polygon{i, i/k*k + (i+1)%k, i/k*k + (i+1)%k + k, clr}
		p2 := object.Polygon{i, i + k, i/k*k + (i+1)%k + k, clr}

		stem.Polygons = append(stem.Polygons, p1)
		stem.Polygons = append(stem.Polygons, p2)
	}

	return stem

}

func MakeLeaf(clr color.NRGBA) object.Model {
	var (
		half1, half2, half3, half4, leaf object.Model
	)

	half1.AddPoint(mymath.Vector3d{0, 0, 0})
	half1.AddPoint(mymath.Vector3d{5, 10, 0})
	half1.AddPoint(mymath.Vector3d{0, 30, 0})

	half1.Vertices[0].Normal = mymath.MakeVector3d(0, 0, -1)
	half1.Vertices[1].Normal = mymath.MakeVector3d(0, 0, -1)
	half1.Vertices[2].Normal = mymath.MakeVector3d(0, 0, -1)

	half1.Polygons = append(half1.Polygons, object.Polygon{0, 1, 2, clr})

	half2.AddPoint(mymath.Vector3d{0, 0, 0})
	half2.AddPoint(mymath.Vector3d{5, 10, 0})
	half2.AddPoint(mymath.Vector3d{0, 30, 0})

	half2.Vertices[0].Normal = mymath.MakeVector3d(0, 0, 1)
	half2.Vertices[1].Normal = mymath.MakeVector3d(0, 0, 1)
	half2.Vertices[2].Normal = mymath.MakeVector3d(0, 0, 1)

	half2.Polygons = append(half1.Polygons, object.Polygon{0, 1, 2, clr})

	half2.Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, 120, 0})

	half1.Vertices[0].Normal.Add(half2.Vertices[0].Normal)
	half1.Vertices[2].Normal.Add(half2.Vertices[2].Normal)

	half1.Vertices[0].Normal.Mul(0.5)
	half1.Vertices[2].Normal.Mul(0.5)

	leaf.Vertices = append(leaf.Vertices, half1.Vertices...)
	leaf.Vertices = append(leaf.Vertices, half2.Vertices[1])

	leaf.Polygons = append(leaf.Polygons, object.Polygon{0, 1, 2, clr}, object.Polygon{0, 2, 3, clr})

	half3.AddPoint(mymath.Vector3d{0, 0, 0})
	half3.AddPoint(mymath.Vector3d{5, 10, 0})
	half3.AddPoint(mymath.Vector3d{0, 30, 0})
	half3.AddPoint(mymath.Vector3d{0, 10, -2})

	half3.Vertices[0].Normal = mymath.MakeVector3d(0, 0, 1)
	half3.Vertices[1].Normal = mymath.MakeVector3d(0, 0, 1)
	half3.Vertices[2].Normal = mymath.MakeVector3d(0, 0, 1)
	half3.Vertices[3].Normal = mymath.MakeVector3d(0, 0, 1)

	half3.Polygons = append(half3.Polygons, object.Polygon{0, 1, 3, clr})
	half3.Polygons = append(half3.Polygons, object.Polygon{1, 2, 3, clr})

	half4.AddPoint(mymath.Vector3d{0, 0, 0})
	half4.AddPoint(mymath.Vector3d{5, 10, 0})
	half4.AddPoint(mymath.Vector3d{0, 30, 0})
	half4.AddPoint(mymath.Vector3d{0, 10, -2})

	half4.Vertices[0].Normal = mymath.MakeVector3d(0, 0, -1)
	half4.Vertices[1].Normal = mymath.MakeVector3d(0, 0, -1)
	half4.Vertices[2].Normal = mymath.MakeVector3d(0, 0, -1)
	half4.Vertices[3].Normal = mymath.MakeVector3d(0, 0, -1)

	half4.Polygons = append(half4.Polygons, object.Polygon{0, 1, 2, clr})
	half4.Polygons = append(half4.Polygons, object.Polygon{1, 2, 3, clr})

	half4.Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, 120, 0})

	half3.Vertices[0].Normal.Add(half4.Vertices[0].Normal)
	half3.Vertices[2].Normal.Add(half4.Vertices[2].Normal)
	half3.Vertices[3].Normal.Add(half4.Vertices[3].Normal)

	half3.Vertices[0].Normal.Mul(0.5)
	half3.Vertices[2].Normal.Mul(0.5)
	half3.Vertices[3].Normal.Mul(0.5)

	leaf.Vertices = append(leaf.Vertices, half3.Vertices...)
	leaf.Vertices = append(leaf.Vertices, half4.Vertices[1])

	leaf.Polygons = append(leaf.Polygons, object.Polygon{4, 5, 7, clr}, object.Polygon{5, 6, 7, clr}, object.Polygon{4, 7, 8, clr}, object.Polygon{7, 2, 8, clr})

	leaf.Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, -60, 0})

	leaf.Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, 0, -30})

	return leaf
}

func (t *Tulip) Rotate(center, angles mymath.Vector3d) {
	t.CompositeModel.Rotate(center, angles)

	for i := 0; i < 4; i++ {
		t.stage1[i].Rotate(mymath.MakeVector3d(0, 0, 0), mymath.MakeVector3d(0, 0, angles.Z))
		t.stage2[i].Rotate(mymath.MakeVector3d(0, 0, 0), mymath.MakeVector3d(0, 0, angles.Z))
	}
}

func (t *Tulip) Scale(center mymath.Vector3d, k float64) {
	t.CompositeModel.Scale(center, k)
	t.stemLen *= k

	for i := 0; i < 4; i++ {
		t.stage1[i].Scale(mymath.MakeVector3d(0, 0, 0), k)
		t.stage2[i].Scale(mymath.MakeVector3d(0, 0, 0), k)
	}
}

func (t *Tulip) MakePetals(stage mymath.BezierCurve) {
	for i := range t.Petals {

		petal := MakePetal(stage, 20, 20, t.clr)

		t.Add(&petal)
		t.Petals[i] = t.Size() - 1
		t.Components[t.Petals[i]].Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, 120 * float64(i), 0})

		if i > 2 {
			t.Components[t.Petals[i]].Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, 60, 0})
		}

		t.Components[t.Petals[i]].Move(mymath.Vector3d{t.pos.X, t.pos.Y + t.stemLen, t.pos.Z})
	}

}

func (t *Tulip) ChangePetals(stage mymath.BezierCurve) {
	for i := range t.Petals {

		petal := MakePetal(stage, 20, 20, t.clr)

		t.Components[t.Petals[i]] = &petal

		t.Components[t.Petals[i]].Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, 120 * float64(i), 0})

		if i > 2 {
			t.Components[t.Petals[i]].Rotate(mymath.Vector3d{0, 0, 0}, mymath.Vector3d{0, 60, 0})
		}

		t.Components[t.Petals[i]].Move(mymath.Vector3d{t.pos.X, t.pos.Y + t.stemLen, t.pos.Z})
	}

}

func (t Tulip) interpolateStagePoint(i int, k float64) mymath.Vector3d {

	var p mymath.Vector3d

	p.X = t.stage1[i].X + (k * (t.stage2[i].X - t.stage1[i].X))
	p.Y = t.stage1[i].Y + (k * (t.stage2[i].Y - t.stage1[i].Y))
	p.Z = t.stage1[i].Z + (k * (t.stage2[i].Z - t.stage1[i].Z))

	return p
}

func (t *Tulip) Animate(k float64) {
	t.OpenPetals(k)
}

func (t *Tulip) OpenPetals(k float64) {
	var stage mymath.BezierCurve

	stage[0] = t.interpolateStagePoint(0, k)
	stage[1] = t.interpolateStagePoint(1, k)
	stage[2] = t.interpolateStagePoint(2, k)
	stage[3] = t.interpolateStagePoint(3, k)

	t.ChangePetals(stage)
}

func NewTulip(pos mymath.Vector3d, stage int, k float64) *Tulip {
	t := new(Tulip)

	t.pos = pos

	t.clr = color.NRGBA{222, 0, 33, 255}

	t.stage1[0] = mymath.Vector3d{0, 0, 0}
	t.stage1[1] = mymath.Vector3d{6, 0, 0}
	t.stage1[2] = mymath.Vector3d{4, 10, 0}
	t.stage1[3] = mymath.Vector3d{0, 10, 0}

	t.stage2[0] = mymath.Vector3d{0, 0, 0}
	t.stage2[1] = mymath.Vector3d{6, 0, 0}
	t.stage2[2] = mymath.Vector3d{6, 8, 0}
	t.stage2[3] = mymath.Vector3d{5, 10.907, 0}

	for i := range t.stage1 {
		t.stage1[i].Scale(mymath.Vector3d{0, 0, 0}, k)
		t.stage2[i].Scale(mymath.Vector3d{0, 0, 0}, k)
	}

	t.stemLen = 30
	t.stemLen *= k

	if stage == 1 {
		t.MakePetals(t.stage1)
	} else {
		t.MakePetals(t.stage2)
	}

	// Stem
	stem := MakeStem(t.stemLen, 0.5*k, 40, 20, color.NRGBA{0, 200, 25, 255})

	t.Add(&stem)
	t.Stem = t.Size() - 1

	t.Components[t.Stem].Move(mymath.Vector3d{pos.X, pos.Y, pos.Z})

	leaf1 := MakeLeaf(color.NRGBA{0, 200, 25, 255})
	leaf1.Scale(mymath.MakeVector3d(0, 0, 0), k)

	t.Add(&leaf1)
	t.Leaves[0] = t.Size() - 1

	t.Components[t.Leaves[0]].Move(mymath.Vector3d{pos.X, pos.Y, pos.Z})

	leaf2 := MakeLeaf(color.NRGBA{0, 200, 25, 255})
	leaf2.Scale(mymath.MakeVector3d(0, 0, 0), k)

	leaf2.Reflect(false, true, false)

	t.Add(&leaf2)
	t.Leaves[1] = t.Size() - 1

	t.Components[t.Leaves[1]].Move(mymath.Vector3d{pos.X, pos.Y, pos.Z})

	return t
}

func TestLeaf() object.Model {

	leaf := MakeLeaf(color.NRGBA{0, 255, 0, 255})

	leaf.Move(mymath.Vector3d{100, 100, 100})

	return leaf

}
