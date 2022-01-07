package flower

import (
	//"fmt"
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
	curve[0] = object.Vertex{0, 0, 0}
	curve[1] = object.Vertex{6, 0, 0}
	curve[2] = object.Vertex{4, 10, 0}
	curve[3] = object.Vertex{0, 10, 0}

	for i := range curve {
		curve[i].ScaleVertex(object.Vertex{0, 0, 0}, 10)
	}

	petal := MakePetal(curve, 5, 5, color.RGBA{255, 0, 0, 255})

	//petal.ScaleModel(object.Vertex{0, 0, 0}, 10)
	petal.Move(object.Vertex{100, 100, 0})

	return petal
}

func MakePetal(curve bezier.BezierCurve, m, n int, clr color.Color) object.Model {
	var (
		petal   object.Model
		surface bezier.BicubicBezierSurface
	)

	surface[0] = curve
	surface[1] = curve
	surface[2] = curve
	surface[3] = curve

	surface[1][1].RotateVertex(object.Vertex{0, surface[1][1].Y, 0},
		object.Vertex{0, 20, 0})
	surface[1][2].RotateVertex(object.Vertex{0, surface[1][2].Y, 0},
		object.Vertex{0, 20, 0})

	surface[2][1].RotateVertex(object.Vertex{0, surface[2][1].Y, 0},
		object.Vertex{0, 40, 0})
	surface[2][2].RotateVertex(object.Vertex{0, surface[2][2].Y, 0},
		object.Vertex{0, 40, 0})

	surface[3][1].RotateVertex(object.Vertex{0, surface[3][1].Y, 0},
		object.Vertex{0, 60, 0})
	surface[3][2].RotateVertex(object.Vertex{0, surface[3][2].Y, 0},
		object.Vertex{0, 60, 0})

	u := 1 / float64(n)
	v := 1 / float64(m)

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			vertex := surface.GetPoint(v*float64(j), u*float64(i))
			petal.Vertices = append(petal.Vertices, vertex)
		}
	}

	for i := 0; i < (n+1)*m; i++ {
		//fmt.Println("jopa")
		if i != 0 && (i+1)%(n+1) == 0 {
			continue
		}
		petal.Polygons = append(petal.Polygons, object.Polygon{i, i + 1, i + n + 2, clr})
		petal.Polygons = append(petal.Polygons, object.Polygon{i, i + n + 1, i + n + 2, clr})
	}

	for i := range petal.Polygons {
		v1 := petal.Vertices[petal.Polygons[i].V1]
		v2 := petal.Vertices[petal.Polygons[i].V2]
		v3 := petal.Vertices[petal.Polygons[i].V3]

		v1.Z = -v1.Z
		v2.Z = -v2.Z
		v3.Z = -v3.Z

		n := len(petal.Vertices)

		petal.Vertices = append(petal.Vertices, v1, v2, v3)
		petal.Polygons = append(petal.Polygons, object.Polygon{n, n + 1, n + 2, clr})
	}

	//fmt.Println(petal)

	return petal
}

func MakeStem(h, r float64, n, k int, clr color.Color) object.Model {
	var stem object.Model

	dy := h / float64(n)
	df := math.Pi * 2 / float64(k)

	//dz := math.Cos(df)
	//dx := math.Sin(df)

	for i := 0; i < n+1; i++ {
		y := float64(i) * dy
		x := r
		z := 0.0

		//stem.Vertices = append(stem.Vertices, object.Vertex{x, y, z})

		for f := math.Pi / 2; f < 5*math.Pi/2; f += df {
			x = math.Sin(f) * r
			z = math.Cos(f) * r

			v := object.Vertex{x, y, z}
			//fmt.Println(v)

			stem.Vertices = append(stem.Vertices, v)
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

	return stem

}

func MakeLeaf(clr color.Color) object.Model {
	var (
		half1, half2, leaf object.Model
		//leaf1 object.CompositeModel
	)
	half1.Vertices = append(half1.Vertices, object.Vertex{0, 0, 0}, object.Vertex{50, 100, 0}, object.Vertex{0, 300, 0})
	half1.Polygons = append(half1.Polygons, object.Polygon{0, 1, 2, clr})

	half2.Vertices = append(half1.Vertices, object.Vertex{0, 0, 0}, object.Vertex{50, 100, 0}, object.Vertex{0, 300, 0})
	half2.Polygons = append(half1.Polygons, object.Polygon{0, 1, 2, clr})

	half2.Rotate(object.Vertex{0, 0, 0}, object.Vertex{0, 120, 0})

	leaf.Vertices = append(leaf.Vertices, half1.Vertices...)
	leaf.Vertices = append(leaf.Vertices, half2.Vertices...)

	leaf.Polygons = append(leaf.Polygons, object.Polygon{0, 1, 2, clr}, object.Polygon{3, 4, 5, clr})

	leaf.Rotate(object.Vertex{0, 0, 0}, object.Vertex{0, -60, 0})

	leaf.Rotate(object.Vertex{0, 0, 0}, object.Vertex{0, 0, -30})

	return leaf
}

// func (c *Tulip) Rotate(center, angles object.Vertex) {
// 	fmt.Println("TULIPPPPPPPPPP")
// 	for i := range c.Components {
// 		c.Components[i].Rotate(center, angles)
// 	}
// }

func NewTulip(pos object.Vertex) *Tulip {
	t := new(Tulip)

	t.clr = color.NRGBA{255, 0, 0, 255}

	t.stage1[0] = object.Vertex{0, 0, 0}
	t.stage1[1] = object.Vertex{6, 0, 0}
	t.stage1[2] = object.Vertex{4, 10, 0}
	t.stage1[3] = object.Vertex{0, 10, 0}

	t.stage2[0] = object.Vertex{0, 0, 0}
	t.stage2[1] = object.Vertex{6, 0, 0}
	t.stage2[2] = object.Vertex{6, 8, 0}
	t.stage2[3] = object.Vertex{5, 10.907, 0}

	for i := range t.stage1 {
		t.stage1[i].ScaleVertex(object.Vertex{0, 0, 0}, 10)
		t.stage2[i].ScaleVertex(object.Vertex{0, 0, 0}, 10)
	}

	var stem_len float64 = 300

	for i := range t.Petals {
		petal := MakePetal(t.stage2, 10, 10, t.clr)
		//petal.Scale(object.Vertex{0, 0, 0}, 10)

		t.Add(&petal)
		t.Petals[i] = t.Size() - 1
		t.Components[t.Petals[i]].Rotate(object.Vertex{0, 0, 0}, object.Vertex{0, 120 * float64(i), 0})
		//t.Petals[i].Rotate(object.Vertex{0, 0, 0}, object.Vertex{0, 120 * float64(i), 0})

		if i > 2 {
			t.Components[t.Petals[i]].Rotate(object.Vertex{0, 0, 0}, object.Vertex{0, 60, 0})
			//t.Petals[i].Rotate(object.Vertex{0, 0, 0}, object.Vertex{0, 60, 0})
		}

		t.Components[t.Petals[i]].Move(object.Vertex{pos.X, pos.Z + stem_len, pos.Y})
	}

	// Stem
	stem := MakeStem(stem_len, 5, 30, 10, color.RGBA{0, 255, 0, 255})

	t.Add(&stem)
	t.Stem = t.Size() - 1

	t.Components[t.Stem].Move(object.Vertex{pos.X, pos.Z, pos.Y})

	leaf1 := MakeLeaf(color.RGBA{0, 255, 0, 255})

	t.Add(&leaf1)
	t.Leaves[0] = t.Size() - 1

	t.Components[t.Leaves[0]].Move(object.Vertex{pos.X, pos.Z, pos.Y})

	leaf2 := MakeLeaf(color.RGBA{0, 255, 0, 255})

	leaf2.Flip(false, true, false)

	t.Add(&leaf2)
	t.Leaves[1] = t.Size() - 1

	t.Components[t.Leaves[1]].Move(object.Vertex{pos.X, pos.Z, pos.Y})

	return t
}

func TestLeaf() object.Model {

	leaf := MakeLeaf(color.RGBA{0, 255, 0, 255})

	leaf.Move(object.Vertex{100, 100, 100})

	return leaf

}
