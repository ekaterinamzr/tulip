package main

import (
	//"fmt"
	//"fmt"

	"image/color"

	// "image/draw"

	"math"
	"time"
	"tulip/flower"
	"tulip/graphics"
	"tulip/object"

	//"tulip/primitives"

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

	// blue := color.NRGBA{0, 0, 255, 255}
	// center := object.MakePoint(100, 100, 100)
	// cube := primitives.NewCube(100, center, blue)

	petal1 := flower.TestPetal()
	petal2 := flower.TestPetal()
	petal2.Move(object.Point{50, 50, 50})

	//leaf := flower.TestLeaf()
	//fmt.Println(leaf)

	var comp object.CompositeModel
	comp.Add(&petal1)
	comp.Add(&petal2)

	//fmt.Println(comp.Components)

	//petal.SortVertices()

	tulip1 := flower.NewTulip(object.Point{100, 0, 0}, 1)
	tulip2 := flower.NewTulip(object.Point{300, 0, 0}, 2)
	//fmt.Println(tulip.Components)

	//myimage := image.NewRGBA(image.Rect(0, 0, 220, 220)) // x1,y1,  x2,y2
	//  R, G, B, Alpha

	// backfill entire surface with green
	//draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

	var delay time.Duration = 50

	var GrEngine graphics.ZBufferGraphicsEngine
	GrEngine.Cnv = graphics.NewImageCanvas(height, width)

	var scene object.Scene
	scene.SetBackground(color.NRGBA{0, 204, 255, 255})
	scene.SetLight(0.5, object.Point{200, 400, 200})
	scene.AddObject(tulip1)
	scene.AddObject(tulip2)

	//red := color.NRGBA{255, 0, 0, 255}

	//scene.Objects[0].RotateObject(center, object.Vertex{30, 30, 30})
	//fmt.Println(tulip.Components[0])

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		scene.Objects[0].Rotate(object.Point{100, 0, 0}, object.Point{0, 10, 0})
	})

	go func() {
		//time.Sleep(time.Millisecond * 500)
		for i := 0; i < 500; i++ {
			time.Sleep(time.Millisecond * delay)

			//scene.BackGround = object.Lightness(red, float32(i)/255)

			GrEngine.RenderScene(scene)

			rast := canvas.NewRasterFromImage(GrEngine.Image())
			w.SetContent(rast)

			//scene.Objects[1].Rotate(object.Point{300, 0, 0}, object.Point{0, 3, 0})

			scene.Objects[0].Animate(math.Abs(object.CalculateIntensityVector(object.Make(0, -1, 0), scene.LightSource)) / 0.6)
			//scene.Objects[1].Animate(1 - float64(i)/100)

			scene.LightSource.Intensity = -scene.LightSource.Direction.Y * 0.6
			scene.SetBackground(object.Lightness(color.NRGBA{0, 204, 255, 255}, scene.LightSource.Intensity))
			scene.LightSource.Direction.Rotate(object.MakePoint(0, 0, 0), object.MakePoint(0, 0, -1))
			if scene.LightSource.Direction.Y > 0 {
				scene.LightSource.Direction.Y = 0
				scene.LightSource.Direction.X = -1
			}

			//scene.Objects[0].Rotate(center, object.Point{3, 3, 3})
			//fmt.Println(tulip.Components[0].Vertices)
		}
	}()

	w.Resize(fyne.NewSize(width, height))

	w.ShowAndRun()
}
