package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/cvanh/forestwatch/internal/forest"
)

const APPNAME = "forest watcher"

func main() {
	// prompt for token
	renderWindow()

	builds := forest.GetBuilds()
	log.Println(builds)
}

// renders the window where forest jwt is entered
func renderWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow(APPNAME)

	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "jwt", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Printf("Form submitted, jwt is: %s", entry.Text)

			// set the jwt token for the forest sdk
			// *&forest.ForestJwt = entry.Text

			// we dont need the window anymore right
			myWindow.Close()
		},
	}

	// if desk, ok := myApp.(desktop.App); ok {
	// 	m := fyne.NewMenu("MyApp",
	// 		fyne.NewMenuItem("Show jwt menu", func() {
	// 			myWindow.Show()
	// 		}))
	// 	desk.SetSystemTrayMenu(m)
	// }

	// myWindow.SetContent(widget.NewLabel("Fyne System Tray"))
	// myWindow.SetCloseIntercept(func() {
	// 	myWindow.Hide()
	// })

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
