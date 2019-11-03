package main

import (
	"fractal-explorer/internal/fractal"
)

const title = "Fractal Explorer"

func main() {
	window := fractal.Window(title)
	window.RequestFocus()
	window.ShowAndRun()
}
