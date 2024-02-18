package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

const APPNAME = "forest watcher"

// store our forest token

type App struct {
	fyne.App
	NewWindow fyne.Window
}

// var a *App

// renders the window/settings where forest jwt is entered
func (a *App) RenderWindow() {
	// a.App = app.New()

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
		},
	}

	a.NewWindow.SetContent(form)
}

// render the widget in taskbar
func (a *App) RenderWidget() {
	if desk, ok := a.App.(desktop.App); ok {
		m := fyne.NewMenu(APPNAME,
			fyne.NewMenuItem("settings", func() {
				a.NewWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	a.NewWindow.SetContent(widget.NewLabel("Fyne System Tray"))
	a.NewWindow.SetCloseIntercept(func() {
		a.NewWindow.Hide()
	})
}

// renders all the defined windows
func (a *App) Render() {
	a.NewWindow.ShowAndRun()
}
