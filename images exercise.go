//Define your own Image type, implement the necessary methods, and call pic.ShowImage.

package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
	color uint8
}

func (p *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, p.w, p.h)
}

func (p *Image) At(x, y int) color.Color {
	return color.RGBA{p.color + uint8(x^y), p.color + uint8(y^x), 255, 255}
}

func main() {
	m := Image{128, 128, 0}
	pic.ShowImage(&m)
}
