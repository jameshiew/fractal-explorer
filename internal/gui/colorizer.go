package gui

import (
	"fractal-explorer/internal/mandelbrot"
	"image/color"
)

// colorizers color complex numbers
type colorizer func(complex128) color.Color

var (
	red   = color.RGBA64{R: 65535, A: 65535}
	green = color.RGBA64{G: 65535, A: 65535}
	blue  = color.RGBA64{B: 65535, A: 65535}
)

func forMandelbrot(base color.RGBA64, fractal mandelbrot.Mandelbrot) colorizer {
	return func(c complex128) color.Color {
		iter := fractal.IterateWhileNotReachingBound(c)
		if iter == fractal.MaxIterations() {
			return color.Black
		}
		scale := float64(iter) / float64(fractal.MaxIterations())
		return color.RGBA64{
			R: uint16(scale * float64(base.R)),
			G: uint16(scale * float64(base.G)),
			B: uint16(scale * float64(base.B)),
			A: base.A,
		}
	}
}
