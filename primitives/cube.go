package primitives

import (
	"image/color"
	"tulip/mymath"
	"tulip/scene"
)

func NewCube(side float64, center mymath.Vec3d, clr color.NRGBA) *scene.Model {
	cube := new(scene.Model)

	// bottom
	cube.AddPoint(mymath.MakeVec3d(0, 0, side))    // 0
	cube.AddPoint(mymath.MakeVec3d(0, 0, 0))       // 1
	cube.AddPoint(mymath.MakeVec3d(side, 0, 0))    // 2
	cube.AddPoint(mymath.MakeVec3d(side, 0, side)) // 3

	cube.Vertices[0].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)
	cube.Vertices[1].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)
	cube.Vertices[2].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)
	cube.Vertices[3].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)

	cube.AddPolygon(0, 1, 2, clr)
	cube.AddPolygon(2, 3, 0, clr)

	// top
	cube.AddPoint(mymath.MakeVec3d(0, side, 0))       // 4
	cube.AddPoint(mymath.MakeVec3d(0, side, side))    // 5
	cube.AddPoint(mymath.MakeVec3d(side, side, side)) // 6
	cube.AddPoint(mymath.MakeVec3d(side, side, 0))    // 7

	cube.Vertices[4].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)
	cube.Vertices[5].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)
	cube.Vertices[6].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)
	cube.Vertices[7].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)

	cube.AddPolygon(4, 5, 6, clr)
	cube.AddPolygon(6, 7, 4, clr)

	// front
	cube.AddPoint(mymath.MakeVec3d(0, side, side))    // 8
	cube.AddPoint(mymath.MakeVec3d(0, 0, side))       // 9
	cube.AddPoint(mymath.MakeVec3d(side, 0, side))    // 10
	cube.AddPoint(mymath.MakeVec3d(side, side, side)) // 11

	cube.Vertices[8].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)
	cube.Vertices[9].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)
	cube.Vertices[10].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)
	cube.Vertices[11].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)

	cube.AddPolygon(8, 9, 10, clr)
	cube.AddPolygon(10, 11, 8, clr)

	// back
	cube.AddPoint(mymath.MakeVec3d(0, 0, 0))       // 12
	cube.AddPoint(mymath.MakeVec3d(0, side, 0))    // 13
	cube.AddPoint(mymath.MakeVec3d(side, side, 0)) // 14
	cube.AddPoint(mymath.MakeVec3d(side, 0, 0))    // 15

	cube.Vertices[12].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)
	cube.Vertices[13].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)
	cube.Vertices[14].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)
	cube.Vertices[15].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)

	cube.AddPolygon(12, 13, 14, clr)
	cube.AddPolygon(14, 15, 12, clr)

	// left
	cube.AddPoint(mymath.MakeVec3d(0, side, 0))    // 16
	cube.AddPoint(mymath.MakeVec3d(0, 0, 0))       // 17
	cube.AddPoint(mymath.MakeVec3d(0, 0, side))    // 18
	cube.AddPoint(mymath.MakeVec3d(0, side, side)) // 19

	cube.Vertices[16].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)
	cube.Vertices[17].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)
	cube.Vertices[18].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)
	cube.Vertices[19].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)

	cube.AddPolygon(16, 17, 18, clr)
	cube.AddPolygon(18, 19, 16, clr)

	// right
	cube.AddPoint(mymath.MakeVec3d(side, side, side)) // 20
	cube.AddPoint(mymath.MakeVec3d(side, 0, side))    // 21
	cube.AddPoint(mymath.MakeVec3d(side, 0, 0))       // 22
	cube.AddPoint(mymath.MakeVec3d(side, side, 0))    // 23

	cube.Vertices[20].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)
	cube.Vertices[21].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)
	cube.Vertices[22].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)
	cube.Vertices[23].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)

	cube.AddPolygon(20, 21, 22, clr)
	cube.AddPolygon(22, 23, 20, clr)

	cube.Move(mymath.Vec3dDiff(center, mymath.MakeVec3d(side/2, side/2, side/2)))
	return cube
}

func NewBlock(a, b, c float64, center mymath.Vec3d, clr color.NRGBA) *scene.Model {
	cube := new(scene.Model)

	// bottom
	cube.AddPoint(mymath.MakeVec3d(0, 0, c)) // 0
	cube.AddPoint(mymath.MakeVec3d(0, 0, 0)) // 1
	cube.AddPoint(mymath.MakeVec3d(a, 0, 0)) // 2
	cube.AddPoint(mymath.MakeVec3d(a, 0, c)) // 3

	cube.Vertices[0].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)
	cube.Vertices[1].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)
	cube.Vertices[2].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)
	cube.Vertices[3].Normal = mymath.MakeVec3d(0.0, 1.0, 0.0)

	cube.AddPolygon(0, 1, 2, clr)
	cube.AddPolygon(2, 3, 0, clr)

	// top
	cube.AddPoint(mymath.MakeVec3d(0, b, 0)) // 4
	cube.AddPoint(mymath.MakeVec3d(0, b, c)) // 5
	cube.AddPoint(mymath.MakeVec3d(a, b, c)) // 6
	cube.AddPoint(mymath.MakeVec3d(a, b, 0)) // 7

	cube.Vertices[4].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)
	cube.Vertices[5].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)
	cube.Vertices[6].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)
	cube.Vertices[7].Normal = mymath.MakeVec3d(0.0, -1.0, 0.0)

	cube.AddPolygon(4, 5, 6, clr)
	cube.AddPolygon(6, 7, 4, clr)

	// front
	cube.AddPoint(mymath.MakeVec3d(0, b, c)) // 8
	cube.AddPoint(mymath.MakeVec3d(0, 0, c)) // 9
	cube.AddPoint(mymath.MakeVec3d(a, 0, c)) // 10
	cube.AddPoint(mymath.MakeVec3d(a, b, c)) // 11

	cube.Vertices[8].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)
	cube.Vertices[9].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)
	cube.Vertices[10].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)
	cube.Vertices[11].Normal = mymath.MakeVec3d(0.0, 0.0, 1.0)

	cube.AddPolygon(8, 9, 10, clr)
	cube.AddPolygon(10, 11, 8, clr)

	// back
	cube.AddPoint(mymath.MakeVec3d(0, 0, 0)) // 12
	cube.AddPoint(mymath.MakeVec3d(0, b, 0)) // 13
	cube.AddPoint(mymath.MakeVec3d(a, b, 0)) // 14
	cube.AddPoint(mymath.MakeVec3d(a, 0, 0)) // 15

	cube.Vertices[12].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)
	cube.Vertices[13].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)
	cube.Vertices[14].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)
	cube.Vertices[15].Normal = mymath.MakeVec3d(0.0, 0.0, -1.0)

	cube.AddPolygon(12, 13, 14, clr)
	cube.AddPolygon(14, 15, 12, clr)

	// left
	cube.AddPoint(mymath.MakeVec3d(0, b, 0)) // 16
	cube.AddPoint(mymath.MakeVec3d(0, 0, 0)) // 17
	cube.AddPoint(mymath.MakeVec3d(0, 0, c)) // 18
	cube.AddPoint(mymath.MakeVec3d(0, b, c)) // 19

	cube.Vertices[16].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)
	cube.Vertices[17].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)
	cube.Vertices[18].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)
	cube.Vertices[19].Normal = mymath.MakeVec3d(1.0, 0.0, 0.0)

	cube.AddPolygon(16, 17, 18, clr)
	cube.AddPolygon(18, 19, 16, clr)

	// right
	cube.AddPoint(mymath.MakeVec3d(a, b, c)) // 20
	cube.AddPoint(mymath.MakeVec3d(a, 0, c)) // 21
	cube.AddPoint(mymath.MakeVec3d(a, 0, 0)) // 22
	cube.AddPoint(mymath.MakeVec3d(a, b, 0)) // 23

	cube.Vertices[20].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)
	cube.Vertices[21].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)
	cube.Vertices[22].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)
	cube.Vertices[23].Normal = mymath.MakeVec3d(-1.0, 0.0, 0.0)

	cube.AddPolygon(20, 21, 22, clr)
	cube.AddPolygon(22, 23, 20, clr)

	cube.Move(mymath.Vec3dDiff(center, mymath.MakeVec3d(a/2, b/2, c/2)))
	return cube
}
