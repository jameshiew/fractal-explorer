package gui

// Run launches the GUI and blocks until exit
func Run(title string) {
	window := setUpWindow(title)
	window.RequestFocus()
	window.ShowAndRun()
}
