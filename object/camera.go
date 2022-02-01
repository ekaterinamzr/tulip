package object

import (
	"tulip/mymath"
)

type Camera struct {
	VCamera  mymath.Vector3d
	VLookDir mymath.Vector3d
	FYaw     float64
	FTheta   float64
	VTarget  mymath.Vector3d

	VForward mymath.Vector3d
}

func MakeCamera() Camera {
	var c Camera
	c.VCamera = mymath.MakeVector3d(0, 0, 0)

	//vUp := mymath.MakeVector3d(0, 1, 0)
	c.VTarget = mymath.MakeVector3d(0, 0, 1)
	mCameraRot := mymath.MakeRotationYM(c.FYaw)
	c.VLookDir = mymath.MulVectorMatrix(c.VTarget, mCameraRot)
	c.VTarget = mymath.Vector3dSum(c.VCamera, c.VLookDir)

	c.VForward = mymath.Vector3dMul(c.VLookDir, 1.0)

	return c
}
