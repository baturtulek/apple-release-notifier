package utils

import "github.com/baturtulek/apple-release-notifier/types"

func IsNewReleaseAvailable(newReleasesNotExistsInOldReleases []types.Release) bool {
	return len(newReleasesNotExistsInOldReleases) > 0
}
