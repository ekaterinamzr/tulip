package object

import (
	"tulip/mymath"
)

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

func (c *CompositeModel) Scale(center mymath.Vector3d, k float64) {
	for i := range c.Components {
		c.Components[i].Scale(center, k)
	}
}

func (c *CompositeModel) Move(delta mymath.Vector3d) {
	for i := range c.Components {
		c.Components[i].Move(delta)
	}
}

func (c *CompositeModel) Rotate(center, angles mymath.Vector3d) {
	for i := range c.Components {
		c.Components[i].Rotate(center, angles)
	}
}
