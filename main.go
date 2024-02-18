package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"github.com/cvanh/forestwatch/internal/forest"
	"github.com/cvanh/forestwatch/internal/gui"
)

const APPNAME = "forest watcher"

// store our forest token
var ForestBearer string

func main() {
	k := &gui.App{App: app.New()}

	k.NewWindow = k.App.NewWindow(APPNAME)

	k.RenderWidget()
	k.RenderWindow()

	// k.SendNotification()

	k.Render()

	f := forest.Connect(k.Preferences().String("ForestBearer"))
	builds := f.GetBuilds()

	log.Println(builds)

}
