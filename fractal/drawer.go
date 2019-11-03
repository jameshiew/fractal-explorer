package fractal

import "image/color"

type drawer struct {
	scale    float64
	position struct {
		x, y float64
	}
	mandelbrot mandelbrot
}

func (d *drawer) pixelColor(pixelX, pixelY, width, height int) color.Color {
	x, y := d.toCartesian(pixelX, pixelY, width, height)
	c := complex(x*d.scale, y*d.scale)
	iter := d.mandelbrot.iterateWhileNotReachingBound(c)
	if iter == d.mandelbrot.maxIterations {
		return color.Black
	}
	scale := float64(iter) / float64(d.mandelbrot.maxIterations)
	return color.RGBA{
		R: uint8(scale * 255),
		G: uint8(scale * 100),
		B: uint8(scale * 100),
		A: 255,
	}
}

func (d *drawer) toCartesian(pixelX, pixelY, width, height int) (x, y float64) {
	return d.position.x + float64(pixelX-width/2), d.position.y + float64(-pixelY+height/2)
}
