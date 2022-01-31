package primitives

import (
	"image/color"
	"tulip/mymath"
	"tulip/object"
)

type Cube struct {
	object.Model
}

func NewCube(side float64, center mymath.Vector3d, clr color.NRGBA) *object.Model {
	cube := new(object.Model)

	// bottom
	cube.AddPoint(mymath.MakeVector3d(0, 0, side))    // 0
	cube.AddPoint(mymath.MakeVector3d(0, 0, 0))       // 1
	cube.AddPoint(mymath.MakeVector3d(side, 0, 0))    // 2
	cube.AddPoint(mymath.MakeVector3d(side, 0, side)) // 3

	cube.Vertices[0].Normal = mymath.MakeVector3d(0.0, -1.0, 0.0)
	cube.Vertices[1].Normal = mymath.MakeVector3d(0.0, -1.0, 0.0)
	cube.Vertices[2].Normal = mymath.MakeVector3d(0.0, -1.0, 0.0)
	cube.Vertices[3].Normal = mymath.MakeVector3d(0.0, -1.0, 0.0)

	cube.AddPolygon(0, 1, 2, clr)
	cube.AddPolygon(2, 3, 0, clr)

	// top
	cube.AddPoint(mymath.MakeVector3d(0, side, 0))       // 4
	cube.AddPoint(mymath.MakeVector3d(0, side, side))    // 5
	cube.AddPoint(mymath.MakeVector3d(side, side, side)) // 6
	cube.AddPoint(mymath.MakeVector3d(side, side, 0))    // 7

	cube.Vertices[4].Normal = mymath.MakeVector3d(0.0, 1.0, 0.0)
	cube.Vertices[5].Normal = mymath.MakeVector3d(0.0, 1.0, 0.0)
	cube.Vertices[6].Normal = mymath.MakeVector3d(0.0, 1.0, 0.0)
	cube.Vertices[7].Normal = mymath.MakeVector3d(0.0, 1.0, 0.0)

	cube.AddPolygon(4, 5, 6, clr)
	cube.AddPolygon(6, 7, 4, clr)

	// front
	cube.AddPoint(mymath.MakeVector3d(0, side, side))    // 8
	cube.AddPoint(mymath.MakeVector3d(0, 0, side))       // 9
	cube.AddPoint(mymath.MakeVector3d(side, 0, side))    // 10
	cube.AddPoint(mymath.MakeVector3d(side, side, side)) // 11

	cube.Vertices[8].Normal = mymath.MakeVector3d(0.0, 0.0, 1.0)
	cube.Vertices[9].Normal = mymath.MakeVector3d(0.0, 0.0, 1.0)
	cube.Vertices[10].Normal = mymath.MakeVector3d(0.0, 0.0, 1.0)
	cube.Vertices[11].Normal = mymath.MakeVector3d(0.0, 0.0, 1.0)

	cube.AddPolygon(8, 9, 10, clr)
	cube.AddPolygon(10, 11, 8, clr)

	// back
	cube.AddPoint(mymath.MakeVector3d(0, 0, 0))       // 12
	cube.AddPoint(mymath.MakeVector3d(0, side, 0))    // 13
	cube.AddPoint(mymath.MakeVector3d(side, side, 0)) // 14
	cube.AddPoint(mymath.MakeVector3d(side, 0, 0))    // 15

	cube.Vertices[12].Normal = mymath.MakeVector3d(0.0, 0.0, -1.0)
	cube.Vertices[13].Normal = mymath.MakeVector3d(0.0, 0.0, -1.0)
	cube.Vertices[14].Normal = mymath.MakeVector3d(0.0, 0.0, -1.0)
	cube.Vertices[15].Normal = mymath.MakeVector3d(0.0, 0.0, -1.0)

	cube.AddPolygon(12, 13, 14, clr)
	cube.AddPolygon(14, 15, 12, clr)

	// left
	cube.AddPoint(mymath.MakeVector3d(0, side, 0))    // 16
	cube.AddPoint(mymath.MakeVector3d(0, 0, 0))       // 17
	cube.AddPoint(mymath.MakeVector3d(0, 0, side))    // 18
	cube.AddPoint(mymath.MakeVector3d(0, side, side)) // 19

	cube.Vertices[16].Normal = mymath.MakeVector3d(-1.0, 0.0, 0.0)
	cube.Vertices[17].Normal = mymath.MakeVector3d(-1.0, 0.0, 0.0)
	cube.Vertices[18].Normal = mymath.MakeVector3d(-1.0, 0.0, 0.0)
	cube.Vertices[19].Normal = mymath.MakeVector3d(-1.0, 0.0, 0.0)

	cube.AddPolygon(16, 17, 18, clr)
	cube.AddPolygon(18, 19, 16, clr)

	// right
	cube.AddPoint(mymath.MakeVector3d(side, side, side)) // 20
	cube.AddPoint(mymath.MakeVector3d(side, 0, side))    // 21
	cube.AddPoint(mymath.MakeVector3d(side, 0, 0))       // 22
	cube.AddPoint(mymath.MakeVector3d(side, side, 0))    // 23

	cube.Vertices[20].Normal = mymath.MakeVector3d(1.0, 0.0, 0.0)
	cube.Vertices[21].Normal = mymath.MakeVector3d(1.0, 0.0, 0.0)
	cube.Vertices[22].Normal = mymath.MakeVector3d(1.0, 0.0, 0.0)
	cube.Vertices[23].Normal = mymath.MakeVector3d(1.0, 0.0, 0.0)

	cube.AddPolygon(20, 21, 22, clr)
	cube.AddPolygon(22, 23, 20, clr)

	cube.Move(mymath.Vector3dDiff(center, mymath.MakeVector3d(side/2, side/2, side/2)))
	return cube
}
