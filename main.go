package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const APPNAME = "forest watcher"

var ForestJwtToken string

type forest_user struct {
	id         string
	username   string
	created_at string
}

type forest_repository struct {
	id               string
	path             string
	builder_path_key string
	environment      forest_environment
}

type forest_environment struct {
	id       string
	branch   string
	project  forest_project
	services []forest_services
}

type forest_services struct {
	id          string
	name        string
	tools       []string
	deploy_path string
	index_path  string
}
type forest_team struct {
	id   string
	name string
	slug string
}

type forest_project struct {
	id   string
	name string
	slug string
	team forest_team
}

type forest_status struct {
	id          int
	name        string
	description string
}

type forest_build struct {
	id           string
	repository   forest_repository
	user         forest_user
	build_status forest_status
	deploy_path  string
	index_file   string
	before       string
	after        string
	deployed_at  string
	started_at   string
	finished_at  string
	message      string
}

type forest_builds_response struct {
	Builds []forest_build `json:"data"`
}

func main() {
	// prompt for token
	// renderWindow()
	getBuilds()

}

func getBuilds() {
	// set the url and method
	req, _ := http.NewRequest("GET", "https://api.forest.host/v1/builds/", nil)

	// add the auth header
	req.Header.Set("Authorization", "Bearer .Ksk5a-R86y7l3zmidTvVzARDVCtyE-eqP1owxARS1OFx1IT6O5rZ-5i81BBTYxk5VkW1_IMHd6FH2pg8MdAPB7cEVHiIjH4fDe-HW4cV8vWp9UIE7dv9QrcQ3Q-Zw0is0XZzDfHuE5mmtzOO_hsfcoF_eoDpJ1qgLm7NXiTqE5kKn4FuYng-fMlSoR_CeO5XhFdWMpPU7zg5lRVkCEYjXo0i07yhv0etrMs3dzf6CwTsv4cto_VwDId4eG_oS6hSZmTdEwfoMHux9KDQ7SjYzpoAGd4LCnRmBQ2zZJ9pzVUAgDPWa_hDsL6n4NFtqbtdWz6PuG_xz47-fLwxTbPdVDyBb6reiGqQH3R6mitakOE-oNWHdrcicK9oOclKVJFUORTXuqdP2ydTYCHFZJ_PjPxZe0cXlSfBleBI4zwtqBZPKPMK7EpYzVAyJsHZ902jXF7m_cxNNmGHLCEOsjNSzONnQJ_B7d6EySh276wYrKLIJX18CyNmVRIvQ73bhtxEufmdFcjQnefye-eTYLg9XG6P6s1K4uQqlszWsrfYtlvujDgB6ndXierc9ZhK13BbOd-vXeXpagwp6EyWBeZAsQ-tmaB_kj4qD4nKs1jDidQShKj5ljqxsKH7zKgKJRaxfoavLiuxvGSwq71mteVUxElLCua65jnQs2y0zb7kps8")
	// req.Header.Set("Authorization", ForestJwtToken)

	// send the request
	resp, err := http.DefaultClient.Do(req)
	log.Println(resp)

	if err != nil {
		log.Printf("got error with http status %s while fetching builds", err)
	}

	// defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("error while reading body")
	}

	var kaas forest_builds_response

	json.Unmarshal(body, &kaas)

	log.Printf("%v", kaas)
	// log.Printf("%v", kaas.builds)
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
