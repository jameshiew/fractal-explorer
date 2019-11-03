package fractal

import (
	"fmt"
	"math/cmplx"
)

type mandelbrot struct {
	maxIterations uint
	bound         float64
}

func (m mandelbrot) String() string {
	return fmt.Sprintf("maxIterations=%v bound=%v", m.maxIterations, m.bound)
}
func (m mandelbrot) iterateWhileNotReachingBound(c complex128) (iterations uint) {
	var n uint
	var z complex128
	for cmplx.Abs(z) <= m.bound && n < m.maxIterations {
		z = z*z + c
		n++
	}
	return n
}
