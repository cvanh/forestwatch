package gui

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/cvanh/forestwatch/internal/forest"
)

const APPNAME = "forest watcher"

// store our forest token

type App struct {
	fyne.App
	NewWindow fyne.Window
}

// renders the window/settings where forest jwt is entered
func (a *App) RenderWindow() {

	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "jwt", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Printf("Form submitted, jwt is: %s", entry.Text)

			// set the jwt token for the forest sdk
			a.Preferences().SetString("ForestBearer", entry.Text)

			// we dont need the window anymore right
			a.NewWindow.Hide()
			a.GetBuildsCron()

		},
	}

	a.NewWindow.SetContent(form)
}

// render the app widget in taskbar
func (a *App) RenderWidget() {
	if desk, ok := a.App.(desktop.App); ok {
		m := fyne.NewMenu(APPNAME,
			fyne.NewMenuItem("settings", func() {
				a.NewWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

}

// renders all the defined windows
func (a *App) Render() {
	a.NewWindow.SetCloseIntercept(func() {
		a.NewWindow.Hide()
	})
	a.NewWindow.ShowAndRun()
}

// send notification to X window
func (a *App) SendNotification(Title string, Content string) {
	notif := fyne.NewNotification(Title, Content)
	a.App.SendNotification(notif)
}

func (a *App) GetBuildsCron() {
	token := a.Preferences().String("ForestBearer")

	Forest := forest.Connect(token)
	builds := Forest.GetBuilds()

	log.Println(builds)
	time.Sleep(1000000)
	go a.GetBuildsCron()
}
