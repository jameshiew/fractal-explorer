package fractal

import "image/color"

type drawer struct {
	mandelbrot mandelbrot
}

func (d *drawer) pixelColor(pixelX, pixelY, width, height int) color.Color {
	x, y := toCartesian(pixelX, pixelY, width, height)
	c := complex(float64(x)/100, float64(y)/100)
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
