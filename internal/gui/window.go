package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
)

func setUpWindow(title string) fyne.Window {
	window := app.New().NewWindow(title)
	wdgt := newFractalWidget()
	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewMaxLayout(),
			&wdgt,
			wdgt.labels.info,
		),
	)
	window.Canvas().SetOnTypedKey(wdgt.controllerFunc())
	return window
}
