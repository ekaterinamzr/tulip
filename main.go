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
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	width  = 800
	height = 800
)

func main() {
	a := app.New()
	w := a.NewWindow("Tulip")

	var delay time.Duration = 50

	cnv := graphics.MakeImageCanvas(height, width)
	engine := graphics.NewMyGrEngine(cnv)

	info := widget.NewLabel("Управление камерой: \n^ - вверх\nv - вниз\n> - вправо\n< - влево\nW - вперед\nS - назад\nD - поворот вправо\nA - поворот влево")
	infoBezier := widget.NewLabel("Модель открытого лепестка:\nМодель открытого\nлепестка строится на\nоснове кривой Безье\nс контрольными точками:\n(0, 0), (x1, 0), (x2, 8), (x3, 11).\nЗначения x1, x2, x3\nможно изменить, используя\nсоответствующие ползунки:\n")
	x1 := widget.NewLabel("x1")
	x2 := widget.NewLabel("x2")
	x3 := widget.NewLabel("x3")
	r := widget.NewLabel("0                                              15")
	slider1 := widget.NewSlider(0.0, 15.0)
	slider2 := widget.NewSlider(0.0, 15.0)
	slider3 := widget.NewSlider(0.0, 15.0)

	slider1.Value = 6
	slider2.Value = 6
	slider3.Value = 5

	meadow := flower.NewTulipScene(slider1.Value, slider2.Value, slider3.Value)
	// meadow := flower.NewTriangleScene()

	button := widget.NewButton("Сгенерировать сцену", func() {
		meadow = flower.NewTulipScene(slider1.Value, slider2.Value, slider3.Value)
	})

	menu := container.New(layout.NewVBoxLayout(), info, infoBezier, x3, slider3, x2, slider2, x1, slider1, r, button)
	menuColumn := container.New(layout.NewGridWrapLayout(fyne.NewSize(220, height)), menu)

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
		for {
			time.Sleep(time.Millisecond * delay)

			engine.RenderScene(&meadow.Scene)

			rast := canvas.NewRasterFromImage(cnv.Image())

			img := container.New(layout.NewGridWrapLayout(fyne.NewSize(width, height)), rast)
			form := container.New(layout.NewFormLayout(), menuColumn, img)

			w.SetContent(form)

			meadow.MoveSun()

		}
	}()

	w.Resize(fyne.NewSize(width+50, height))
	w.ShowAndRun()
}
