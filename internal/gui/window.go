package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
)

func layoutWindow(window fyne.Window, wdgt fractalWidget) {
	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewMaxLayout(),
			&wdgt,
			wdgt.InfoLabel(),
		),
	)
	window.Canvas().SetOnTypedKey(wdgt.controllerFunc())
}
