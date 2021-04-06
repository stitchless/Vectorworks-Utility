package software

import (
	"context"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v33/github"
	"log"
)

func SearchForUpdate(currentVersion string) {
	client := github.NewClient(nil)

	release, _, err := client.Repositories.GetLatestRelease(context.Background(), "jpeizer", "Vectorworks-Utility")

	if release != nil {
		gitLatestReleaseConstraint, _ := semver.NewConstraint("> " + *release.TagName)
		currentSemVersion, _ := semver.NewVersion(currentVersion)
		updateAvailable := gitLatestReleaseConstraint.Check(currentSemVersion)
		if updateAvailable == true {
			// Do Something
			fmt.Println("Updates Found")
			// Prompt for update
		} else {
			fmt.Println("No Update Found")
			// No Update Icon
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
