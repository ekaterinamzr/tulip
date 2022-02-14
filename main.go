package main

import (

	// "math"
	"time"
	"tulip/flower"
	"tulip/graphics"
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

	var delay time.Duration = 50

	cnv := graphics.MakeImageCanvas(height, width)
	engine := graphics.NewMyGrEngine(cnv)

	meadow := flower.NewTulipScene()

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		if k.Name == "Right" {
			meadow.MoveCamera(mymath.MakeVec3(1, 0, 0))
		}
		if k.Name == "Left" {
			meadow.MoveCamera(mymath.MakeVec3(-1, 0, 0))
		}
		if k.Name == "Up" {
			meadow.MoveCamera(mymath.MakeVec3(0, 1, 0))
		}
		if k.Name == "Down" {
			meadow.MoveCamera(mymath.MakeVec3(0, -1, 0))
		}

		if k.Name == "W" {
			meadow.MoveCameraForward()
		}
		if k.Name == "S" {
			meadow.MoveCameraBackward()
		}
		if k.Name == "A" {
			meadow.RotateCameraLeft()
		}
		if k.Name == "D" {
			meadow.RotateCameraRight()
		}
	})

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * delay)

			engine.RenderScene(&meadow.Scene)
			rast := canvas.NewRasterFromImage(cnv.Image())
			w.SetContent(rast)

			// scn.Objects[0].Animate(math.Abs(graphics.Intensity(mymath.MakeVec3(0, -1, 0), scn.LightSource)) / 0.6)
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

}
