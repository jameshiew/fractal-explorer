package mandelbrot

import (
	"fmt"
	"math/cmplx"
)

// Mandelbrot defines a Mandelbrot set
type Mandelbrot struct {
	maxIterations uint16
	bound         float64
}

func New(maxIterations uint16, bound float64) Mandelbrot {
	return Mandelbrot{
		maxIterations: maxIterations,
		bound:         bound,
	}
}

func (m Mandelbrot) MaxIterations() uint16 {
	return m.maxIterations
}

func (m Mandelbrot) Bound() float64 {
	return m.bound
}

func (m Mandelbrot) String() string {
	return fmt.Sprintf("maxIterations=%v bound=%v", m.maxIterations, m.bound)
}

func (m Mandelbrot) IterateWhileNotReachingBound(c complex128) (iterations uint16) {
	var (
		n uint16
		z complex128
	)
	for cmplx.Abs(z) <= m.bound && n < m.maxIterations {
		z = z*z + c
		n++
	}
	return n
}
