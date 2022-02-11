package main

import (
	"image/color"
	// "math"
	"time"
	// "tulip/flower"
	// "tulip/graphics"
	"tulip/primitives"
	"tulip/scene"

	"tulip/mymath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
)

const (
	width  = 500
	height = 500
)

func main() {
	a := app.New()
	w := a.NewWindow("Tulip")
	// ws := a.NewWindow("Tulip shadows")

	pink := color.NRGBA{255, 135, 141, 255}
	// yellow := color.NRGBA{251, 206, 43, 255}
	// red := color.NRGBA{226, 34, 46, 255}

	// tulip1 := flower.NewTulip(pink, mymath.MakeVec3(0, -0.5, 0), 1, 0.03)
	// tulip2 := flower.NewTulip(yellow, mymath.MakeVec3(0.3, -0.5, -0.3), 1, 0.03)
	//tulip3 := flower.NewTulip(red, mymath.MakeVec3(-20, 0, 220), 1, 2)

	cube := primitives.NewCube(0.5, mymath.MakeVec3(0, 0.25, 0), pink)

	var delay time.Duration = 50

	// var GrEngine graphics.ZBufferGraphicsEngine
	// GrEngine.Cnv = graphics.NewImageCanvas(height, width)

	//GrEngine := graphics.MakeZBuffEngine(graphics.MakeImageCanvas(height, width))
	// cnv := graphics.MakeImageCanvas(height, width)
	// engine := graphics.NewMyGrEngine()

	var scn scene.Scene
	scn.SetBackground(color.NRGBA{0, 204, 255, 255})
	scn.SetGroundClr(color.NRGBA{0, 154, 23, 255})
	scn.SetGround(mymath.MakeVec3(10, 0, 1))
	scn.SetLight(1, mymath.MakeVec3(10, 10, 10), mymath.MakeVec3(-1, -0.3, 0))
	ground := primitives.NewBlock(2, 0.5, 2, mymath.MakeVec3(0, -0.25, 0), scn.GroundClr)
	scn.AddObject(ground)
	// scn.AddObject(tulip1)
	// scn.AddObject(tulip2)
	scn.AddObject(cube)
	//scn.AddObject(tulip2)
	//scn.AddObject(tulip3)

	scn.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), scn.LightSource.Pos)
	scn.LightSource.Direction.Normalize()

	//cam := scene.MakeCamera()
	var cam scene.Camera
	cam.SetPos(mymath.MakeVec3(0, 0, 50))
	scn.SetCamera(cam)

	//scn.Camera.VCamera.Y += 21

	// w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
	// 	if k.Name == "Right" {
	// 		scn.Camera.VCamera.X += 1
	// 	}
	// 	if k.Name == "Left" {
	// 		scn.Camera.VCamera.X -= 1
	// 	}
	// 	if k.Name == "Up" {
	// 		scn.Camera.VCamera.Y += 1
	// 	}
	// 	if k.Name == "Down" {
	// 		scn.Camera.VCamera.Y -= 1
	// 	}

	// 	if k.Name == "W" {
	// 		scn.Camera.VCamera.Add(scn.Camera.VForward)
	// 	}
	// 	if k.Name == "S" {
	// 		scn.Camera.VCamera.Sub(scn.Camera.VForward)
	// 	}
	// 	if k.Name == "A" {
	// 		scn.Camera.FYaw += 0.1
	// 	}
	// 	if k.Name == "D" {
	// 		scn.Camera.FYaw -= 0.1
	// 	}
	// })

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "Right" {
			scn.Camera.Pos.X += 1
			// scn.Camera.Pos.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 3, 0))
			//scn.Camera.Center.X += 1
		}
		if k.Name == "Left" {
			scn.Camera.Pos.X -= 1
			//scn.Camera.Pos.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, -3, 0))
			//scn.Camera.Center.X -= 1
		}
		if k.Name == "Up" {
			scn.Camera.Pos.Y += 1
		}
		if k.Name == "Down" {
			scn.Camera.Pos.Y -= 1
		}

		if k.Name == "W" {
			scn.Camera.Pos.Z += 1
		}
		if k.Name == "S" {
			scn.Camera.Pos.Z -= 1
		}
	})

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * delay)

			// engine.RenderScene(&scn, true, true)
			// rastShadows := canvas.NewRasterFromImage(engine.Image())
			// ws.SetContent(rastShadows)

			// //engine.RenderScene(&scn, false, false)
			// rast := canvas.NewRasterFromImage(engine.Image())
			// w.SetContent(rast)

			//scn.Objects[0].Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(3, 3, 3))

			// scn.Objects[1].Animate(math.Abs(scene.VectorIntensity(mymath.MakeVec3(0, -1, 0), scn.LightSource)) / 0.6)
			//scn.Objects[1].Animate(math.Abs(scn.VectorIntensity(mymath.MakeVec3(0, -1, 0), scn.LightSource)) / 0.6)
			// scn.Objects[2].Animate(math.Abs(scene.VectorIntensity(mymath.MakeVec3(0, -1, 0), scn.LightSource)) / 0.6)

			// scn.LightSource.Intensity = -scn.LightSource.Direction.Y * 0.6
			// scn.SetBackground(scene.Lightness(color.NRGBA{0, 204, 255, 255}, scn.LightSource.Intensity))
			// scn.LightSource.Pos.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 0, 1))

			// scn.LightSource.Direction = mymath.Vec3Diff(mymath.MakeVec3(0, 0, 0), scn.LightSource.Pos)
			// scn.LightSource.Direction.Normalize()

			// scn.LightSource.Direction.Rotate(mymath.MakeVec3(0, 0, 0), mymath.MakeVec3(0, 0, -1))
			// if scn.LightSource.Pos.Y < 0 {
			// 	scn.LightSource.Pos.Y = 0
			// 	scn.LightSource.Pos.X = -10
			// }
		}
	}()

	w.Resize(fyne.NewSize(width, height))

	w.ShowAndRun()

	// ws.Resize(fyne.NewSize(width, height))

	// ws.ShowAndRun()
}
