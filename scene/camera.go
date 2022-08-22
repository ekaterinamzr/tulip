package scene

import (
	"tulip/mymath"
)

type Camera struct {
	Pos mymath.Vec3

	VCamera  mymath.Vec3
	VLookDir mymath.Vec4
	FYaw     float64
	FTheta   float64
	VTarget  mymath.Vec3

	VForward mymath.Vec3
}

func MakeCamera(pos mymath.Vec3) Camera {
	var c Camera
	c.VCamera = pos

	//vUp := mymath.MakeVec3(0, 1, 0)
	c.VTarget = mymath.MakeVec3(0, 0, 1)
	//mCameraRot := mymath.MakeRotationYM(c.FYaw)
	//c.VLookDir = mymath.MulVectorMatrix(c.VTarget, mCameraRot)
	//c.VTarget = mymath.Vec3Sum(c.VCamera, c.VLookDir.Vec3)

	c.VForward = mymath.Vec3Mul(c.VLookDir.Vec3, 1.0)

	return c
}

// func MakeCamera(pos mymath.Vec3, ar, hfov, speed float64) Camera {
// 	var c Camera

// 	c.Pos = pos
// 	c.Aspect_ratio = ar
// 	c.Hfov = 95.0
// 	c.Vfov = c.Hfov / c.Aspect_ratio

// 	// c.Htrack = c.Hfov / float64()

// 	c.Speed = speed

// 	return c
// }

func (c *Camera) SetPos(pos mymath.Vec3) {
	c.Pos = pos
}

func (c *Camera) Move(delta mymath.Vec3) {
	c.Pos.Move(delta)
	//c.Center.Move(delta)
}

