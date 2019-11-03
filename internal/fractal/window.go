package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// Window gets a window containing the GUI
func Window(title string) fyne.Window {
	window := app.New().NewWindow(title)
	setUpWindow(window)
	return window
}

func setUpWindow(window fyne.Window) {
	cnvs := newFractalCanvas()
	window.SetContent(
		widget.NewVBox(
			&cnvs,
			cnvs.labels.info,
		),
	)
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		const zoomIncrement = 0.001
		switch event.Name {
		case fyne.KeyUp:
			cnvs.viewport.center.y += 1
		case fyne.KeyLeft:
			cnvs.viewport.center.x -= 1
		case fyne.KeyRight:
			cnvs.viewport.center.x += 1
		case fyne.KeyDown:
			cnvs.viewport.center.y -= 1
		case fyne.KeyW:
			cnvs.viewport.scale -= zoomIncrement
		case fyne.KeyS:
			cnvs.viewport.scale += zoomIncrement
		}
		widget.Refresh(&cnvs)
	})
}
