package fractal

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"log"
)

func (f *fractalCanvas) Tapped(*fyne.PointEvent) {
	// TODO: implement zoom here
	log.Print("tapped")
	widget.Refresh(f)
}

func (f *fractalCanvas) TappedSecondary(*fyne.PointEvent) {
	// do nothing
}
