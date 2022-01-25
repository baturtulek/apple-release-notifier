package utils

import "github.com/baturtulek/apple-release-notifier/types"

func CompareOldReleaseDataWithNewReleaseData(oldReleaseArr []types.Release, newReleaseArr []types.Release) []types.Release {
	var releasesNotExistInOldReleases []types.Release

	for _, newRelease := range newReleaseArr {
		if !isNewReleaseExistsiInOldRelease(newRelease, oldReleaseArr) {
			releasesNotExistInOldReleases = append(releasesNotExistInOldReleases, newRelease)
		}
	}

	return releasesNotExistInOldReleases
}

func isNewReleaseExistsiInOldRelease(newRelease types.Release, oldReleaseArr []types.Release) bool {
	for _, oldRelease := range oldReleaseArr {
		if newRelease.Platform == oldRelease.Platform && newRelease.Code == oldRelease.Code {
			return true
		}
	}
	return false
}
