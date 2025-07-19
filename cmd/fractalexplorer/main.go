package main

import (
	"log/slog"
	"os"

	"github.com/jameshiew/fractal-explorer/internal/gui"
)

const title = "Fractal Explorer"

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)

	logger.Info("Starting up")
	gui.Run(logger, title)
	logger.Info("Shutting down")
}
