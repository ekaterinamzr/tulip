package main

import (
	//"fmt"
	"fmt"
	"image/color"

	// "image/draw"

	"time"
	"tulip/flower"
	"tulip/graphics"
	"tulip/object"

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

	mygreen := color.RGBA{0, 100, 0, 255}

	v := []object.Vertex{{0, 0, 0}, {0, 10, 0}, {0, 0, 10}, {10, 0, 0}, {10, 10, 0}, {10, 0, 10}, {0, 10, 10}, {10, 10, 10}}
	p := []object.Polygon{{2, 6, 5, mygreen}, {6, 7, 5, mygreen}, {5, 7, 3, mygreen}, {7, 4, 3, mygreen}, {4, 3, 0, mygreen}, {4, 1, 0, mygreen}, {1, 2, 0, mygreen}, {1, 6, 0, mygreen}, {1, 6, 7, mygreen}, {1, 4, 7, mygreen}, {0, 2, 5, mygreen}, {0, 3, 5, mygreen}}
	cube := object.Model{v, p}

	v1 := []object.Vertex{{50, 50, 0}, {100, 100, 0}, {80, 20, 0}}
	p1 := []object.Polygon{{0, 1, 2, mygreen}}
	triangle := object.Model{v1, p1}
	triangle.SortVertices()

	cube.SortVertices()

	petal1 := flower.TestPetal()
	petal2 := flower.TestPetal()
	petal2.Move(object.Vertex{50, 50, 50})

	leaf := flower.TestLeaf()
	fmt.Println(leaf)

	var comp object.CompositeModel
	comp.Add(&petal1)
	comp.Add(&petal2)

	//fmt.Println(comp.Components)

	//petal.SortVertices()

	tulip := flower.NewTulip(object.Vertex{100, 100, 0})
	//fmt.Println(tulip.Components)

	//myimage := image.NewRGBA(image.Rect(0, 0, 220, 220)) // x1,y1,  x2,y2
	//  R, G, B, Alpha

	// backfill entire surface with green
	//draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

	var delay time.Duration = 50

	var GrEngine graphics.ZBufferGraphicsEngine
	GrEngine.Cnv = graphics.NewImageCanvas(height, width)

	cube.Scale(object.Vertex{0, 0, 0}, 10)
	cube.Move(object.Vertex{100, 100, 0})
	center := object.Vertex{5, 5, 5}
	center.ScaleVertex(object.Vertex{0, 0, 0}, 10)

	center.MoveVertex(object.Vertex{100, 100, 0})

	var scene object.Scene
	scene.BackGround = color.NRGBA{255, 255, 255, 255}
	scene.AddObject(tulip)

	//red := color.NRGBA{255, 0, 0, 255}

	//scene.Objects[0].RotateObject(center, object.Vertex{30, 30, 30})

	go func() {
		//time.Sleep(time.Millisecond * 500)
		for i := 0; i < 500; i++ {
			time.Sleep(time.Millisecond * delay)

			//scene.BackGround = object.Lightness(red, float32(i)/255)

			GrEngine.RenderScene(scene)

			rast := canvas.NewRasterFromImage(GrEngine.Image())
			w.SetContent(rast)

			scene.Objects[0].Rotate(object.Vertex{100, 300, 100}, object.Vertex{0, 3, 0})
		}
	}()

	w.Resize(fyne.NewSize(width, height))

	w.ShowAndRun()
}
