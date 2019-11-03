package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
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
	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		const zoomFactor = 1.1
		switch event.Name {
		case fyne.KeyQ:
			cnvs.drawer.scale.x /= zoomFactor
			cnvs.drawer.scale.y /= zoomFactor
			widget.Refresh(&cnvs)
		case fyne.KeyW:
			cnvs.drawer.scale.x *= zoomFactor
			cnvs.drawer.scale.y *= zoomFactor
			widget.Refresh(&cnvs)
		}
	})
	w.ShowAndRun()
}
