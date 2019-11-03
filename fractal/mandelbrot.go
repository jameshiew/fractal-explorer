package fractal

import (
	"math/cmplx"
)

const (
	maxIterations = 2
	bound         = 2
)

func mandelbrot(c complex128) uint {
	var n uint
	var z complex128
	for cmplx.Abs(z) <= bound && n < maxIterations {
		z = z*z + c
		n++
	}
	return n
}
