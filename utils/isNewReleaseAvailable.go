package utils

import "github.com/baturtulek/apple-release-notifier/types"

// IsNewReleaseAvailable Returns true if New Releases are available
func IsNewReleaseAvailable(newReleasesNotExistsInOldReleases []types.Release) bool {
	return len(newReleasesNotExistsInOldReleases) > 0
}
