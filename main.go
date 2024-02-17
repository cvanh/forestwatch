package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/cvanh/forestwatch/internal/forest"
)

const APPNAME = "forest watcher"

var ForestJwtToken string

func main() {
	// prompt for token
	renderWindow()
	forest.GetBuilds(ForestJwtToken)
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

			// set the jwt token for the application
			ForestJwtToken = entry.Text

			// we dont need the window anymore right
			myWindow.Close()
		},
	}

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
