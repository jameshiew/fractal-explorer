package fractal

import (
	"fmt"
	"image/color"
)

type viewport struct {
	scale  float64
	center struct {
		x, y float64
	}
	mandelbrot mandelbrot
}

func (v *viewport) String() string {
	return fmt.Sprintf("%v - (%v, %v) @ %vx", v.mandelbrot.String(), v.center.x, v.center.y, v.scale)
}

func (d *viewport) pixelColor(pixelX, pixelY, width, height int) color.Color {
	x, y := d.toCartesian(pixelX, pixelY, width, height)
	c := complex(x*d.scale, y*d.scale)
	iter := d.mandelbrot.iterateWhileNotReachingBound(c)
	if iter == d.mandelbrot.maxIterations {
		return color.Black
	}
	scale := float64(iter) / float64(d.mandelbrot.maxIterations)
	return color.RGBA{
		R: uint8(scale * 255),
		G: uint8(scale * 100),
		B: uint8(scale * 100),
		A: 255,
	}
}

func (d *viewport) toCartesian(pixelX, pixelY, width, height int) (x, y float64) {
	return d.center.x + float64(pixelX-width/2), d.center.y + float64(-pixelY+height/2)
}
