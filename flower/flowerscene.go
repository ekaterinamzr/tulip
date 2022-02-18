package flower

import (
	"image/color"
	"math/rand"
	"time"
	"tulip/graphics"
	"tulip/mymath"
	"tulip/primitives"
	"tulip/scene"
)

type TulipScene struct {
	scene.Scene

	center                       mymath.Vec3
	centerNormal                 mymath.Vec3
	lightDayStep, lightNightStep mymath.Vec3

	minLight, maxLight float64

	skyBlue color.NRGBA

	night bool

	tulips []Tulip
}

func randomPos(xMin, xMax, zMin, zMax float64) mymath.Vec3 {
	var pos mymath.Vec3

	rand.Seed(time.Now().UnixNano())
	pos.X = xMin + rand.Float64()*(xMax-xMin)
	rand.Seed(time.Now().UnixNano())
	pos.Z = zMin + rand.Float64()*(zMax-zMin)
	pos.Y = 0

	return pos
}

func NewTulipScene(x1, x2, x3 float64) *TulipScene {
	meadow := new(TulipScene)
	meadow.center = mymath.MakeVec3(0, 0, 0)
	meadow.centerNormal = mymath.MakeVec3(0, -1, 0)
	meadow.lightDayStep = mymath.MakeVec3(0, 0, 0.1)
	meadow.lightDayStep = mymath.MakeVec3(0, 0, 3)
	meadow.minLight = 0.
	meadow.maxLight = 0.6
	meadow.night = false

	meadow.skyBlue = color.NRGBA{135, 206, 235, 255}
	meadow.SetBackground(meadow.skyBlue)

	// Ground
	groundWidth := 15.0
	groundDepth := 10.0
	green := color.NRGBA{0, 154, 23, 255}
	ground := primitives.NewBlock(groundWidth, 1, groundDepth, mymath.MakeVec3(0, -0.5, 0), green)
	meadow.AddModel(ground)

	// Tulip(s)
	pink := color.NRGBA{255, 135, 141, 255}
	tulip1 := NewTulip(pink, randomPos(-groundWidth/2.0, groundWidth/2.0, groundDepth/2.0, -groundDepth/2.0), 0.2, x1, x2, x3)
	meadow.tulips = append(meadow.tulips, *tulip1)
	meadow.AddComposite(meadow.tulips[0].Components)

	yellow := color.NRGBA{251, 206, 43, 255}
	tulip2 := NewTulip(yellow, randomPos(-groundWidth/2.0, groundWidth/2.0, groundDepth/2.0, -groundDepth/2.0), 0.2, x1, x2, x3)
	meadow.tulips = append(meadow.tulips, *tulip2)
	meadow.AddComposite(meadow.tulips[1].Components)

	// red := color.NRGBA{226, 34, 46, 255}
	white := color.NRGBA{243, 226, 192, 255}
	tulip3 := NewTulip(white, randomPos(-groundWidth/2.0, groundWidth/2.0, groundDepth/2.0, -groundDepth/2.0), 0.2, x1, x2, x3)
	meadow.tulips = append(meadow.tulips, *tulip3)
	meadow.AddComposite(meadow.tulips[2].Components)

	// Light
	meadow.SetLight(meadow.minLight, mymath.MakeVec3(-10, 0, 0), mymath.MakeVec3(1, 0, 0))
	meadow.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), meadow.LightSource.Pos)
	meadow.LightSource.Direction.Normalize()

	// Camera
	cam := scene.MakeCamera(mymath.MakeVec3(0, 5, -15))
	meadow.SetCamera(cam)

	return meadow
}

func NewPetalScene() *TulipScene {
	meadow := new(TulipScene)
	meadow.center = mymath.MakeVec3(0, 0, 0)
	meadow.centerNormal = mymath.MakeVec3(0, -1, 0)
	meadow.lightDayStep = mymath.MakeVec3(0, 0, 0.5)
	meadow.lightDayStep = mymath.MakeVec3(0, 0, 3)
	meadow.minLight = 0.2
	meadow.maxLight = 0.6
	meadow.night = false

	meadow.skyBlue = color.NRGBA{135, 206, 235, 255}
	meadow.SetBackground(meadow.skyBlue)

	var stage mymath.BezierCurve
	stage[0] = mymath.MakeVec3(0, 0, 0)
	stage[1] = mymath.MakeVec3(6, 0, 0)
	stage[2] = mymath.MakeVec3(6, 8, 0)
	stage[3] = mymath.MakeVec3(5, 10.907, 0)

	// Petal
	pink := color.NRGBA{255, 135, 141, 255}

	petal := MakePetal(stage, 10, 10, pink)
	petal.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 270, 0))

	meadow.AddModel(&petal)

	// Light
	meadow.SetLight(meadow.minLight, mymath.MakeVec3(-10, 0, 0), mymath.MakeVec3(1, 0, 0))
	meadow.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), meadow.LightSource.Pos)
	meadow.LightSource.Direction.Normalize()

	// Camera
	cam := scene.MakeCamera(mymath.MakeVec3(0, 5, -15))
	meadow.SetCamera(cam)

	return meadow
}

func NewTriangleScene() *TulipScene {
	meadow := new(TulipScene)
	meadow.center = mymath.MakeVec3(0, 0, 0)
	meadow.centerNormal = mymath.MakeVec3(0, -1, 0)
	meadow.lightDayStep = mymath.MakeVec3(0, 0, 0.5)
	meadow.lightDayStep = mymath.MakeVec3(0, 0, 3)
	meadow.minLight = 0.2
	meadow.maxLight = 0.6
	meadow.night = false

	meadow.skyBlue = color.NRGBA{135, 206, 235, 255}
	meadow.SetBackground(meadow.skyBlue)
	

	petal := primitives.NewTriangle()
	//petal.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 270, 0))

	meadow.AddModel(petal)

	// Light
	meadow.SetLight(meadow.minLight, mymath.MakeVec3(-10, 0, 0), mymath.MakeVec3(1, 0, 0))
	meadow.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), meadow.LightSource.Pos)
	meadow.LightSource.Direction.Normalize()

	// Camera
	cam := scene.MakeCamera(mymath.MakeVec3(0, 5, -15))
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
	if meadow.night {
		meadow.LightSource.Pos.Rotate(meadow.center, meadow.lightNightStep)
	} else {
		meadow.LightSource.Pos.Rotate(meadow.center, meadow.lightDayStep)
	}
	if meadow.LightSource.Pos.Y <= 0 {
		meadow.LightSource.Pos.X = -10
		meadow.night = !meadow.night
	}

	meadow.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), meadow.LightSource.Pos)
	meadow.LightSource.Direction.Normalize()

	k := 0.0
	if !meadow.night {
		meadow.LightSource.Intensity = meadow.LightSource.Direction.Y/-1.0*(meadow.maxLight-meadow.minLight) + meadow.minLight
		k = meadow.LightSource.NormalIntensity(meadow.centerNormal) / meadow.maxLight
	} else {
		meadow.LightSource.Intensity = meadow.minLight //meadow.LightSource.Direction.Y / -1.0 * meadow.minLight
	}

	meadow.SetBackground(graphics.Shade(meadow.skyBlue, 0.5-meadow.LightSource.Intensity))

	meadow.Animate(k)
}
