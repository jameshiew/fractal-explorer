package fractal

import (
	"fyne.io/fyne/app"
)

const title = "Fractal Explorer"

// Run runs the application
func Run() {
	app := app.New()

	window := app.NewWindow(title)
	setUpWindow(window)
	window.RequestFocus()
	window.ShowAndRun()
}
