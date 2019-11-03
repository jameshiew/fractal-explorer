package viewport

import (
	"fmt"
)

const defaultScale = 0.01

type Viewport struct {
	scale, x, y float64
}

func New() Viewport {
	return Viewport{
		scale: defaultScale,
	}
}

func (v *Viewport) Move(deltaX, deltaY float64) {
	v.x += deltaX * v.scale
	v.y += deltaY * v.scale
}

func (v *Viewport) Up() {
	v.y++
}

func (v *Viewport) Left() {
	v.x--
}

func (v *Viewport) Right() {
	v.x++
}

func (v *Viewport) Down() {
	v.y--
}

func (v *Viewport) Zoom(factor float64) {
	v.scale *= factor
}

func (v *Viewport) String() string {
	return fmt.Sprintf("(%v, %v) @ %vx", v.x, v.y, v.scale)
}

func (v *Viewport) PixelToCartesian(pixelX, pixelY, width, height int) (x, y float64) {
	x, y = toVector(pixelX, pixelY, width, height)
	x *= v.scale
	y *= v.scale
	x += v.x
	y += v.y
	return x, y
}

func toVector(pixelX, pixelY, width, height int) (x, y float64) {
	return float64(pixelX - width/2), float64(-pixelY + height/2)
}
