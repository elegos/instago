package ui

import (
	"github.com/therecipe/qt/widgets"
)

// InstagoMainWindow the main window
type InstagoMainWindow struct {
	widgets.QMainWindow
}

// BuildInstagoMainWindow generate a new main window
func BuildInstagoMainWindow() InstagoMainWindow {
	mw := InstagoMainWindow{}

	return mw
}

// Binc nothing
func Binc() {

}
