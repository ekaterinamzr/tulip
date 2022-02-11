package scene

import (
	"tulip/mymath"
)

type Camera struct {
	Pos mymath.Vec3
	Center mymath.Vec3 // камера смотрит в центр
	// VCamera  mymath.Vec3
	// VLookDir mymath.Vec3
	// FYaw     float64
	// FTheta   float64
	// VTarget  mymath.Vec3

	// VForward mymath.Vec3
}

func (c *Camera) SetPos(pos mymath.Vec3) {
	c.Pos = pos
}

func (c *Camera) Move(delta mymath.Vec3) {
	c.Pos.Move(delta)
	c.Center.Move(delta)
}

// func (c Camera) Pos() mymath.Vec3 {
// 	return c.pos
// }

// func MakeCamera(pos mymath.Vec3) Camera {
// 	var c Camera
// 	c.VCamera = mymath.MakeVec3(0, 0, 0)

// 	//vUp := mymath.MakeVec3(0, 1, 0)
// 	c.VTarget = mymath.MakeVec3(0, 0, 1)
// 	//mCameraRot := mymath.MakeRotationYM(c.FYaw)
// 	//c.VLookDir = mymath.MulVectorMatrix(c.VTarget, mCameraRot)
// 	c.VTarget = mymath.Vec3Sum(c.VCamera, c.VLookDir)

// 	c.VForward = mymath.Vec3Mul(c.VLookDir, 1.0)

// 	return c
// }
