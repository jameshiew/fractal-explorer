package main

import (
	"os"

	"github.com/op/go-logging"

	"gitlab.com/jameshiew/fractal-explorer/internal/gui"
)

const title = "Fractal Explorer"

func main() {
	logging.SetBackend(logging.NewLogBackend(os.Stdout, "", 0))
	logging.SetFormatter(logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`))
	log := logging.MustGetLogger(title)

	log.Info("Starting up")
	gui.Run(title)
	log.Info("Shutting down")
}
