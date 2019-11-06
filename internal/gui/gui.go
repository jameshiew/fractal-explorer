package gui

import "fyne.io/fyne/app"

// Run launches the GUI and blocks until exit
func Run(title string) {
	window := app.New().NewWindow(title)
	wdgt := newFractalWidget()
	layoutWindow(window, wdgt)
	window.RequestFocus()
	window.ShowAndRun()
}
