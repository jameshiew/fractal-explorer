package fractal

import "fyne.io/fyne"

const (
	minWidth  = 320
	minHeight = 240
)

type viewport struct {
	canvas fyne.CanvasObject
}

func (v *viewport) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	v.canvas.Resize(size)
}

func (v *viewport) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(minWidth, minHeight)
}
