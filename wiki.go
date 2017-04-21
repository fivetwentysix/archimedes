package main

import (
	"fmt"

	"github.com/bndr/gopencils"
	"github.com/kevinburke/rest"
)

type pageStruct struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Space struct {
		Key string `json:"key"`
	} `json:"space"`
	Body struct {
		Storage struct {
			Value          string `json:"value"`
			Representation string `json:"representation"`
		} `json:"storage"`
	} `json:"body"`
	Version struct {
		Number int `json:"number"`
	} `json:"version"`
}

func main() {

	// Create Basic Auth
	auth := gopencils.BasicAuth{Username: "archimedes", Password: "LOOKINLASTPASS"}

	// Create New Api with our auth
	api := gopencils.Api("https://newcontext.atlassian.net/wiki/rest/api/content/78025190", &auth)

	rest.DefaultTransport.Debug = false
	api.Api.Client.Transport = rest.DefaultTransport

	// api.Client.Transport = rest.DefaultTransport

	querystring := map[string]string{"expand": "space,title,body.view,body.storage,version"}

	// Perform a GET request
	// URL Requested: http://your-api-url.com/api/users
	// foo, _ := api.Res("78025190", &respStruct{}).Get()

	fetchData := &pageStruct{}

	_, err := api.Res("", fetchData).Get(querystring)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Fetch Data: ", fetchData)

	fmt.Printf("Body: %v \n", fetchData) // resp.Body.Storage.Value)

	version := fetchData.Version.Number + 1
	body := fetchData.Body.Storage.Value + "<p>NEW STUFF 123</p>"

	// postData.Expandable.Version = version
	fetchData.Body.Storage.Value = body
	fetchData.Version.Number = version

	rest.DefaultTransport.Debug = true

	newresponse, err := api.Res("", fetchData).Put(fetchData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("New response: ", newresponse.Response)

}
