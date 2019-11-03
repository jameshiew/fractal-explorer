package viewport

import (
	"fmt"
)

const (
	defaultScale  = 0.01
	zoomIncrement = 0.001
)

type Viewport struct {
	scale  float64
	center struct {
		x, y float64
	}
}

func New() Viewport {
	return Viewport{
		scale: defaultScale,
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

func (v *Viewport) Zoom(factor float64) {
	v.scale *= factor
}

func (v *Viewport) String() string {
	return fmt.Sprintf("(%v, %v) @ %vx", v.center.x, v.center.y, v.scale)
}

func (v *Viewport) PixelToComplex(pixelX, pixelY, width, height int) complex128 {
	x, y := toVector(pixelX, pixelY, width, height)
	x *= v.scale
	y *= v.scale
	x += v.center.x
	y += v.center.y
	return complex(x, y)
}

func toVector(pixelX, pixelY, width, height int) (x, y float64) {
	return float64(pixelX - width/2), float64(-pixelY + height/2)
}
