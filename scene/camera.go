package scene

import (
	"tulip/mymath"
)

type Camera struct {
	Pos mymath.Vec3
	Center mymath.Vec3 // Camera is viewing center
}

func (c *Camera) SetPos(pos mymath.Vec3) {
	c.Pos = pos
}

func (c *Camera) Move(delta mymath.Vec3) {
	c.Pos.Move(delta)
	c.Center.Move(delta)
}

