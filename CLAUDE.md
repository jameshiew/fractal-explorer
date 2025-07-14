# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based desktop GUI application for exploring the Mandelbrot set fractal. It uses the Fyne framework for cross-platform GUI rendering.

## Development Commands

```bash
# Build the main application
just app

# Run all tests
just test

# Run tests with race detector (used in CI)
just test-ci

# Build all packages
just build

# Upgrade all dependencies
just upgrade-deps

# Format Go code
go fmt ./...
```

## Architecture

The codebase follows Go's internal package pattern with clear separation of concerns:

### Core Components

1. **Entry Point** (`cmd/fractalexplorer/main.go`): Initializes logging and launches the GUI

2. **GUI Layer** (`internal/gui/`): Fyne-based user interface
   - `widget.go`: Main fractal widget that handles rendering and user interaction
   - `controls.go`: Keyboard input handling (arrow keys for pan, W/S for zoom)
   - `taps.go`: Mouse/touch input handling
   - `window.go`: Window layout and info display

3. **Fractal Mathematics** (`internal/mandelbrot/`): Core Mandelbrot set calculations
   - Uses builder pattern for configuration
   - Implements the iterative algorithm: z = z² + c

4. **Coordinate System** (`internal/cartesian/`): Pixel to mathematical coordinate conversion
   - `viewport.go`: Manages view position, scale, and navigation
   - Handles zoom and pan transformations

5. **Rendering Engine** (`internal/draw/`): Parallel image generation
   - Multi-threaded rendering with 1024 worker goroutines
   - `colorizer.go`: Maps iteration counts to colors
   - Performance tracking (pixels/second)

### Key Design Patterns

- **Parallel Processing**: The renderer uses a worker pool pattern to compute fractal pixels in parallel
- **Builder Pattern**: Mandelbrot set configuration uses a builder for flexible setup
- **Interface-based Design**: Widget renderer follows Fyne's interface patterns
- **Separation of Concerns**: Math, rendering, and UI logic are cleanly separated

## Testing Approach

- Unit tests use `stretchr/testify` for assertions
- Test files follow Go convention: `*_test.go` alongside implementation files
- Run individual test files with: `go test ./internal/cartesian/cartesian_test.go`

## Important Implementation Details

- The application uses centered coordinates where (0,0) is at the center of the screen
- Default viewport scale is 0.01 (100 pixels = 1 unit in fractal space)
- Mandelbrot calculations use 70 iterations by default with a bound of 2
- The renderer creates exactly 1024 worker goroutines for parallel processing
- Logging uses the `go-logging` package with colored console output