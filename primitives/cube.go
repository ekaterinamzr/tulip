package primitives

import (
	"image/color"
	"tulip/object"
)

type Cube struct {
	object.Model
}

func NewCube(side float64, center object.Point, clr color.NRGBA) *object.Model {
	cube := new(object.Model)

	//p := []object.Polygon{{2, 6, 5, clr}, {6, 7, 5, clr}, {5, 7, 3, clr}, {7, 4, 3, clr}, {4, 0, 3, clr}, {4, 1, 0, clr}, {1, 2, 0, clr}, {1, 6, 2, clr}, {1, 7, 6, clr}, {1, 4, 7, clr}, {0, 2, 5, clr}, {2, 5, 3, clr}}

	// bottom
	cube.AddPoint(object.MakePoint(0, 0, side))    // 0
	cube.AddPoint(object.MakePoint(0, 0, 0))       // 1
	cube.AddPoint(object.MakePoint(side, 0, 0))    // 2
	cube.AddPoint(object.MakePoint(side, 0, side)) // 3

	cube.Vertices[0].Normal = object.Make(0.0, -1.0, 0.0)
	cube.Vertices[1].Normal = object.Make(0.0, -1.0, 0.0)
	cube.Vertices[2].Normal = object.Make(0.0, -1.0, 0.0)
	cube.Vertices[3].Normal = object.Make(0.0, -1.0, 0.0)

	cube.AddPolygon(0, 1, 2, clr)
	cube.AddPolygon(2, 3, 0, clr)

	// top
	cube.AddPoint(object.MakePoint(0, side, 0))       // 4
	cube.AddPoint(object.MakePoint(0, side, side))    // 5
	cube.AddPoint(object.MakePoint(side, side, side)) // 6
	cube.AddPoint(object.MakePoint(side, side, 0))    // 7

	cube.Vertices[4].Normal = object.Make(0.0, 1.0, 0.0)
	cube.Vertices[5].Normal = object.Make(0.0, 1.0, 0.0)
	cube.Vertices[6].Normal = object.Make(0.0, 1.0, 0.0)
	cube.Vertices[7].Normal = object.Make(0.0, 1.0, 0.0)

	cube.AddPolygon(4, 5, 6, clr)
	cube.AddPolygon(6, 7, 4, clr)

	// front
	cube.AddPoint(object.MakePoint(0, side, side))    // 8
	cube.AddPoint(object.MakePoint(0, 0, side))       // 9
	cube.AddPoint(object.MakePoint(side, 0, side))    // 10
	cube.AddPoint(object.MakePoint(side, side, side)) // 11

	cube.Vertices[8].Normal = object.Make(0.0, 0.0, 1.0)
	cube.Vertices[9].Normal = object.Make(0.0, 0.0, 1.0)
	cube.Vertices[10].Normal = object.Make(0.0, 0.0, 1.0)
	cube.Vertices[11].Normal = object.Make(0.0, 0.0, 1.0)

	cube.AddPolygon(8, 9, 10, clr)
	cube.AddPolygon(10, 11, 8, clr)

	// back
	cube.AddPoint(object.MakePoint(0, 0, 0))       // 12
	cube.AddPoint(object.MakePoint(0, side, 0))    // 13
	cube.AddPoint(object.MakePoint(side, side, 0)) // 14
	cube.AddPoint(object.MakePoint(side, 0, 0))    // 15

	cube.Vertices[12].Normal = object.Make(0.0, 0.0, -1.0)
	cube.Vertices[13].Normal = object.Make(0.0, 0.0, -1.0)
	cube.Vertices[14].Normal = object.Make(0.0, 0.0, -1.0)
	cube.Vertices[15].Normal = object.Make(0.0, 0.0, -1.0)

	cube.AddPolygon(12, 13, 14, clr)
	cube.AddPolygon(14, 15, 12, clr)

	// left
	cube.AddPoint(object.MakePoint(0, side, 0))    // 16
	cube.AddPoint(object.MakePoint(0, 0, 0))       // 17
	cube.AddPoint(object.MakePoint(0, 0, side))    // 18
	cube.AddPoint(object.MakePoint(0, side, side)) // 19

	cube.Vertices[16].Normal = object.Make(-1.0, 0.0, 0.0)
	cube.Vertices[17].Normal = object.Make(-1.0, 0.0, 0.0)
	cube.Vertices[18].Normal = object.Make(-1.0, 0.0, 0.0)
	cube.Vertices[19].Normal = object.Make(-1.0, 0.0, 0.0)

	cube.AddPolygon(16, 17, 18, clr)
	cube.AddPolygon(18, 19, 16, clr)

	// right
	cube.AddPoint(object.MakePoint(side, side, side)) // 20
	cube.AddPoint(object.MakePoint(side, 0, side))    // 21
	cube.AddPoint(object.MakePoint(side, 0, 0))       // 22
	cube.AddPoint(object.MakePoint(side, side, 0))    // 23

	cube.Vertices[20].Normal = object.Make(1.0, 0.0, 0.0)
	cube.Vertices[21].Normal = object.Make(1.0, 0.0, 0.0)
	cube.Vertices[22].Normal = object.Make(1.0, 0.0, 0.0)
	cube.Vertices[23].Normal = object.Make(1.0, 0.0, 0.0)

	cube.AddPolygon(20, 21, 22, clr)
	cube.AddPolygon(22, 23, 20, clr)

	//cube.CalculateNormals()

	//fmt.Println(len(cube.Vertices))

	//cube.Move(object.MakePoint(center.X/2, center.Y/2, center.Z/2))
	cube.Move(center.Sub(object.MakePoint(side/2, side/2, side/2)))
	return cube
}
