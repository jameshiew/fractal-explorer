package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

const title = "Fractal Explorer"

// Run runs the application
func Run() {
	app := app.New()

	w := app.NewWindow(title)
	drwr := &drawer{
		scale: struct {
			x, y float64
		}{
			x: 0.01,
			y: 0.01,
		},
		mandelbrot: mandelbrot{
			maxIterations: 50,
			bound:         2,
		},
	}
	vp := &viewport{
		canvas: canvas.NewRasterWithPixels(drwr.pixelColor),
	}
	w.SetContent(fyne.NewContainerWithLayout(vp, vp.canvas))

	w.ShowAndRun()
}
