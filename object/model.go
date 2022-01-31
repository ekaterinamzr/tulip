package object

import (
	"image/color"
)

type Model struct {
	Vertices []Vertex
	Polygons []Polygon
	//Normals  []Vector3d
}

type PolygonialFunc func(v1, v2, v3 Vertex, clr color.NRGBA)

type PolygonialModel interface {
	IterateOverPolygons(f PolygonialFunc)

	Scale(center Point, k float64)
	Move(delta Point)
	Rotate(center, angles Point)

	Animate(k float64)
}

func (m *Model) Animate(k float64) {

}

func (m *Model) CalculateNormals() {
	for _, p := range m.Polygons {
		v1, v2, v3 := m.Vertices[p.V1], m.Vertices[p.V2], m.Vertices[p.V3]

		//fmt.Println(v1.Point, v3.Point, v3.Point)
		normal := PolygonNormal(v1, v2, v3)

		//fmt.Println(normal)

		m.Vertices[p.V1].Normal.Add(normal)
		m.Vertices[p.V2].Normal.Add(normal)
		m.Vertices[p.V3].Normal.Add(normal)

		m.Vertices[p.V1].cnt++
		m.Vertices[p.V2].cnt++
		m.Vertices[p.V3].cnt++
	}

	for i := range m.Vertices {
		m.Vertices[i].Normal.Mul(1.0 / float64(m.Vertices[i].cnt))
	}
}

func (m Model) IterateOverPolygons(f PolygonialFunc) {

	for i := 0; i < len(m.Polygons); i++ {
		v1, v2, v3 := m.Polygons[i].V1, m.Polygons[i].V2, m.Polygons[i].V3
		f(m.Vertices[v1], m.Vertices[v2], m.Vertices[v3], m.Polygons[i].Clr)
	}
}

func (m *Model) AddPoint(p Point) {
	var v Vertex
	v.X, v.Y, v.Z = p.X, p.Y, p.Z
	m.Vertices = append(m.Vertices, v)
}

func (m *Model) AddPolygon(v1, v2, v3 int, clr color.NRGBA) {
	var p Polygon
	p.V1, p.V2, p.V3 = v1, v2, v3
	p.Clr = clr
	m.Polygons = append(m.Polygons, p)
}

func (o *Model) SortVertices() {
	for i, p := range o.Polygons {
		if o.Vertices[p.V1].Y > o.Vertices[p.V2].Y {
			o.Polygons[i].V1, o.Polygons[i].V2 = o.Polygons[i].V2, o.Polygons[i].V1
			//o.Vertices[p.V1], o.Vertices[p.V2] = o.Vertices[p.V2], o.Vertices[p.V1]
		}
		if o.Vertices[p.V1].Y > o.Vertices[p.V3].Y {
			o.Polygons[i].V1, o.Polygons[i].V3 = o.Polygons[i].V3, o.Polygons[i].V1
			//o.Vertices[p.V1], o.Vertices[p.V3] = o.Vertices[p.V3], o.Vertices[p.V1]
		}
		if o.Vertices[p.V2].Y > o.Vertices[p.V3].Y {
			o.Polygons[i].V2, o.Polygons[i].V3 = o.Polygons[i].V3, o.Polygons[i].V2
			//o.Vertices[p.V2], o.Vertices[p.V3] = o.Vertices[p.V3], o.Vertices[p.V2]
		}
	}
}

func (m *Model) Scale(center Point, k float64) {
	for i := range m.Vertices {
		m.Vertices[i].Scale(center, k)
	}
}

func (m *Model) Move(delta Point) {
	for i := range m.Vertices {
		m.Vertices[i].Move(delta)
	}
}

func (m *Model) Rotate(center, angles Point) {
	for i := range m.Vertices {
		m.Vertices[i].Rotate(center, angles)
	}
}

func (m *Model) Flip(x, y, z bool) {
	for i := range m.Vertices {
		m.Vertices[i].Flip(x, y, z)
	}
}
