package main

import (
	//"fmt"
	//"fmt"

	"image/color"
	"math"

	// "image/draw"

	"time"
	"tulip/flower"
	"tulip/graphics"
	"tulip/object"

	"tulip/mymath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const (
	width  = 500
	height = 500
)

func main() {
	a := app.New()
	w := a.NewWindow("Tulip")

	tulip1 := flower.NewTulip(mymath.MakeVector3d(0, 0, 150), 1, 2)
	tulip2 := flower.NewTulip(mymath.MakeVector3d(50, 0, 200), 2, 2)

	var delay time.Duration = 50

	var GrEngine graphics.ZBufferGraphicsEngine
	GrEngine.Cnv = graphics.NewImageCanvas(height, width)

	var scene object.Scene
	scene.SetBackground(color.NRGBA{0, 204, 255, 255})
	scene.SetGroundClr(color.NRGBA{0, 154, 23, 255})
	scene.SetGround(mymath.MakeVector3d(100, 0, 100))
	scene.SetLight(0.5, mymath.MakeVector3d(200, 400, 200), mymath.MakeVector3d(-1, 0, 0))
	scene.AddObject(tulip1)
	scene.AddObject(tulip2)

	cam := object.MakeCamera()
	scene.SetCamera(cam)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "Right" {
			scene.Camera.VCamera.X += 1
		}
		if k.Name == "Left" {
			scene.Camera.VCamera.X -= 1
		}
		if k.Name == "Up" {
			scene.Camera.VCamera.Y += 1
		}
		if k.Name == "Down" {
			scene.Camera.VCamera.Y -= 1
		}

		if k.Name == "W" {
			scene.Camera.VCamera.Add(scene.Camera.VForward)
		}
		if k.Name == "S" {
			scene.Camera.VCamera.Sub(scene.Camera.VForward)
		}
		if k.Name == "A" {
			scene.Camera.FYaw -= 0.1
		}
		if k.Name == "D" {
			scene.Camera.FYaw += 0.1
		}
		//scene.Objects[0].Rotate(mymath.MakeVector3d(0, 0, 100}, mymath.MakeVector3d(0, 10, 0})
	})

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * delay)

			GrEngine.RenderScene(scene)

			rast := canvas.NewRasterFromImage(GrEngine.Image())
			w.SetContent(rast)

			scene.Objects[0].Animate(math.Abs(object.VectorIntensity(mymath.MakeVector3d(0, -1, 0), scene.LightSource)) / 0.6)

			scene.LightSource.Intensity = -scene.LightSource.Direction.Y * 0.6
			scene.SetBackground(object.Lightness(color.NRGBA{0, 204, 255, 255}, scene.LightSource.Intensity))
			scene.LightSource.Direction.Rotate(mymath.MakeVector3d(0, 0, 0), mymath.MakeVector3d(0, 0, -1))
			if scene.LightSource.Direction.Y > 0 {
				scene.LightSource.Direction.Y = 0
				scene.LightSource.Direction.X = -1
			}
		}
	}()

	w.Resize(fyne.NewSize(width, height))

	w.ShowAndRun()
}
