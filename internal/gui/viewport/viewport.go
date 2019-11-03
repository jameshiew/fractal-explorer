package viewport

import (
	"fmt"
	"image/color"
)

const zoomIncrement = 0.001

type complexColorer func(complex128) color.Color

type Viewport struct {
	colorer complexColorer
	scale   float64
	center  struct {
		x, y float64
	}
}

func New(colorer complexColorer) Viewport {
	return Viewport{
		scale:   0.01,
		colorer: colorer,
	}
}

func (v *Viewport) Move(deltaX, deltaY float64) {
	v.center.x += deltaX * v.scale
	v.center.y += deltaY * v.scale
}

func (v *Viewport) Up() {
	v.center.y++
}

func (v *Viewport) Left() {
	v.center.x--
}

func (v *Viewport) Right() {
	v.center.x++
}

func (v *Viewport) Down() {
	v.center.y--
}

func (v *Viewport) ZoomIn() {
	v.scale -= zoomIncrement
}

func (v *Viewport) ZoomOut() {
	v.scale += zoomIncrement
}

func (v *Viewport) String() string {
	return fmt.Sprintf("(%v, %v) @ %vx", v.center.x, v.center.y, v.scale)
}

func (v *Viewport) PixelColor(pixelX, pixelY, width, height int) color.Color {
	x, y := toVector(pixelX, pixelY, width, height)
	x *= v.scale
	y *= v.scale
	x += v.center.x
	y += v.center.y
	c := complex(x, y)
	return v.colorer(c)
}

func toVector(pixelX, pixelY, width, height int) (x, y float64) {
	return float64(pixelX - width/2), float64(-pixelY + height/2)
}
