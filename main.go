package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/cvanh/forestwatch/internal/forest"
	"github.com/cvanh/forestwatch/internal/gui"
)

const APPNAME = "forest watcher"

var Forest *forest.Auth

// store our forest token
func main() {
	// setup the window
	a := &gui.App{App: app.NewWithID("dev.imaretarded.forestwatcher")}
	a.NewWindow = a.App.NewWindow(APPNAME)

	// get token from storage
	token := a.Preferences().String("ForestBearer")

	// if there is no token given yet ask for one otherwise run get builds every x
	if token != "" {
		a.RenderWindow()
	} else {
		a.GetBuildsCron()
	}

	a.RenderWidget()

	a.Render()
}
