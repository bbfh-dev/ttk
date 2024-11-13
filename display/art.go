package display

type TtkArt struct {
	Content string
}

func (element *TtkArt) Init() {
}

func (widget *TtkArt) Render(width int, height int) string {
	return widget.Content
}

func (widget *TtkArt) ExpandWidth() bool {
	return false
}

func (widget *TtkArt) ExpandHeight() bool {
	return false
}
