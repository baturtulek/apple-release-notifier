package main

import (
	"log"

	"github.com/baturtulek/apple-release-notifier/constants"
	"github.com/baturtulek/apple-release-notifier/services"
	"github.com/baturtulek/apple-release-notifier/utils"
)

func main() {
	utils.LoadEnvironmentVariables()
	lastReleaseData := services.ReadLastReleaseFromFile()
	response := services.GetPageContentFromURL(constants.APPLE_RELEASE_PAGE_URL)
	newReleaseData := utils.ParsePageContent(response)
	newReleasesNotExistsInOldReleases := utils.CompareOldReleaseDataWithNewReleaseData(lastReleaseData, newReleaseData)
	if utils.IsNewReleaseAvailable(newReleasesNotExistsInOldReleases) {
		log.Print("New Releases Available.")
		services.SendMail(newReleasesNotExistsInOldReleases)
		services.WriteNewReleaseDataToFile(newReleaseData)
	} else {
		log.Print("No new release Available")
	}
}
