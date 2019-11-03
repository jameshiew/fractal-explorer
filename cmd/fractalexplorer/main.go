package main

import (
	"gitlab.com/jameshiew/fractal-explorer/internal/gui"
)

const title = "Fractal Explorer"

func main() {
	window := gui.Window(title)
	window.RequestFocus()
	window.ShowAndRun()
}
