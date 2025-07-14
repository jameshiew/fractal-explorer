package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func layoutWindow(window fyne.Window, wdgt fractalWidget) {
	ctnr := container.New(layout.NewMaxLayout())
	ctnr.Add(&wdgt)
	ctnr.Add(wdgt.InfoLabel())
	window.SetContent(ctnr)
	window.Canvas().SetOnTypedKey(wdgt.controllerFunc())
}
