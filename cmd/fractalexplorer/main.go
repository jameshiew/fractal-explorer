package main

import (
	"log"

	"gitlab.com/jameshiew/fractal-explorer/internal/gui"
)

const title = "Fractal Explorer"

func main() {
	log.Println("Starting up")
	gui.Run(title)
	log.Println("Shutting down")
}
