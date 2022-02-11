package graphics

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/StephaneBunel/bresenham"
)

type Canvas interface {
	setPixel(x, y int, c color.Color)
	getPixel(x, y int) color.Color

	fill(c color.Color)
	drawLine(x1, y1, x2, y2 int, c color.Color)

	height() int
	width() int

	Image() image.Image
}

type ImageCanvas struct {
	Img *image.NRGBA
}

func MakeImageCanvas(h, w int) ImageCanvas {
	var img ImageCanvas
	img.Img = image.NewNRGBA(image.Rect(0, 0, w, h))
	return img
}

func (cnv ImageCanvas) setPixel(x, y int, c color.Color) {
	cnv.Img.Set(x, y, c)
}

func (cnv ImageCanvas) getPixel(x, y int) color.Color {
	return cnv.Img.At(x, y)
}

func (cnv ImageCanvas) drawLine(x1, y1, x2, y2 int, c color.Color) {
	bresenham.DrawLine(cnv.Img, x1, y1, x2, y2, c)
}

func (cnv ImageCanvas) fill(c color.Color) {
	draw.Draw(cnv.Img, cnv.Img.Bounds(), &image.Uniform{c}, image.Point{}, draw.Src)
}

func (cnv ImageCanvas) size() (int, int) {
	imageRect := cnv.Img.Bounds()
	return imageRect.Dx(), imageRect.Dy()
}

func (cnv ImageCanvas) height() int {
	return cnv.Img.Rect.Dy()
}

func (cnv ImageCanvas) width() int {
	return cnv.Img.Rect.Dx()
}

func (cnv ImageCanvas) Image() image.Image {
	return cnv.Img
}
