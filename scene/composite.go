package scene

import (
	"tulip/mymath"
)

type CompositeModel struct {
	Components []Model
}

func (c *CompositeModel) Add(m Model) {
	c.Components = append(c.Components, m)
}

func (c CompositeModel) Size() int {
	return len(c.Components)
}

func (c CompositeModel) IsComposit() bool{
	return true
}

func (c CompositeModel) GetVertices() ([]Vertex, []int){
	allVertices := make([]Vertex, 0, 1000)
	allIndices := make([]int, 0, 1000 * 3)
	verticesLen := 0
	
	for i := range(c.Components) {
		vertices, indices := c.Components[i].GetVertices()

		allVertices = append(allVertices, vertices...)

		for j := range(indices) {
			allIndices = append(allIndices, indices[j] + verticesLen)
		}
		verticesLen = len(vertices)
	}

	return allVertices, allIndices
}

// func (c CompositeModel) IterateOverPolygons(f PolygonialFunc) {
// 	for i := range c.Components {
// 		c.Components[i].IterateOverPolygons(f)
// 	}
// }

func (c *CompositeModel) Scale(center mymath.Vec3, k float64) {
	for i := range c.Components {
		c.Components[i].Scale(center, k)
	}
}

func (c *CompositeModel) Move(delta mymath.Vec3) {
	for i := range c.Components {
		c.Components[i].Move(delta)
	}
}

func (c *CompositeModel) Rotate(center, angles mymath.Vec3) {
	for i := range c.Components {
		c.Components[i].Rotate(center, angles)
	}
}
