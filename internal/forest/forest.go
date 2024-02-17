package forest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type forest_user struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Created  string `json:"created_at"`
}

type forest_repository struct {
	Id            string             `json:"id"`
	Path          string             `json:"Path"`
	BuilderPubKey string             `json:"builder_path_key"`
	Environment   forest_environment `json:"environment"`
}

type forest_environment struct {
	Id       string            `json:"id"`
	Branch   string            `json:"branch"`
	Project  forest_project    `json:"project"`
	Services []forest_services `json:"services"`
}

type forest_services struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Tools      []string `json:"tools"`
	DeployPath string   `json:"deploy_path"`
	IndexPath  string   `json:"index_path"`
}
type forest_team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type forest_project struct {
	Id   string      `json:"id"`
	Name string      `json:"name"`
	Slug string      `json:"slug"`
	Team forest_team `json:"team"`
}

type forest_status struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"description"`
}

type forest_build struct {
	Id          string            `json:"id"`
	Repository  forest_repository `json:"repository"`
	User        forest_user       `json:"user"`
	Status      forest_status     `json:"build_status"`
	DeployPath  string            `json:"deploy_path"`
	IndexFile   string            `json:"index_file"`
	Before      string            `json:"before"`
	After       string            `json:"after"`
	Deployed    string            `json:"deployed_at"`
	Started     string            `json:"started_at"`
	Finished_at string            `json:"finished_at"`
	Message     string            `json:"message"`
}

var token string

type forest_builds_response struct {
	Builds []forest_build `json:"builds"`
}

func GetBuilds(token string) {
	// set the url and method
	req, _ := http.NewRequest("GET", "https://api.forest.host/v1/builds/", nil)

	// add the auth header
	req.Header.Set("Authorization", token)

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

	var builds forest_builds_response

	json.Unmarshal(body, &builds)
}
