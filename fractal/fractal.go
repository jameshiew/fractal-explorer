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
		const zoomIncrement = 0.001
		switch event.Name {
		case fyne.KeyUp:
			cnvs.drawer.center.y += 1
		case fyne.KeyLeft:
			cnvs.drawer.center.x -= 1
		case fyne.KeyRight:
			cnvs.drawer.center.x += 1
		case fyne.KeyDown:
			cnvs.drawer.center.y -= 1
		case fyne.KeyW:
			cnvs.drawer.scale -= zoomIncrement
		case fyne.KeyS:
			cnvs.drawer.scale += zoomIncrement
		case fyne.KeyA:
			cnvs.drawer.mandelbrot.maxIterations--
		case fyne.KeyD:
			cnvs.drawer.mandelbrot.maxIterations++
		case fyne.KeyQ:
			cnvs.drawer.mandelbrot.bound--
		case fyne.KeyE:
			cnvs.drawer.mandelbrot.bound++
		}
		widget.Refresh(&cnvs)
	})
	w.RequestFocus()
	w.ShowAndRun()
}
