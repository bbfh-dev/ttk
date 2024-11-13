package main

import (
	"fmt"

	"github.com/bbfh-dev/ttk/display"
	"github.com/bbfh-dev/ttk/scene"
	"github.com/bbfh-dev/ttk/ttk"
)

func main() {
	scene := scene.TtkWindow{
		Widgets: []ttk.Widget{
			&display.TtkLabel{
				Label:    "Hello World!\nSomething else",
				Justify:  0,
				Width:    display.AUTO,
				MaxWidth: display.AUTO,
			},
		},
	}
	scene.Init()
	fmt.Println(scene.Scene())
}
