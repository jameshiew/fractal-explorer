package gui

import (
	"fractal-explorer/internal/mandelbrot"
	"image/color"
)

// colorizers color complex numbers
type colorizer func(complex128) color.Color

func forMandelbrot(fractal mandelbrot.Mandelbrot) colorizer {
	return func(c complex128) color.Color {
		iter := fractal.IterateWhileNotReachingBound(c)
		if iter == fractal.MaxIterations() {
			return color.Black
		}
		scale := float64(iter) / float64(fractal.MaxIterations())
		return color.NRGBA64{
			R: uint16(scale * 65535),
			G: uint16(scale * 16000),
			B: uint16(scale * 16000),
			A: 65535,
		}
	}
}
