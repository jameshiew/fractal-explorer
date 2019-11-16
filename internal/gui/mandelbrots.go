package gui

import (
	"image/color"
	"math"

	"github.com/jameshiew/fractal-explorer/internal/draw"
	"github.com/jameshiew/fractal-explorer/internal/mandelbrot"
)

var (
	red   = color.RGBA64{R: 65535, A: 65535}
	green = color.RGBA64{G: 65535, A: 65535}
	blue  = color.RGBA64{B: 65535, A: 65535}
)

// darkBlend is quite dark
func darkBlend(z complex128) color.Color {
	return draw.Blend(
		draw.NewColorizer(green, mandelbrot.NewImageBuilder().SetMaxIterations(125).Build())(z),
		draw.NewColorizer(blue, mandelbrot.NewImageBuilder().SetMaxIterations(250).Build())(z),
		draw.NewColorizer(red, mandelbrot.NewImageBuilder().SetMaxIterations(500).Build())(z),
	)
}

func otherBlend(z complex128) color.Color {
	return draw.Blend(
		draw.NewColorizer(green, mandelbrot.NewImageBuilder().SetMaxIterations(120).SetBound(math.Phi).Build())(z),
		draw.NewColorizer(color.RGBA64{
			R: 20000,
			G: 50000,
			B: 20000,
			A: 65535,
		}, mandelbrot.NewImageBuilder().SetMaxIterations(100).SetBound(math.E).Build())(z),
		draw.NewColorizer(color.RGBA64{
			R: 16000,
			G: 65335,
			B: 16000,
			A: 65535,
		}, mandelbrot.NewImageBuilder().SetMaxIterations(75).SetBound(math.Pi).Build())(z),
	)
}
