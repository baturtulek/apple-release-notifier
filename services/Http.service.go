package services

import (
	"log"
	"net/http"
)

func GetPageContentFromURL(url string) *http.Response {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal("GET PAGE CONTENT ERROR:", err)
	}

	if response.StatusCode != 200 {
		log.Fatal("STATUS CODE ERROR: ", response.StatusCode)
	}

	return response
}
