package object

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

func (c *CompositeModel) Scale(center Point, k float64) {
	for i := range c.Components {
		c.Components[i].Scale(center, k)
	}
}

func (c *CompositeModel) Move(delta Point) {
	for i := range c.Components {
		c.Components[i].Move(delta)
	}
}

func (c *CompositeModel) Rotate(center, angles Point) {
	for i := range c.Components {
		c.Components[i].Rotate(center, angles)
	}
}

// func (c CompositeModel) SingleModel() Model {
// 	var combined Model

// 	for i := range c.Components {

// 	}
// }
