package object

import (
	"image/color"
)

type PolygonialFunc func(v1, v2, v3 Vertex, clr color.Color)

type PolygonialModel interface {
	IterateOverPolygons(f PolygonialFunc)

	Scale(center Vertex, k float64)
	Move(delta Vertex)
	Rotate(center, angles Vertex)
}

type CompositeModel struct {
	Components []PolygonialModel
}

func (c *CompositeModel) Add(m PolygonialModel) {
	c.Components = append(c.Components, m)
}

func (c CompositeModel) Size() int {
	return len(c.Components)
}

func (c CompositeModel) IterateOverPolygons(f PolygonialFunc) {
	for i := range c.Components {
		c.Components[i].IterateOverPolygons(f)
	}
}

func (c *CompositeModel) Scale(center Vertex, k float64) {
	for i := range c.Components {
		c.Components[i].Scale(center, k)
	}
}

func (c *CompositeModel) Move(delta Vertex) {
	for i := range c.Components {
		c.Components[i].Move(delta)
	}
}

func (c *CompositeModel) Rotate(center, angles Vertex) {
	for i := range c.Components {
		c.Components[i].Rotate(center, angles)
	}
}

// func (c CompositeModel) SingleModel() Model {
// 	var combined Model

// 	for i := range c.Components {

// 	}
// }
