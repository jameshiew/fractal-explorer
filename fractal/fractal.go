package fractal

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// Run runs the application
func Run() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
