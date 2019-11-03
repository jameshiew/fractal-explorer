package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"image/color"
)

const title = "Fractal Explorer"

type viewport struct {
	canvas fyne.CanvasObject
}

func (v *viewport) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	v.canvas.Resize(size)
}

func (v *viewport) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(320, 240)
}

func stubPixelColor(px, py, w, h int) color.Color {
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
		canvas: canvas.NewRasterWithPixels(stubPixelColor),
	}
	w.SetContent(widget.NewVBox(
		fyne.NewContainerWithLayout(vp, vp.canvas),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
