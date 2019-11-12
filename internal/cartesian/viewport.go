package cartesian

import (
	"fmt"
)

const defaultScale = 0.01

// Viewport is a view into a Cartesian plane
type Viewport struct {
	scale float64
	x, y  float64 // location
}

// NewViewport constructs a new Viewport with the default scale
func NewViewport() Viewport {
	return Viewport{
		scale: defaultScale,
	}
}

func (v *Viewport) Move(deltaX, deltaY float64) {
	v.x += deltaX * v.scale
	v.y += deltaY * v.scale
}

func (v *Viewport) Up() {
	v.Move(0, 1)
}

func (v *Viewport) Left() {
	v.Move(-1, 0)
}

func (v *Viewport) Right() {
	v.Move(1, 0)
}

func (v *Viewport) Down() {
	v.Move(0, -1)
}

func (v *Viewport) Zoom(factor float64) {
	// from https://stackoverflow.com/a/30410948
	newScale := factor * v.scale
	delta := newScale - v.scale
	v.scale = newScale
	v.Move(-(v.x * delta), -(v.y * delta))
}

func (v *Viewport) String() string {
	return fmt.Sprintf("(%v, %v) @ %vx", v.x, v.y, v.scale)
}

func (v *Viewport) PixelToCartesian(pixelX, pixelY, width, height int) (x, y float64) {
	x, y = Convert(pixelX, pixelY, width, height)
	x *= v.scale
	y *= v.scale
	x += v.x
	y += v.y
	return x, y
}
