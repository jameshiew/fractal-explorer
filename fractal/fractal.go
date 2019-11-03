package fractal

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

const title = "Fractal Explorer"

// Run runs the application
func Run() {
	app := app.New()

	w := app.NewWindow(title)
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
