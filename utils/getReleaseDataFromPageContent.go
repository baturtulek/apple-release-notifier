package utils

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/baturtulek/apple-release-notifier/types"
)

func getPlatformDataFromString(str string) string {
	value := strings.IndexAny(str, "0123456789")
	if value >= 0 && value <= len(str) {
		return str[:value]
	}
	return str
}

func getVersionDataFromString(str string) string {
	value := strings.IndexAny(str, "(")
	if value >= 0 && value <= len(str) {
		return str[:value]
	}
	return str
}

func parseAndCreateReleaseObject(release string) types.Release {
	platform := strings.Trim(getPlatformDataFromString(release), " ")
	release = strings.ReplaceAll(release, platform, "")
	version := strings.Trim(getVersionDataFromString(release), " ")
	release = strings.ReplaceAll(release, version, "")
	code := strings.ReplaceAll(release, "(", "")
	code = strings.ReplaceAll(code, ")", "")
	code = strings.Trim(code, " ")
	releaseObj := types.Release{
		Platform: platform,
		Version:  version,
		Code:     code,
	}
	return releaseObj
}

func ParsePageContent(response *http.Response) []types.Release {
	var releaseArr []types.Release

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".article-title").Each(func(i int, s *goquery.Selection) {
		release := s.Find("h2").Text()
		releaseObj := parseAndCreateReleaseObject(release)
		releaseArr = append(releaseArr, releaseObj)
	})

	return releaseArr
}
