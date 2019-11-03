package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

const title = "Fractal Explorer"

// Run runs the application
func Run() {
	app := app.New()

	w := app.NewWindow(title)
	mandel := mandelbrot{
		maxIterations: 50,
		bound:         2,
	}
	drwr := &drawer{mandel}
	vp := &viewport{
		canvas: canvas.NewRasterWithPixels(drwr.pixelColor),
	}
	w.SetContent(widget.NewVBox(
		fyne.NewContainerWithLayout(vp, vp.canvas),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
