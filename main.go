package main

import (
	"fmt"
	"os"

	"github.com/bbfh-dev/ttk/display"
	"github.com/bbfh-dev/ttk/scene"
	"github.com/bbfh-dev/ttk/ttk"
	"golang.org/x/crypto/ssh/terminal"
)

const FIGLET = ` _____ _____ _  __
|_   _|_   _| |/ /
  | |   | | | ' <
  |_|   |_| |_|\_\`

func main() {
	scene := scene.TtkWindow{
		Widgets: []ttk.Widget{
			&display.TtkArt{Content: FIGLET},
			&display.TtkLabel{
				Label:    "Hello World! And other species! What a wonderful long text about Fooga that is so based and fat and based, did I mention based? Anyway. Fooga. Yes, Fooga is incredible. It's so round and... fooga. Really, no other explanation.",
				Justify:  0,
				Width:    display.AUTO,
				MaxWidth: display.AUTO,
			},
		},
	}
	scene.Init()

	var err error
	ttk.Width, ttk.Height, err = terminal.GetSize(int(os.Stdin.Fd()))
	ttk.Height -= 4 // My hardcoded prompt height™ + command
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(scene.Scene())
}
