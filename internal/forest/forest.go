package forest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Auth struct {
	// Forests user id
	// TODO implement
	Id string

	// Forests username
	// TODO implement
	Username string

	// Forests access_token
	// TODO implement
	AccessToken string

	// Forests bearer token, defined by external method
	BearerToken string
}

func Connect(jwt string) *Auth {
	if jwt == "" {
		log.Fatalf("No bearer token was given, was given %v", jwt)
	}

	return &Auth{
		Id:          "notImplemented",
		Username:    "notImplemented",
		AccessToken: "notImplemented",
		BearerToken: jwt,
	}
}

// get global builds
func (f *Auth) GetBuilds() forest_builds {
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

	var builds forest_builds

	// parse the json into something we can use
	err = json.Unmarshal(body, &builds)

	if err != nil {
		log.Println("error while parsing json")
	}

	return builds
}

// func (f *Auth) GetRefreshToken() forest_refresh_token {
// 	// set the url and method
// 	req, _ := http.NewRequest("GET", "https://api.forest.host/v1/auth/refresh_token/", nil)

// 	// add the auth header
// 	req.Header.Set("Authorization", f.BearerToken)

// 	// send the request
// 	resp, err := http.DefaultClient.Do(req)
// 	log.Println(resp)

// 	if err != nil {
// 		log.Printf("got error with http status %s while fetching builds", err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)

// 	if err != nil {
// 		log.Fatalln("error while reading body")
// 	}

// 	var builds forest_builds

// 	// parse the json into something we can use
// 	err = json.Unmarshal(body, &builds)

// 	if err != nil {
// 		log.Println("error while parsing json")
// 	}

// 	return builds
// }
