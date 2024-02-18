package forest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// type Api interface {
// 	GetBuilds()
// 	GetRefreshToken()
// }

type Auth struct {
	// Forests user id
	Id string

	// Forests username
	Username string

	// Forests access_token
	AccessToken string

	// Forests bearer token, defined by external method
	BearerToken string
}

// get global builds
func (f *Auth) GetBuilds() forest_builds_response {
	// set the url and method
	req, _ := http.NewRequest("GET", "https://api.forest.host/v1/builds/", nil)

	// add the auth header
	log.Println("using", f.BearerToken)
	req.Header.Set("Authorization", f.BearerToken)

	// send the request
	resp, err := http.DefaultClient.Do(req)
	log.Println(resp)

	if err != nil {
		log.Printf("got error with http status %s while fetching builds", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("error while reading body")
	}

	var builds forest_builds_response

	// parse the json into something we can use
	err = json.Unmarshal(body, &builds)

	if err != nil {
		log.Println("error while parsing json")
	}

	return builds
}

func (f *Auth) GetRefreshToken() forest_builds_response {
	// set the url and method
	req, _ := http.NewRequest("GET", "https://api.forest.host/v1/auth/refresh_token/", nil)

	// add the auth header
	req.Header.Set("Authorization", f.BearerToken)

	// send the request
	resp, err := http.DefaultClient.Do(req)
	log.Println(resp)

	if err != nil {
		log.Printf("got error with http status %s while fetching builds", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("error while reading body")
	}

	var builds forest_builds_response

	// parse the json into something we can use
	err = json.Unmarshal(body, &builds)

	if err != nil {
		log.Println("error while parsing json")
	}

	return builds
}
