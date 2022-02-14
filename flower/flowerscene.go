package flower

import (
	"image/color"
	"tulip/mymath"
	"tulip/primitives"
	"tulip/scene"
	// "tulip/graphics"
)

type TulipScene struct {
	scene.Scene

	center       mymath.Vec3
	centerNormal mymath.Vec3
	lightStep    mymath.Vec3

	tulips []Tulip
}

func NewTulipScene() *TulipScene {
	meadow := new(TulipScene)
	meadow.center = mymath.MakeVec3(0, 0, 0)
	meadow.centerNormal = mymath.MakeVec3(0, -1, 0)
	meadow.lightStep = mymath.MakeVec3(0, 0, 1)

	skyBlue := color.NRGBA{135, 206, 235, 255}
	meadow.SetBackground(skyBlue)

	// Ground
	green := color.NRGBA{0, 154, 23, 255}
	ground := primitives.NewBlock(10, 1, 5, mymath.MakeVec3(0, -0.5, 0), green)
	meadow.AddModel(ground)

	// Tulip(s)
	pink := color.NRGBA{255, 135, 141, 255}
	tulip1 := NewTulip(pink, mymath.MakeVec3(0, 0, 0), 0.2)
	meadow.tulips = append(meadow.tulips, *tulip1)
	meadow.AddComposite(meadow.tulips[0].Components)

	// yellow := color.NRGBA{251, 206, 43, 255}
	//tulip2 := flower.NewTulip(yellow, mymath.MakeVec3(0.3, -0.5, -0.3), 1, 0.03)
	// red := color.NRGBA{226, 34, 46, 255}
	//tulip3 := flower.NewTulip(red, mymath.MakeVec3(-20, 0, 220), 1, 2)

	// Light
	meadow.SetLight(1, mymath.MakeVec3(-10, 0, 0), mymath.MakeVec3(1, 0, 0.001))
	meadow.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), meadow.LightSource.Pos)
	meadow.LightSource.Direction.Normalize()

	// Camera
	cam := scene.MakeCamera(mymath.MakeVec3(0, 0, -10))
	meadow.SetCamera(cam)

	return meadow
}

func (meadow TulipScene) Animate(k float64) {
	if k > 1 {
		k = 1
	}
	if k < 0 {
		k = 0
	}

	for i := range meadow.tulips {
		meadow.tulips[i].OpenPetals(k)
	}
}

func (meadow *TulipScene) MoveSun() {
	meadow.LightSource.Pos.Rotate(meadow.center, meadow.lightStep)
	meadow.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), meadow.LightSource.Pos)
	meadow.LightSource.Direction.Normalize()

	k := meadow.LightSource.NormalIntensity(meadow.centerNormal)
	// fmt.Println(k)

	meadow.Animate(k)
}
