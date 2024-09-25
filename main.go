package main

import (
	"gioui.org/unit"
	"log"
	"os"

	"gioui.org/app"

	mainAPP "github.com/Mr-Ao-Dragon/tinyRconConsole/ui/app"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("RconConsole"), app.Size(unit.Dp(1200), unit.Dp(800)))

		mainUI, err := mainAPP.NewUi(window)
		if err != nil {
			log.Fatal(err)
		}
		if err := mainUI.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
