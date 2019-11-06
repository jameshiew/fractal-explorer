package gui

// hideable is something which can be hidden
//
// subset of the fyne.CanvasObject interface
type hideable struct {
	hidden bool
}

func (h *hideable) Visible() bool {
	return !h.hidden
}

func (h *hideable) Show() {
	h.hidden = false
}

func (h *hideable) Hide() {
	h.hidden = true
}
