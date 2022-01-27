package services

import (
	"log"
	"net/http"
)

// Sends HTTP Request to Web Page and Returns the Response
func GetPageContentFromURL(url string) *http.Response {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal("ERROR - GetPageContentFromURL: ", err)
	}

	if response.StatusCode != 200 {
		log.Fatal("ERROR - GetPageContentFromURL - Status Code: ", response.StatusCode)
	}

	return response
}
