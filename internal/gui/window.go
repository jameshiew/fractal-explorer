package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func setUpWindow(title string) fyne.Window {
	window := app.New().NewWindow(title)
	cnvs := newFractalWidget()
	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewMaxLayout(),
			&cnvs,
			cnvs.labels.info,
		),
	)
	ctrlr := controllerFor(&cnvs.viewport)
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		ctrlr(event)
		widget.Refresh(&cnvs)
	})
	return window
}
