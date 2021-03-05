package software

import (
	"context"
	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v33/github"
	"log"
)

func SearchForUpdate(currentVersion string) {
	client := github.NewClient(nil)

	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "jpeizer", "vectorworks-app-cleaner")

	if release != nil {
		gitLatestReleaseConstraint, _ := semver.NewConstraint("> " + *release.TagName)
		currentSemVersion, _ := semver.NewVersion(currentVersion)
		updateAvailable := gitLatestReleaseConstraint.Check(currentSemVersion)
		if updateAvailable == true {
			// Do Something
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
