package gui

import (
	"gitlab.com/jameshiew/fractal-explorer/internal/mandelbrot"
	"image/color"
	"math"
)

// colorizers color complex numbers
type colorizer func(complex128) color.Color

var (
	red   = color.RGBA64{R: 65535, A: 65535}
	green = color.RGBA64{G: 65535, A: 65535}
	blue  = color.RGBA64{B: 65535, A: 65535}
)

// darkBlend is quite dark
func darkBlend(z complex128) color.Color {
	return blend(
		newColorizer(green, mandelbrot.NewImageBuilder().SetMaxIterations(125).Build())(z),
		newColorizer(blue, mandelbrot.NewImageBuilder().SetMaxIterations(250).Build())(z),
		newColorizer(red, mandelbrot.NewImageBuilder().SetMaxIterations(500).Build())(z),
	)
}

func otherBlend(z complex128) color.Color {
	return blend(
		newColorizer(green, mandelbrot.NewImageBuilder().SetMaxIterations(120).SetBound(math.Phi).Build())(z),
		newColorizer(color.RGBA64{
			R: 20000,
			G: 50000,
			B: 20000,
			A: 65535,
		}, mandelbrot.NewImageBuilder().SetMaxIterations(100).SetBound(math.E).Build())(z),
		newColorizer(color.RGBA64{
			R: 16000,
			G: 65335,
			B: 16000,
			A: 65535,
		}, mandelbrot.NewImageBuilder().SetMaxIterations(75).SetBound(math.Pi).Build())(z),
	)
}

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

type colorizable interface {
	MaxIterations() uint16
	IterateWhileNotReachingBound(complex128) uint16
}

func newColorizer(base color.Color, colorizable colorizable) colorizer {
	return func(c complex128) color.Color {
		iter := colorizable.IterateWhileNotReachingBound(c)
		if iter == colorizable.MaxIterations() {
			return color.Black
		}
		scale := float64(iter) / float64(colorizable.MaxIterations())
		r, g, b, a := base.RGBA()
		return color.RGBA64{
			R: uint16(scale * float64(r)),
			G: uint16(scale * float64(g)),
			B: uint16(scale * float64(b)),
			A: uint16(a),
		}
	}
}
