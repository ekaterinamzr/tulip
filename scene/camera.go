package scene

import (
	"tulip/mymath"
)

type Camera struct {
	Pos mymath.Vec3d
	Center mymath.Vec3d // камера смотрит в центр
	// VCamera  mymath.Vec3d
	// VLookDir mymath.Vec3d
	// FYaw     float64
	// FTheta   float64
	// VTarget  mymath.Vec3d

	// VForward mymath.Vec3d
}

func (c *Camera) SetPos(pos mymath.Vec3d) {
	c.Pos = pos
}

func (c *Camera) Move(delta mymath.Vec3d) {
	c.Pos.Move(delta)
	c.Center.Move(delta)
}

// func (c Camera) Pos() mymath.Vec3d {
// 	return c.pos
// }

// func MakeCamera(pos mymath.Vec3d) Camera {
// 	var c Camera
// 	c.VCamera = mymath.MakeVec3d(0, 0, 0)

// 	//vUp := mymath.MakeVec3d(0, 1, 0)
// 	c.VTarget = mymath.MakeVec3d(0, 0, 1)
// 	//mCameraRot := mymath.MakeRotationYM(c.FYaw)
// 	//c.VLookDir = mymath.MulVectorMatrix(c.VTarget, mCameraRot)
// 	c.VTarget = mymath.Vec3dSum(c.VCamera, c.VLookDir)

// 	c.VForward = mymath.Vec3dMul(c.VLookDir, 1.0)

// 	return c
// }
