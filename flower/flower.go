package flower

import (
	"image/color"
	"math"
	"tulip/mymath"
	"tulip/scene"
)

type Tulip struct {
	scene.CompositeModel

	Petals [6]int
	Leaves [2]int
	Stem   int

	clr color.NRGBA

	stemLen float64

	pos mymath.Vec3

	stage1 mymath.BezierCurve
	stage2 mymath.BezierCurve
}

func MakeStem(h, r float64, n, k int, clr color.NRGBA) scene.Model {
	var stem scene.Model

	dy := h / float64(n)
	df := math.Pi * 2 / float64(k)

	for i := 0; i < n+1; i++ {
		y := float64(i) * dy
		x := r
		z := 0.0

		for f := math.Pi / 2; f < 5*math.Pi/2; f += df {
			x = math.Sin(f) * r
			z = math.Cos(f) * r

			p := mymath.MakeVec4(x, y, z)
			var vertex scene.Vertex
			vertex.Point = p

			center := mymath.MakeVec3(0, y, 0)
			vertex.Normal = mymath.Vec3Diff(center, p.Vec3).ToVec4()

			stem.Vertices = append(stem.Vertices, vertex)
		}
	}

	for i := 0; i < n*k; i++ {
		// p1 := scene.Polygon{i, i/k*k + (i+1)%k, i/k*k + (i+1)%k + k, clr}
		// p2 := scene.Polygon{i, i + k, i/k*k + (i+1)%k + k, clr}

		stem.AddPolygon(i, i/k*k + (i+1)%k, i/k*k + (i+1)%k + k, clr)
		stem.AddPolygon(i, i + k, i/k*k + (i+1)%k + k, clr)
		// stem.Polygons = append(stem.Polygons, p1)
		// stem.Polygons = append(stem.Polygons, p2)
	}


	return stem

}

func MakeLeaf(clr color.NRGBA) scene.Model {
	var (
		half1, half2, half3, half4, leaf scene.Model
	)

	half1.AddPoint(mymath.MakeVec4(0, 0, 0))
	half1.AddPoint(mymath.MakeVec4(5, 10, 0))
	half1.AddPoint(mymath.MakeVec4(0, 30, 0))

	half1.Vertices[0].Normal = mymath.MakeVec4(0, 0, -1)
	half1.Vertices[1].Normal = mymath.MakeVec4(0, 0, -1)
	half1.Vertices[2].Normal = mymath.MakeVec4(0, 0, -1)

	// half1.Polygons = append(half1.Polygons, scene.Polygon{0, 1, 2, clr})
	half1.AddPolygon(0, 1, 2, clr)

	half2.AddPoint(mymath.MakeVec4(0, 0, 0))
	half2.AddPoint(mymath.MakeVec4(5, 10, 0))
	half2.AddPoint(mymath.MakeVec4(0, 30, 0))

	half2.Vertices[0].Normal = mymath.MakeVec4(0, 0, 1)
	half2.Vertices[1].Normal = mymath.MakeVec4(0, 0, 1)
	half2.Vertices[2].Normal = mymath.MakeVec4(0, 0, 1)

	// half2.Polygons = append(half1.Polygons, scene.Polygon{0, 1, 2, clr})
	half2.AddPolygon(0, 1, 2, clr)

	half2.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 120, 0))

	half1.Vertices[0].Normal.Add(half2.Vertices[0].Normal)
	half1.Vertices[2].Normal.Add(half2.Vertices[2].Normal)

	half1.Vertices[0].Normal.Mul(0.5)
	half1.Vertices[2].Normal.Mul(0.5)

	leaf.Vertices = append(leaf.Vertices, half1.Vertices...)
	leaf.Vertices = append(leaf.Vertices, half2.Vertices[1])

	// leaf.Polygons = append(leaf.Polygons, scene.Polygon{0, 1, 2, clr}, scene.Polygon{0, 2, 3, clr})
	leaf.AddPolygon(0, 1, 2, clr)
	leaf.AddPolygon(0, 2, 3, clr)

	half3.AddPoint(mymath.MakeVec4(0, 0, 0))
	half3.AddPoint(mymath.MakeVec4(5, 10, 0))
	half3.AddPoint(mymath.MakeVec4(0, 30, 0))
	half3.AddPoint(mymath.MakeVec4(0, 10, -2))

	half3.Vertices[0].Normal = mymath.MakeVec4(0, 0, 1)
	half3.Vertices[1].Normal = mymath.MakeVec4(0, 0, 1)
	half3.Vertices[2].Normal = mymath.MakeVec4(0, 0, 1)
	half3.Vertices[3].Normal = mymath.MakeVec4(0, 0, 1)

	// half3.Polygons = append(half3.Polygons, scene.Polygon{0, 1, 3, clr})
	// half3.Polygons = append(half3.Polygons, scene.Polygon{1, 2, 3, clr})
	half3.AddPolygon(0, 1, 3, clr)
	half3.AddPolygon(0, 2, 3, clr)


	half4.AddPoint(mymath.MakeVec4(0, 0, 0))
	half4.AddPoint(mymath.MakeVec4(5, 10, 0))
	half4.AddPoint(mymath.MakeVec4(0, 30, 0))
	half4.AddPoint(mymath.MakeVec4(0, 10, -2))

	half4.Vertices[0].Normal = mymath.MakeVec4(0, 0, -1)
	half4.Vertices[1].Normal = mymath.MakeVec4(0, 0, -1)
	half4.Vertices[2].Normal = mymath.MakeVec4(0, 0, -1)
	half4.Vertices[3].Normal = mymath.MakeVec4(0, 0, -1)

	// half4.Polygons = append(half4.Polygons, scene.Polygon{0, 1, 2, clr})
	// half4.Polygons = append(half4.Polygons, scene.Polygon{1, 2, 3, clr})

	half4.AddPolygon(0, 1, 2, clr)
	half4.AddPolygon(1, 2, 3, clr)

	half4.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 120, 0))

	half3.Vertices[0].Normal.Add(half4.Vertices[0].Normal)
	half3.Vertices[2].Normal.Add(half4.Vertices[2].Normal)
	half3.Vertices[3].Normal.Add(half4.Vertices[3].Normal)

	half3.Vertices[0].Normal.Mul(0.5)
	half3.Vertices[2].Normal.Mul(0.5)
	half3.Vertices[3].Normal.Mul(0.5)

	leaf.Vertices = append(leaf.Vertices, half3.Vertices...)
	leaf.Vertices = append(leaf.Vertices, half4.Vertices[1])

	// leaf.Polygons = append(leaf.Polygons, scene.Polygon{4, 5, 7, clr}, scene.Polygon{5, 6, 7, clr}, scene.Polygon{4, 7, 8, clr}, scene.Polygon{7, 2, 8, clr})
	leaf.AddPolygon(4, 5, 7, clr)
	leaf.AddPolygon(5, 6, 7, clr)
	leaf.AddPolygon(4, 7, 8 ,clr)
	leaf.AddPolygon(7, 2, 8 ,clr)

	leaf.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, -60, 0))

	leaf.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 0, -30))

	return leaf
}

func MakePetal(curve mymath.BezierCurve, m, n int, clr color.NRGBA) scene.Model {
	var (
		petal   scene.Model
		surface mymath.BicubicBezierSurface
	)

	surface[0] = curve
	surface[1] = curve
	surface[2] = curve
	surface[3] = curve

	surface[1][1].Rotate(mymath.MakeVec3(0, surface[1][1].Y, 0),
		mymath.MakeVec3(0, 20, 0))
	surface[1][2].Rotate(mymath.MakeVec3(0, surface[1][2].Y, 0),
		mymath.MakeVec3(0, 20, 0))

	surface[2][1].Rotate(mymath.MakeVec3(0, surface[2][1].Y, 0),
		mymath.MakeVec3(0, 40, 0))
	surface[2][2].Rotate(mymath.MakeVec3(0, surface[2][2].Y, 0),
		mymath.MakeVec3(0, 40, 0))

	surface[3][1].Rotate(mymath.MakeVec3(0, surface[3][1].Y, 0),
		mymath.MakeVec3(0, 60, 0))
	surface[3][2].Rotate(mymath.MakeVec3(0, surface[3][2].Y, 0),
		mymath.MakeVec3(0, 60, 0))

	for j := 0; j <= n; j++ {
		v := float64(j) / float64(n)
		for i := 0; i <= m; i++ {
			u := float64(i) / float64(m)
			point := surface.GetPoint(float64(i)/float64(m), float64(j)/float64(n))

			var vertex scene.Vertex
			vertex.Point = point.ToVec4()

			dU := surface.DUBezier(u, v)
			dV := surface.DVBezier(u, v)

			vertex.Normal = dU.CrossProduct(dV).ToVec4()

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
		//petal.Polygons = append(petal.Polygons, scene.Polygon{n, n + 1, n + 2, clr})
		petal.AddPolygon(n, n+1, n+2, clr)
	}

	return petal
}

func (t *Tulip) MakePetals(stage mymath.BezierCurve) {
	for i := range t.Petals {

		petal := MakePetal(stage, 10, 10, t.clr)

		t.Add(&petal)
		t.Petals[i] = t.Size() - 1
		t.Components[t.Petals[i]].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 120*float64(i), 0))

		if i > 2 {
			t.Components[t.Petals[i]].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 60, 0))
		}

		t.Components[t.Petals[i]].Move(mymath.MakeVec3(t.pos.X, t.pos.Y+t.stemLen, t.pos.Z))
	}

}

func NewTulip(clr color.NRGBA, pos mymath.Vec3, stage int, k float64) *Tulip {
	t := new(Tulip)

	t.pos = pos

	t.clr = clr

	t.stage1[0] = mymath.MakeVec3(0, 0, 0)
	t.stage1[1] = mymath.MakeVec3(6, 0, 0)
	t.stage1[2] = mymath.MakeVec3(4, 10, 0)
	t.stage1[3] = mymath.MakeVec3(0, 10, 0)

	t.stage2[0] = mymath.MakeVec3(0, 0, 0)
	t.stage2[1] = mymath.MakeVec3(6, 0, 0)
	t.stage2[2] = mymath.MakeVec3(6, 8, 0)
	t.stage2[3] = mymath.MakeVec3(5, 10.907, 0)

	for i := range t.stage1 {
		t.stage1[i].Scale(mymath.MakeVec3(0, 0, 0), k)
		t.stage2[i].Scale(mymath.MakeVec3(0, 0, 0), k)
	}

	t.stemLen = 30
	t.stemLen *= k

	if stage == 1 {
		t.MakePetals(t.stage1)
	} else {
		t.MakePetals(t.stage2)
	}

	// Stem
	stem := MakeStem(t.stemLen, 0.5 * k, 10, 10, color.NRGBA{0, 200, 25, 255})

	t.Add(&stem)
	t.Stem = t.Size() - 1

	t.Components[t.Stem].Move(mymath.MakeVec3(pos.X, pos.Y, pos.Z))

	leaf1 := MakeLeaf(color.NRGBA{0, 200, 25, 255})
	leaf1.Scale(mymath.MakeVec3(0, 0, 0), k)

	t.Add(&leaf1)
	t.Leaves[0] = t.Size() - 1

	t.Components[t.Leaves[0]].Move(mymath.MakeVec3(pos.X, pos.Y, pos.Z))

	leaf2 := MakeLeaf(color.NRGBA{0, 200, 25, 255})
	leaf2.Scale(mymath.MakeVec3(0, 0, 0), k)

	leaf2.Reflect(true, false, false)

	t.Add(&leaf2)
	t.Leaves[1] = t.Size() - 1

	t.Components[t.Leaves[1]].Move(mymath.MakeVec3(pos.X, pos.Y, pos.Z))

	return t
}

func (t *Tulip) Rotate(center, angles mymath.Vec3) {
	t.CompositeModel.Rotate(center, angles)

	for i := 0; i < 4; i++ {
		t.stage1[i].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 0, angles.Z))
		t.stage2[i].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 0, angles.Z))
	}
}

func (t *Tulip) Scale(center mymath.Vec3, k float64) {
	t.CompositeModel.Scale(center, k)
	t.stemLen *= k

	for i := 0; i < 4; i++ {
		t.stage1[i].Scale(mymath.MakeVec3(0, 0, 0), k)
		t.stage2[i].Scale(mymath.MakeVec3(0, 0, 0), k)
	}
}

func (t *Tulip) ChangePetals(stage mymath.BezierCurve) {

	var clrs [6]color.NRGBA
	clrs[0] = color.NRGBA{255, 0, 0, 255}
	clrs[1] = color.NRGBA{0, 255, 0, 255}
	clrs[2] = color.NRGBA{0, 0, 255, 255}
	clrs[3] = color.NRGBA{255, 255, 0, 255}
	clrs[4] = color.NRGBA{255, 0, 255, 255}
	clrs[5] = color.NRGBA{0, 255, 255, 255}

	for i := range t.Petals {

		petal := MakePetal(stage, 10, 10, t.clr)

		t.Components[t.Petals[i]] = &petal

		if i > 2 {
			//t.Components[t.Petals[i]].Move(mymath.MakeVec3(1, 0, 0))
			t.Components[t.Petals[i]].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 60, 0))
		}

		t.Components[t.Petals[i]].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 120*float64(i), 0))

		t.Components[t.Petals[i]].Move(mymath.MakeVec3(t.pos.X, t.pos.Y+t.stemLen, t.pos.Z))
	}

}

func (t *Tulip) ChangePetalsNew(stageIn, stageOut mymath.BezierCurve) {

	var clrs [6]color.NRGBA
	clrs[0] = color.NRGBA{255, 0, 0, 255}
	clrs[1] = color.NRGBA{0, 255, 0, 255}
	clrs[2] = color.NRGBA{0, 0, 255, 255}
	clrs[3] = color.NRGBA{255, 255, 0, 255}
	clrs[4] = color.NRGBA{255, 0, 255, 255}
	clrs[5] = color.NRGBA{0, 255, 255, 255}

	for i := 0; i < 3; i++ {

		petal := MakePetal(stageOut, 10, 10, t.clr)

		t.Components[t.Petals[i]] = &petal
		//t.Components[t.Petals[i]].Move(mymath.MakeVec3(1, 0, 0))

		t.Components[t.Petals[i]].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 120*float64(i), 0))

		t.Components[t.Petals[i]].Move(mymath.MakeVec3(t.pos.X, t.pos.Y+t.stemLen, t.pos.Z))
	}

	for i := 3; i < 6; i++ {

		petal := MakePetal(stageIn, 10, 10, t.clr)

		t.Components[t.Petals[i]] = &petal

		t.Components[t.Petals[i]].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 120*float64(i)+60, 0))

		// t.Components[t.Petals[i]].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, , 0))

		t.Components[t.Petals[i]].Move(mymath.MakeVec3(t.pos.X, t.pos.Y+t.stemLen, t.pos.Z))
	}

}

func (t Tulip) interpolateStagePoint(i int, k float64) mymath.Vec3 {

	var p mymath.Vec3

	p.X = t.stage1[i].X + (k * (t.stage2[i].X - t.stage1[i].X))
	p.Y = t.stage1[i].Y + (k * (t.stage2[i].Y - t.stage1[i].Y))
	p.Z = t.stage1[i].Z + (k * (t.stage2[i].Z - t.stage1[i].Z))

	return p
}

func (t *Tulip) OpenPetals(k float64) {
	var stage mymath.BezierCurve

	stage[0] = t.interpolateStagePoint(0, k)
	stage[1] = t.interpolateStagePoint(1, k)
	stage[2] = t.interpolateStagePoint(2, k)
	stage[3] = t.interpolateStagePoint(3, k)

	t.ChangePetals(stage)
}

func (t *Tulip) OpenPetalsNew(k float64) {
	var stageOut, stageIn mymath.BezierCurve

	kIn := k - 0.15
	if kIn < 0.0 {
		kIn = 0.0
	}

	stageOut[0] = t.interpolateStagePoint(0, k)
	stageOut[1] = t.interpolateStagePoint(1, k)
	stageOut[2] = t.interpolateStagePoint(2, k)
	stageOut[3] = t.interpolateStagePoint(3, k)

	stageIn[0] = t.interpolateStagePoint(0, kIn)
	stageIn[1] = t.interpolateStagePoint(1, kIn)
	stageIn[2] = t.interpolateStagePoint(2, kIn)
	stageIn[3] = t.interpolateStagePoint(3, kIn)

	t.ChangePetalsNew(stageIn, stageOut)
}

func (t *Tulip) Animate(k float64) {
	t.OpenPetals(k)
}
