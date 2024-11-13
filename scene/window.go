package scene

import "github.com/bbfh-dev/ttk/ttk"

type TtkWindow struct {
	Widgets []ttk.Widget
}

func (window *TtkWindow) Init() {
	for _, widget := range window.Widgets {
		widget.Init()
	}
}

func (window *TtkWindow) Scene() string {
	return window.Widgets[0].Render(0, 0)
}
