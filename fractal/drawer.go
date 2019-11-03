package fractal

import "image/color"

type drawer struct {
	scale struct {
		x, y float64
	}
	mandelbrot mandelbrot
}

func (d *drawer) pixelColor(pixelX, pixelY, width, height int) color.Color {
	x, y := toCartesian(pixelX, pixelY, width, height)
	c := complex(float64(x)*d.scale.x, float64(y)*d.scale.y)
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

func toCartesian(pixelX, pixelY, width, height int) (x, y int) {
	return pixelX - width/2, -pixelY + height/2
}
