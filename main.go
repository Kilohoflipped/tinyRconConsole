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
		windowSize := app.Size(unit.Dp(1200), unit.Dp(800))
		window.Option(app.Title("RconConsole"), windowSize)
		// window.Option(app.Decorated(false))

		mainUI, err := mainAPP.NewUI(window)
		if err != nil {
			log.Fatal(err)
		}
		// Run主窗体
		err = mainUI.Run()
		if err != nil {
			log.Fatal(err)
		}
		// 安全退出
		os.Exit(0)
	}()
	app.Main()
}
