package fractal

import (
	"fyne.io/fyne/app"
)

const (
	title           = "Fractal Explorer"
	minWidthPixels  = 320
	minHeightPixels = 240
)

// Run runs the application
func Run() {
	app := app.New()

	w := app.NewWindow(title)
	cnvs := newFractalCanvas()
	w.SetContent(&cnvs)
	w.ShowAndRun()
}
