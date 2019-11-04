package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
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
	ctrlr := controllerFor(&wdgt.viewport)
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		ctrlr(event)
		widget.Refresh(&wdgt)
	})
	return window
}
