package main

import (
	"os"

	"github.com/elegos/instago/source/gui/ui"
	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	mainWindow := ui.BuildInstagoMainWindow()
	mainWindow.Show()

	app.Exec()
}
