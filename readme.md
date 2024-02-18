# forest watcher
This is a repo for Forestwatch. A simple program that sends you a notification when your [forest](forest.host) build has failed or passed. This solves a common frustration of Forest not refreshing when you leave the browser window.


This program creates a window where you enter your Forest bearer token and a widget where you can see the build status.
The program also sends you a notification when the status of your build has changed

## development
The app is build using go and [fyne](https://fyne.io/) for the gui

### install
you need [dlv](https://github.com/derekparker/delve/tree/master/Documentation/installation)

### commands
`make build` builds a new version and outputs it to `./bin/main`
`make run` builds a binary and directly run it
`make lint` run the go linter
`make fmt` format the code
`make publish` create a packaged macos app
`make test` run the unit tests


### structure
the folder structure is based on [https://github.com/hashicorp/terraform/](https://github.com/hashicorp/terraform/)