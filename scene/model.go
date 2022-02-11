package scene

import (
	"image/color"
	"tulip/mymath"
)

type Model struct {
	Vertices []Vertex
	Polygons []Polygon
}

type PolygonialFunc func(v1, v2, v3 Vertex, clr color.NRGBA)

type PolygonialModel interface {
	IterateOverPolygons(f PolygonialFunc)

	Scale(center mymath.Vec3, k float64)
	Move(delta mymath.Vec3)
	Rotate(center, angles mymath.Vec3)

	Animate(k float64)
}

func (m *Model) Animate(k float64) {

}

func (m Model) IterateOverPolygons(f PolygonialFunc) {

	for i := 0; i < len(m.Polygons); i++ {
		v1, v2, v3 := m.Polygons[i].V1, m.Polygons[i].V2, m.Polygons[i].V3
		f(m.Vertices[v1], m.Vertices[v2], m.Vertices[v3], m.Polygons[i].Clr)
	}
}

func (m *Model) AddPoint(p mymath.Vec4) {
	var v Vertex
	v.Point = p
	m.Vertices = append(m.Vertices, v)
}

func (m *Model) AddPolygon(v1, v2, v3 int, clr color.NRGBA) {
	var p Polygon
	p.V1, p.V2, p.V3 = v1, v2, v3
	p.Clr = clr
	m.Polygons = append(m.Polygons, p)
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
