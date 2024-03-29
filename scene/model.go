package scene

import (
	"image/color"
	"tulip/mymath"
)

type Model struct {
	Vertices []Vertex
	// Polygons []Polygon
	Indices []int
}

type PolygonialFunc func(v1, v2, v3 Vertex, clr color.NRGBA)

type PolygonialModel interface {
	GetVertices() ([]Vertex, []int)
	IterateOverPolygons(f PolygonialFunc)

	Scale(center mymath.Vec3, k float64)
	Move(delta mymath.Vec3)
	Rotate(center, angles mymath.Vec3)

	Animate(k float64)
}

func (m *Model) Animate(k float64) {

}

func (m Model) GetVertices() ([]Vertex, []int){ 
	return m.Vertices, m.Indices
}

// func (m Model) GetVerticesOld() ([]Vertex, []int){
// 	for i := range(m.Polygons) {
// 		vIdx0, vIdx1, vIdx2 := m.Polygons[i].V1, m.Polygons[i].V2, m.Polygons[i].V3
// 		clr := m.Polygons[i].Clr
// 		// v0, v1, v2 := m.Vertices[vIdx0], m.Vertices[vIdx1], m.Vertices[vIdx2]
// 		m.Vertices[vIdx0].Clr = clr
// 		m.Vertices[vIdx1].Clr = clr
// 		m.Vertices[vIdx2].Clr = clr
// 	}

// 	vertices := make([]Vertex, len(m.Vertices))
// 	indices := make([]int, len(m.Polygons) * 3)

// 	copy(vertices, m.Vertices)

// 	for _, p := range(m.Polygons) {
// 		indices = append(indices, p.V1)
// 		indices = append(indices, p.V2)
// 		indices = append(indices, p.V3)
// 	}

// 	return vertices, indices
// }

// func (m Model) IterateOverPolygons(f PolygonialFunc) {

// 	for i := 0; i < len(m.Polygons); i++ {
// 		v1, v2, v3 := m.Polygons[i].V1, m.Polygons[i].V2, m.Polygons[i].V3
// 		f(m.Vertices[v1], m.Vertices[v2], m.Vertices[v3], m.Polygons[i].Clr)
// 	}
// }

func (m *Model) AddPoint(p mymath.Vec4) {
	var v Vertex
	v.Point = p
	m.Vertices = append(m.Vertices, v)
}

func (m *Model) AddPolygon(v0, v1, v2 int, clr color.NRGBA) {
	// var p Polygon
	// p.V1, p.V2, p.V3 = v1, v2, v3
	// p.Clr = clr
	// m.Polygons = append(m.Polygons, p)

	m.Vertices[v0].Clr = clr
	m.Vertices[v1].Clr = clr
	m.Vertices[v2].Clr = clr
	m.Indices = append(m.Indices, v0, v1, v2)
}

func (m *Model) Scale(center mymath.Vec3, k float64) {
	for i := range m.Vertices {
		m.Vertices[i].Scale(center, k)
	}
}

func (m *Model) Move(delta mymath.Vec3) {
	for i := range m.Vertices {
		m.Vertices[i].Move(delta)
	}
}

func (m *Model) Rotate(center, angles mymath.Vec3) {
	for i := range m.Vertices {
		m.Vertices[i].Rotate(center, angles)
	}
}

func (m *Model) Reflect(x, y, z bool) {
	for i := range m.Vertices {
		m.Vertices[i].Reflect(x, y, z)
	}
}
