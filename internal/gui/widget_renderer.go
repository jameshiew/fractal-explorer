package gui

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"

	"gitlab.com/jameshiew/fractal-explorer/internal/draw"
	"gitlab.com/jameshiew/fractal-explorer/internal/mandelbrot"
)

func (f *fractalWidget) CreateRenderer() fyne.WidgetRenderer {
	raster := canvas.NewRaster(draw.New(func(pixelX, pixelY, width, height int) color.Color {
		x, y := f.viewport.PixelToCartesian(pixelX, pixelY, width, height)
		z := complex(x, y)
		return draw.NewColorizer(green, mandelbrot.NewImageBuilder().SetMaxIterations(70).Build())(z)
	}))
	return newWidgetRenderer(raster, f.refresh)
}
