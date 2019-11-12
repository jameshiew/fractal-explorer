package draw

import (
	"image/color"
)

// Colorizers color complex numbers
type Colorizer func(complex128) color.Color

func Blend(colors ...color.Color) color.Color {
	var r, g, b, a uint32
	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r += cr
		g += cg
		b += cb
		a += ca
	}
	return color.RGBA64{
		R: uint16(r / uint32(len(colors))),
		G: uint16(g / uint32(len(colors))),
		B: uint16(b / uint32(len(colors))),
		A: uint16(a / uint32(len(colors))),
	}
}

type Colorizable interface {
	MaxIterations() uint16
	IterateWhileNotReachingBound(complex128) uint16
}

func NewColorizer(base color.Color, colorizable Colorizable) Colorizer {
	return func(c complex128) color.Color {
		iter := colorizable.IterateWhileNotReachingBound(c)
		if iter == colorizable.MaxIterations() {
			return color.Black
		}
		scale := float64(iter) / float64(colorizable.MaxIterations())
		r, g, b, a := base.RGBA()
		return color.RGBA64{
			R: uint16(scale * float64(r)),
			G: uint16(scale * float64(g)),
			B: uint16(scale * float64(b)),
			A: uint16(a),
		}
	}
}
