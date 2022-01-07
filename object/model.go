package object

import (
	"image/color"
	"math"
)

type Vertex struct {
	X float64
	Y float64
	Z float64
}

type Polygon struct {
	V1 int
	V2 int
	V3 int

	Clr color.Color
}

type Model struct {
	Vertices []Vertex
	Polygons []Polygon
}

func (m Model) IterateOverPolygons(f PolygonialFunc) {

	for i := 0; i < len(m.Polygons); i++ {
		v1, v2, v3 := m.Polygons[i].V1, m.Polygons[i].V2, m.Polygons[i].V3
		//fmt.Println(m.Vertices[v1], m.Vertices[v2], m.Vertices[v3])
		f(m.Vertices[v1], m.Vertices[v2], m.Vertices[v3], m.Polygons[i].Clr)
	}
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

func (a Vertex) Add(b Vertex) Vertex {
	return Vertex{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vertex) Sub(b Vertex) Vertex {
	return Vertex{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vertex) Mul(k float64) Vertex {
	return Vertex{a.X * k, a.Y * k, a.Z * k}
}

func (v *Vertex) ScaleVertex(center Vertex, k float64) {
	v.X = center.X + k*(v.X-center.X)
	v.Y = center.Y + k*(v.Y-center.Y)
	v.Z = center.Z + k*(v.Z-center.Z)
}

func (o *Model) Scale(center Vertex, k float64) {
	for i := 0; i < len(o.Vertices); i++ {
		o.Vertices[i].ScaleVertex(center, k)
	}
}

func (v *Vertex) MoveVertex(delta Vertex) {
	v.X += delta.X
	v.Y += delta.Y
	v.Z += delta.Z
}

func (o *Model) Move(delta Vertex) {
	for i := 0; i < len(o.Vertices); i++ {
		o.Vertices[i].MoveVertex(delta)
	}
}

func toRadians(angle float64) float64 {
	return angle * 3.14 / 180
}

func (v *Vertex) rotateZ(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*v.X + sin*v.Y
	newY := cos*v.Y - sin*v.X

	v.X = newX
	v.Y = newY
}

func (v *Vertex) rotateX(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newZ := cos*v.Z - sin*v.Y
	newY := sin*v.Z + cos*v.Y

	v.Z = newZ
	v.Y = newY
}

func (v *Vertex) rotateY(angle float64) {
	cos := math.Cos(float64(angle))
	sin := math.Sin(float64(angle))

	newX := cos*v.X + sin*v.Z
	newZ := cos*v.Z - sin*v.X

	v.X = newX
	v.Z = newZ
}

func (v *Vertex) RotateVertex(center, angles Vertex) {
	angleX := toRadians(angles.X)
	angleY := toRadians(angles.Y)
	angleZ := toRadians(angles.Z)

	v.MoveVertex(Vertex{-center.X, -center.Y, -center.Z})

	if angleX != 0 {
		v.rotateX(angleX)
	}

	if angleY != 0 {
		v.rotateY(angleY)
	}

	if angleZ != 0 {
		v.rotateZ(angleZ)
	}

	v.MoveVertex(center)
}

func (o *Model) Rotate(center, angles Vertex) {
	for i := 0; i < len(o.Vertices); i++ {
		o.Vertices[i].RotateVertex(center, angles)
	}
}

func (v *Vertex) Flip(x, y, z bool) {
	if x {
		v.X = -v.X
	}
	if y {
		v.X = -v.X
	}
	if z {
		v.Z = -v.Z
	}
}

func (m *Model) Flip(x, y, z bool) {
	for i := 0; i < len(m.Vertices); i++ {
		m.Vertices[i].Flip(x, y, z)
	}
}
