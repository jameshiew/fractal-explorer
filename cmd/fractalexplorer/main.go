package main

import (
	"github.com/op/go-logging"

	"gitlab.com/jameshiew/fractal-explorer/internal/gui"
)

const title = "Fractal Explorer"

func main() {
	log := logging.MustGetLogger(title)

	log.Info("Starting up")
	gui.Run(title)
	log.Info("Shutting down")
}
