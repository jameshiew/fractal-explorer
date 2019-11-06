package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
)

func layoutWindow(window fyne.Window, wdgt fractalWidget) {
	container := fyne.NewContainerWithLayout(layout.NewMaxLayout())
	container.AddObject(&wdgt)
	container.AddObject(wdgt.InfoLabel())
	window.SetContent(container)
	window.Canvas().SetOnTypedKey(wdgt.controllerFunc())
}
