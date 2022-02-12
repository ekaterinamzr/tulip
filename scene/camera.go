package scene

import (
	"tulip/mymath"
)

type Camera struct {
	Pos mymath.Vec3
	Aspect_ratio float64
	Hfov, Vfov float64

	Htrack, Vtrack float64
	Speed float64
	//Center mymath.Vec3 // Camera is viewing center
}

func MakeCamera(pos mymath.Vec3, ar, hfov, speed float64) Camera {
	var c Camera

	c.Pos = pos
	c.Aspect_ratio = ar
	c.Hfov = 95.0
	c.Vfov = c.Hfov / c.Aspect_ratio

	// c.Htrack = c.Hfov / float64()

	c.Speed = speed

	return c
}

func (c *Camera) SetPos(pos mymath.Vec3) {
	c.Pos = pos
}

func (c *Camera) Move(delta mymath.Vec3) {
	c.Pos.Move(delta)
	//c.Center.Move(delta)
}

