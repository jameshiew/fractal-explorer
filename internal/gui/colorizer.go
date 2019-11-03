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
		return color.RGBA{
			R: uint8(scale * 255),
			G: uint8(scale * 100),
			B: uint8(scale * 100),
			A: 255,
		}
	}
}
