package flower

import (
	"fmt"
	"image/color"
	"math"
	"tulip/bezier"
	"tulip/object"
)

type Tulip struct {
	object.CompositeModel

	Petals [6]int
	Leaves [2]int
	Stem   int

	clr color.NRGBA

	stemLen float64

	pos object.Point

	stage1 bezier.BezierCurve
	stage2 bezier.BezierCurve
}

type TulipOld struct {
	Petals [6]object.Model
	Leafs  [2]object.Model
	Stem   object.Model

	clr color.NRGBA

	stage1 bezier.BezierCurve
	stage2 bezier.BezierCurve
}

func TestPetal() object.Model {
	var curve bezier.BezierCurve
	curve[0] = object.Point{0, 0, 0}
	curve[1] = object.Point{6, 0, 0}
	curve[2] = object.Point{4, 10, 0}
	curve[3] = object.Point{0, 10, 0}

	for i := range curve {
		curve[i].Scale(object.Point{0, 0, 0}, 10)
	}

	petal := MakePetal(curve, 5, 5, color.NRGBA{255, 0, 0, 255})

	//petal.ScaleModel(object.Point{0, 0, 0}, 10)
	petal.Move(object.Point{100, 100, 0})

	return petal
}

func MakePetal(curve bezier.BezierCurve, m, n int, clr color.NRGBA) object.Model {
	var (
		petal   object.Model
		surface bezier.BicubicBezierSurface
	)

	surface[0] = curve
	surface[1] = curve
	surface[2] = curve
	surface[3] = curve

	surface[1][1].Rotate(object.Point{0, surface[1][1].Y, 0},
		object.Point{0, 20, 0})
	surface[1][2].Rotate(object.Point{0, surface[1][2].Y, 0},
		object.Point{0, 20, 0})

	surface[2][1].Rotate(object.Point{0, surface[2][1].Y, 0},
		object.Point{0, 40, 0})
	surface[2][2].Rotate(object.Point{0, surface[2][2].Y, 0},
		object.Point{0, 40, 0})

	surface[3][1].Rotate(object.Point{0, surface[3][1].Y, 0},
		object.Point{0, 60, 0})
	surface[3][2].Rotate(object.Point{0, surface[3][2].Y, 0},
		object.Point{0, 60, 0})

	for j := 0; j <= n; j++ {
		v := float64(j) / float64(n)
		for i := 0; i <= m; i++ {
			u := float64(i) / float64(m)
			point := surface.GetPoint(float64(i)/float64(m), float64(j)/float64(n))

			var vertex object.Vertex
			vertex.Point = point

			dU := surface.DUBezier(u, v)
			dV := surface.DVBezier(u, v)

			//fmt.Println(dU, dV)
			vertex.Normal = dU.Cross(dV)

			petal.Vertices = append(petal.Vertices, vertex)

			//petal.AddPoint(point)
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

	// for i := 0; i < (n+1)*m; i++ {
	// 	//fmt.Println("jopa")
	// 	if i != 0 && (i+1)%(n+1) == 0 {
	// 		continue
	// 	}
	// 	petal.Polygons = append(petal.Polygons, object.Polygon{i, i + 1, i + n + 2, clr})
	// 	petal.Polygons = append(petal.Polygons, object.Polygon{i, i + n + 1, i + n + 2, clr})
	// }

	//petal.CalculateNormals()

	for i := range petal.Polygons {
		v1 := petal.Vertices[petal.Polygons[i].V1]
		v2 := petal.Vertices[petal.Polygons[i].V2]
		v3 := petal.Vertices[petal.Polygons[i].V3]

		v1.Z = -v1.Z
		v2.Z = -v2.Z
		v3.Z = -v3.Z

		v1.Normal.Z = -v1.Normal.Z
		v2.Normal.Z = -v2.Normal.Z
		v3.Normal.Z = -v3.Normal.Z

		n := len(petal.Vertices)

		petal.Vertices = append(petal.Vertices, v1, v2, v3)
		petal.Polygons = append(petal.Polygons, object.Polygon{n, n + 1, n + 2, clr})
	}

	//fmt.Println(petal.Vertices)

	return petal
}

func MakeStem(h, r float64, n, k int, clr color.NRGBA) object.Model {
	var stem object.Model

	dy := h / float64(n)
	df := math.Pi * 2 / float64(k)

	//dz := math.Cos(df)
	//dx := math.Sin(df)

	for i := 0; i < n+1; i++ {
		y := float64(i) * dy
		x := r
		z := 0.0

		//stem.Vertices = append(stem.Vertices, object.Point{x, y, z})

		for f := math.Pi / 2; f < 5*math.Pi/2; f += df {
			x = math.Sin(f) * r
			z = math.Cos(f) * r

			p := object.Point{x, y, z}
			var vertex object.Vertex
			vertex.Point = p

			center := object.MakePoint(0, y, 0)
			vertex.Normal = object.MakeTwoPoints(p, center)

			stem.Vertices = append(stem.Vertices, vertex)
			//fmt.Println(v)

			//stem.AddPoint(p)
		}
	}

	for i := 0; i < n*k; i++ {
		p1 := object.Polygon{i, i/k*k + (i+1)%k, i/k*k + (i+1)%k + k, clr}
		p2 := object.Polygon{i, i + k, i/k*k + (i+1)%k + k, clr}
		//fmt.Println(p1, p2)
		stem.Polygons = append(stem.Polygons, p1)
		stem.Polygons = append(stem.Polygons, p2)
	}

	//fmt.Println(stem.Polygons)

	//stem.CalculateNormals()

	return stem

}

func MakeLeaf(clr color.NRGBA) object.Model {
	var (
		half1, half2, half3, half4, leaf object.Model
		//leaf1 object.CompositeModel
	)

	half1.AddPoint(object.Point{0, 0, 0})
	half1.AddPoint(object.Point{50, 100, 0})
	half1.AddPoint(object.Point{0, 300, 0})

	half1.Vertices[0].Normal = object.Make(0, 0, -1)
	half1.Vertices[1].Normal = object.Make(0, 0, -1)
	half1.Vertices[2].Normal = object.Make(0, 0, -1)

	half1.Polygons = append(half1.Polygons, object.Polygon{0, 1, 2, clr})

	half2.AddPoint(object.Point{0, 0, 0})
	half2.AddPoint(object.Point{50, 100, 0})
	half2.AddPoint(object.Point{0, 300, 0})

	half2.Vertices[0].Normal = object.Make(0, 0, 1)
	half2.Vertices[1].Normal = object.Make(0, 0, 1)
	half2.Vertices[2].Normal = object.Make(0, 0, 1)

	half2.Polygons = append(half1.Polygons, object.Polygon{0, 1, 2, clr})

	//half2.CalculateNormals()

	half2.Rotate(object.Point{0, 0, 0}, object.Point{0, 120, 0})

	half1.Vertices[0].Normal.Add(half2.Vertices[0].Normal)
	half1.Vertices[2].Normal.Add(half2.Vertices[2].Normal)

	half1.Vertices[0].Normal.Mul(0.5)
	half1.Vertices[2].Normal.Mul(0.5)

	leaf.Vertices = append(leaf.Vertices, half1.Vertices...)
	leaf.Vertices = append(leaf.Vertices, half2.Vertices[1])

	//leaf.Vertices = append(leaf.Vertices, half1.Vertices...)
	//leaf.Vertices = append(leaf.Vertices, half2.Vertices...)

	leaf.Polygons = append(leaf.Polygons, object.Polygon{0, 1, 2, clr}, object.Polygon{0, 2, 3, clr})

	half3.AddPoint(object.Point{0, 0, 0})
	half3.AddPoint(object.Point{50, 100, 0})
	half3.AddPoint(object.Point{0, 300, 0})
	half3.AddPoint(object.Point{0, 100, -20})

	half3.Vertices[0].Normal = object.Make(0, 0, 1)
	half3.Vertices[1].Normal = object.Make(0, 0, 1)
	half3.Vertices[2].Normal = object.Make(0, 0, 1)
	half3.Vertices[3].Normal = object.Make(0, 0, 1)

	half3.Polygons = append(half3.Polygons, object.Polygon{0, 1, 3, clr})
	half3.Polygons = append(half3.Polygons, object.Polygon{1, 2, 3, clr})

	half4.AddPoint(object.Point{0, 0, 0})
	half4.AddPoint(object.Point{50, 100, 0})
	half4.AddPoint(object.Point{0, 300, 0})
	half4.AddPoint(object.Point{0, 100, -20})

	half4.Vertices[0].Normal = object.Make(0, 0, -1)
	half4.Vertices[1].Normal = object.Make(0, 0, -1)
	half4.Vertices[2].Normal = object.Make(0, 0, -1)
	half4.Vertices[3].Normal = object.Make(0, 0, -1)

	half4.Polygons = append(half4.Polygons, object.Polygon{0, 1, 2, clr})
	half4.Polygons = append(half4.Polygons, object.Polygon{1, 2, 3, clr})

	half4.Rotate(object.Point{0, 0, 0}, object.Point{0, 120, 0})

	half3.Vertices[0].Normal.Add(half4.Vertices[0].Normal)
	half3.Vertices[2].Normal.Add(half4.Vertices[2].Normal)
	half3.Vertices[3].Normal.Add(half4.Vertices[3].Normal)

	half3.Vertices[0].Normal.Mul(0.5)
	half3.Vertices[2].Normal.Mul(0.5)
	half3.Vertices[3].Normal.Mul(0.5)

	leaf.Vertices = append(leaf.Vertices, half3.Vertices...)
	leaf.Vertices = append(leaf.Vertices, half4.Vertices[1])

	leaf.Polygons = append(leaf.Polygons, object.Polygon{4, 5, 7, clr}, object.Polygon{5, 6, 7, clr}, object.Polygon{4, 7, 8, clr}, object.Polygon{7, 2, 8, clr})

	leaf.Rotate(object.Point{0, 0, 0}, object.Point{0, -60, 0})

	leaf.Rotate(object.Point{0, 0, 0}, object.Point{0, 0, -30})

	return leaf
}

// func (c *Tulip) Rotate(center, angles object.Point) {
// 	fmt.Println("TULIPPPPPPPPPP")
// 	for i := range c.Components {
// 		c.Components[i].Rotate(center, angles)
// 	}
// }

func (t *Tulip) Rotate(center, angles object.Point) {
	t.CompositeModel.Rotate(center, angles)

	for i := 0; i < 4; i++ {
		t.stage1[i].Rotate(object.MakePoint(0, 0, 0), object.MakePoint(0, 0, angles.Z))
		t.stage2[i].Rotate(object.MakePoint(0, 0, 0), object.MakePoint(0, 0, angles.Z))
	}
}

func (t *Tulip) MakePetals(stage bezier.BezierCurve) {
	for i := range t.Petals {

		petal := MakePetal(stage, 20, 20, t.clr)

		//petal.Scale(object.Point{0, 0, 0}, 10)

		t.Add(&petal)
		t.Petals[i] = t.Size() - 1
		t.Components[t.Petals[i]].Rotate(object.Point{0, 0, 0}, object.Point{0, 120 * float64(i), 0})
		//t.Petals[i].Rotate(object.Point{0, 0, 0}, object.Point{0, 120 * float64(i), 0})

		if i > 2 {
			t.Components[t.Petals[i]].Rotate(object.Point{0, 0, 0}, object.Point{0, 60, 0})
			//t.Petals[i].Rotate(object.Point{0, 0, 0}, object.Point{0, 60, 0})
		}

		t.Components[t.Petals[i]].Move(object.Point{t.pos.X, t.pos.Y + t.stemLen, t.pos.Z})
	}

}

func (t *Tulip) ChangePetals(stage bezier.BezierCurve) {
	for i := range t.Petals {

		petal := MakePetal(stage, 20, 20, t.clr)

		//petal.Scale(object.Point{0, 0, 0}, 10)

		t.Components[t.Petals[i]] = &petal

		t.Components[t.Petals[i]].Rotate(object.Point{0, 0, 0}, object.Point{0, 120 * float64(i), 0})
		//t.Petals[i].Rotate(object.Point{0, 0, 0}, object.Point{0, 120 * float64(i), 0})

		if i > 2 {
			t.Components[t.Petals[i]].Rotate(object.Point{0, 0, 0}, object.Point{0, 60, 0})
			//t.Petals[i].Rotate(object.Point{0, 0, 0}, object.Point{0, 60, 0})
		}

		t.Components[t.Petals[i]].Move(object.Point{t.pos.X, t.pos.Y + t.stemLen, t.pos.Z})
	}

}

func (t Tulip) interpolateStagePoint(i int, k float64) object.Point {

	var p object.Point

	p.X = t.stage1[i].X + (k * (t.stage2[i].X - t.stage1[i].X))
	p.Y = t.stage1[i].Y + (k * (t.stage2[i].Y - t.stage1[i].Y))
	p.Z = t.stage1[i].Z + (k * (t.stage2[i].Z - t.stage1[i].Z))

	return p
}

func (t *Tulip) Animate(k float64) {
	t.OpenPetals(k)
	fmt.Println(t.Petals[0])
}

func (t *Tulip) OpenPetals(k float64) {
	var stage bezier.BezierCurve

	stage[0] = t.interpolateStagePoint(0, k)
	stage[1] = t.interpolateStagePoint(1, k)
	stage[2] = t.interpolateStagePoint(2, k)
	stage[3] = t.interpolateStagePoint(3, k)

	fmt.Println(stage)

	t.ChangePetals(stage)
}

func NewTulip(pos object.Point, stage int) *Tulip {
	t := new(Tulip)

	t.pos = pos

	t.clr = color.NRGBA{222, 0, 33, 255}

	t.stage1[0] = object.Point{0, 0, 0}
	t.stage1[1] = object.Point{6, 0, 0}
	t.stage1[2] = object.Point{4, 10, 0}
	t.stage1[3] = object.Point{0, 10, 0}

	t.stage2[0] = object.Point{0, 0, 0}
	t.stage2[1] = object.Point{6, 0, 0}
	t.stage2[2] = object.Point{6, 8, 0}
	t.stage2[3] = object.Point{5, 10.907, 0}

	for i := range t.stage1 {
		t.stage1[i].Scale(object.Point{0, 0, 0}, 10)
		t.stage2[i].Scale(object.Point{0, 0, 0}, 10)
	}

	t.stemLen = 300

	if stage == 1 {
		t.MakePetals(t.stage1)
	} else {
		t.MakePetals(t.stage2)
	}

	// Stem
	stem := MakeStem(t.stemLen, 5, 40, 20, color.NRGBA{0, 200, 25, 255})

	t.Add(&stem)
	t.Stem = t.Size() - 1

	t.Components[t.Stem].Move(object.Point{pos.X, pos.Y, pos.Z})

	leaf1 := MakeLeaf(color.NRGBA{0, 200, 25, 255})

	t.Add(&leaf1)
	t.Leaves[0] = t.Size() - 1

	t.Components[t.Leaves[0]].Move(object.Point{pos.X, pos.Y, pos.Z})

	leaf2 := MakeLeaf(color.NRGBA{0, 200, 25, 255})

	leaf2.Flip(false, true, false)

	t.Add(&leaf2)
	t.Leaves[1] = t.Size() - 1

	t.Components[t.Leaves[1]].Move(object.Point{pos.X, pos.Y, pos.Z})

	return t
}

func TestLeaf() object.Model {

	leaf := MakeLeaf(color.NRGBA{0, 255, 0, 255})

	leaf.Move(object.Point{100, 100, 100})

	return leaf

}
