package services

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/baturtulek/apple-release-notifier/constants"
	"github.com/baturtulek/apple-release-notifier/types"
)

// ReadLastReleaseFromFile Read Last Crawled Release Data From the File
func ReadLastReleaseFromFile() []types.Release {
	var lastReleases []types.Release

	if isLastReleaseFileExists() {
		jsonFile, err := os.Open(constants.LAST_CRAWL_RELEASES_FILE)
		if err != nil {
			log.Fatal("ERROR - ReadLastReleaseFromFile - File Open: ", err)
		}
		defer jsonFile.Close()

		jsonData, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatal("ERROR - ReadLastReleaseFromFile - JSON", err)
		}

		if err := json.Unmarshal(jsonData, &lastReleases); err != nil {
			log.Fatal("ERROR - ReadLastReleaseFromFile - UnMarshal: ", err)
		}
	}
	return lastReleases
}

// WriteNewReleaseDataToFile Write New Release To the File
func WriteNewReleaseDataToFile(lines []types.Release) {
	file, _ := json.MarshalIndent(lines, "", "  ")
	_ = ioutil.WriteFile(constants.LAST_CRAWL_RELEASES_FILE, file, 0644)
	log.Print("Last Release file updated.")
}

// ReadMailContactsFromFile Read Mail Addresses of the Clients
func ReadMailContactsFromFile() []string {
	var mailContacts []string

	file, err := os.Open(constants.MAIL_CONTACTS_FILE)
	if err != nil {
		log.Fatal("ERROR - ReadMailContactsFromFile - File Open: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mailContacts = append(mailContacts, strings.Trim(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("ERROR - ReadMailContactsFromFile - Scan File: ", err)
	}

	return mailContacts
}

func isLastReleaseFileExists() bool {
	_, err := os.Stat(constants.LAST_CRAWL_RELEASES_FILE)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
