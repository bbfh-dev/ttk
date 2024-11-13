package display

import (
	"github.com/charmbracelet/lipgloss"
)

const PREFERED_TEXT_WIDTH = 60

const AUTO = -1

var defaultStyle = lipgloss.NewStyle()

type TtkLabel struct {
	Label    string
	Justify  lipgloss.Position
	Width    int
	MaxWidth int

	style lipgloss.Style
}

func (element *TtkLabel) Init() {
	element.style = lipgloss.NewStyle().AlignHorizontal(element.Justify)
}

func (widget *TtkLabel) Render(width int, height int) string {
	label := widget.Label
	labelWidth := lipgloss.Width(label)

	if widget.MaxWidth != AUTO && labelWidth > widget.MaxWidth {
		label = defaultStyle.MaxWidth(widget.MaxWidth).Render(label)
	}

	if width == 0 {
		if widget.Width == AUTO {
			contentWidth := lipgloss.Width(label)
			if contentWidth < PREFERED_TEXT_WIDTH {
				return widget.style.Render(label)
			}

			return widget.style.Width(PREFERED_TEXT_WIDTH).Render(label)
		}

		return widget.style.Width(widget.Width).Render(label)
	}

	return widget.style.Width(width).Height(height).Render(label)
}

func (widget *TtkLabel) ExpandWidth() bool {
	return widget.Width == AUTO
}

func (widget *TtkLabel) ExpandHeight() bool {
	return false
}
