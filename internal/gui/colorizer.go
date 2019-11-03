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

func blend(colors ...color.Color) color.Color {
	var r, g, b, a uint32
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r += cr
		g += cg
		b += cb
		a += ca
	}
	return color.RGBA64{
		R: uint16(r / uint32(len(colors))),
		G: uint16(g / uint32(len(colors))),
		B: uint16(b / uint32(len(colors))),
		A: uint16(a / uint32(len(colors))),
	}
}

func forMandelbrot(base color.Color, fractal mandelbrot.Mandelbrot) colorizer {
	return func(c complex128) color.Color {
		iter := fractal.IterateWhileNotReachingBound(c)
		if iter == fractal.MaxIterations() {
			return color.Black
		}
		scale := float64(iter) / float64(fractal.MaxIterations())
		r, g, b, a := base.RGBA()
		return color.RGBA64{
			R: uint16(scale * float64(r)),
			G: uint16(scale * float64(g)),
			B: uint16(scale * float64(b)),
			A: uint16(a),
		}
	}
}
