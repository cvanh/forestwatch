package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"github.com/cvanh/forestwatch/internal/forest"
	"github.com/cvanh/forestwatch/internal/gui"
)

const APPNAME = "forest watcher"

// store our forest token
func main() {
	a := &gui.App{App: app.New()}

	a.NewWindow = a.App.NewWindow(APPNAME)

	token := a.Preferences().String("ForestBearer")

	if token == "" {
		a.RenderWidget()
	}
	f := forest.Connect(token)

	// fetchBuilds(f)

	a.RenderWindow()

	// k.SendNotification()

	a.Render()

}

func fetchBuilds(f *forest.Auth) {
	builds := f.GetBuilds()

	log.Println(builds)
}
