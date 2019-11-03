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
		case fyne.KeyUp:
			cnvs.drawer.position.y += 1
		case fyne.KeyLeft:
			cnvs.drawer.position.x -= 1
		case fyne.KeyRight:
			cnvs.drawer.position.x += 1
		case fyne.KeyDown:
			cnvs.drawer.position.y -= 1
		case fyne.KeyW:
			cnvs.drawer.scale /= zoomFactor
		case fyne.KeyS:
			cnvs.drawer.scale *= zoomFactor
		}
		widget.Refresh(&cnvs)
	})
	w.RequestFocus()
	w.ShowAndRun()
}
