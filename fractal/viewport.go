package fractal

import (
	"fmt"
	"fractal-explorer/fractal/mandelbrot"
	"image/color"
)

type complexColorer func(complex128) color.Color

type viewport struct {
	colorer complexColorer
	scale   float64
	center  struct {
		x, y float64
	}
}

func (v *viewport) String() string {
	return fmt.Sprintf("(%v, %v) @ %vx", v.center.x, v.center.y, v.scale)
}

func forMandelbrot(fractal mandelbrot.Mandelbrot) complexColorer {
	return func(c complex128) color.Color {
		iter := fractal.IterateWhileNotReachingBound(c)
		if iter == fractal.MaxIterations() {
			return color.Black
		}
		scale := float64(iter) / float64(fractal.MaxIterations())
		return color.RGBA{
			R: uint8(scale * 255),
			G: uint8(scale * 100),
			B: uint8(scale * 100),
			A: 255,
		}
	}
}

func (v *viewport) pixelColor(pixelX, pixelY, width, height int) color.Color {
	x, y := toCartesian(pixelX, pixelY, width, height)
	x *= v.scale
	y *= v.scale
	x += v.center.x
	y += v.center.y
	c := complex(x, y)
	return v.colorer(c)
}

func toCartesian(pixelX, pixelY, width, height int) (x, y float64) {
	return float64(pixelX - width/2), float64(-pixelY + height/2)
}
