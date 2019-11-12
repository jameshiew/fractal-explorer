package cartesian

func Convert(pixelX, pixelY, width, height int) (x, y float64) {
	return float64(pixelX - width/2), float64(-pixelY + height/2)
}
