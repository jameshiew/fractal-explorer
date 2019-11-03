package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"image/color"
)

const title = "Fractal Explorer"

func toCartesian(pixelX, pixelY, width, height int) (x, y int) {
	return pixelX - width/2, -pixelY + height/2
}

var myMandelbrot = mandelbrot{
	maxIterations: 50,
	bound:         2,
}

func colorForPixel(pixelX, pixelY, width, height int) color.Color {
	x, y := toCartesian(pixelX, pixelY, width, height)
	c := complex(float64(x)/100, float64(y)/100)
	if myMandelbrot.iterateWhileNotReachingBound(c) == myMandelbrot.maxIterations {
		return color.Black
	}
	return color.RGBA{
		R: 255,
		G: 100,
		B: 100,
		A: 255,
	}
}

// Run runs the application
func Run() {
	app := app.New()

	w := app.NewWindow(title)
	vp := &viewport{
		canvas: canvas.NewRasterWithPixels(colorForPixel),
	}
	w.SetContent(widget.NewVBox(
		fyne.NewContainerWithLayout(vp, vp.canvas),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
