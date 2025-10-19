package mandelbrot

import (
	"fmt"
)

const (
	defaultMaxIterations = 70
	defaultBound         = 2
)

type ImageBuilder struct {
	Image
}

func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{
		Image{
			maxIterations: defaultMaxIterations,
			bound:         defaultBound,
		},
	}
}

func (ib *ImageBuilder) SetMaxIterations(maxIterations uint16) *ImageBuilder {
	ib.maxIterations = maxIterations
	return ib
}

func (ib *ImageBuilder) SetBound(bound float64) *ImageBuilder {
	ib.bound = bound
	return ib
}

func (ib *ImageBuilder) Build() Image {
	return ib.Image
}

// Image defines a Image set
type Image struct {
	maxIterations uint16
	bound         float64
}

func (m Image) MaxIterations() uint16 {
	return m.maxIterations
}

func (m Image) Bound() float64 {
	return m.bound
}

func (m Image) String() string {
	return fmt.Sprintf("maxIterations=%v bound=%v", m.maxIterations, m.bound)
}

func (m Image) IterateWhileNotReachingBound(c complex128) (iterations uint16) {
	var (
		n uint16
		z complex128
	)
	boundSquared := m.bound * m.bound
	for real(z)*real(z)+imag(z)*imag(z) <= boundSquared && n < m.maxIterations {
		z = z*z + c
		n++
	}
	return n
}
