package scene

import (
	"github.com/bbfh-dev/ttk/ttk"
	"github.com/charmbracelet/lipgloss"
)

type TtkWindow struct {
	Widgets []ttk.Widget
	length  int // Cached len(Widgets)
}

func (window *TtkWindow) Init() {
	window.length = len(window.Widgets)

	for _, widget := range window.Widgets {
		widget.Init()
	}
}

func (window *TtkWindow) Scene() string {
	var blocks = make([]string, window.length)
	for i, widget := range window.Widgets {
		blocks[i] = widget.Render(0, 0)
	}

	return lipgloss.Place(
		ttk.Width, ttk.Height, 0.5, 0.5,
		lipgloss.JoinVertical(0, blocks...),
	)
}
