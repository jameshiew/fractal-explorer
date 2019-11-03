package fractal

import "image/color"

type drawer struct {
	scale struct {
		x, y float64
	}
	position struct {
		x, y float64
	}
	mandelbrot mandelbrot
}

func (d *drawer) pixelColor(pixelX, pixelY, width, height int) color.Color {
	x, y := d.toCartesian(pixelX, pixelY, width, height)
	c := complex(x*d.scale.x, y*d.scale.y)
	if d.mandelbrot.iterateWhileNotReachingBound(c) == d.mandelbrot.maxIterations {
		return color.Black
	}
	return color.RGBA{
		R: 255,
		G: 100,
		B: 100,
		A: 255,
	}
}

func (d *drawer) toCartesian(pixelX, pixelY, width, height int) (x, y float64) {
	return d.position.x + float64(pixelX-width/2), d.position.y + float64(-pixelY+height/2)
}
