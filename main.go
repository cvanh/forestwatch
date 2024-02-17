package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const APPNAME = "forest watcher"

var ForestJwtToken string

func main() {
	// prompt for token
	renderWindow()

}

// renders the window where forest jwt is entered
func renderWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow(APPNAME)

	entry := widget.NewEntry()
	// textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Entry", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted, jwt is:", entry.Text)

			// set the jwt token for the application
			ForestJwtToken = entry.Text

			// we dont need the window anymore right
			myWindow.Close()
		},
	}

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
