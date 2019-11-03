package gui

import (
	"fractal-explorer/internal/mandelbrot"
	"image/color"
)

// colorizers color complex numbers
type colorizer func(complex128) color.Color

var (
	red   = color.RGBA64{65336, 16000, 16000, 65535}
	other = color.RGBA64{12345, 23456, 34567, 65535}
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
