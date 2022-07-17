package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"github.com/jameshiew/fractal-explorer/internal/draw"
	"github.com/jameshiew/fractal-explorer/internal/mandelbrot"
)

func (f *fractalWidget) CreateRenderer() fyne.WidgetRenderer {
	if f.renderer == nil {
		raster := canvas.NewRaster(draw.New(f.log, func(pixelX, pixelY, width, height int) color.Color {
			x, y := f.viewport.PixelToCartesian(pixelX, pixelY, width, height)
			z := complex(x, y)
			return draw.NewColorizer(green, mandelbrot.NewImageBuilder().SetMaxIterations(70).Build())(z)
		}))
		f.renderer = newWidgetRenderer(raster, f.Refresh)
	}
	return f.renderer
}
