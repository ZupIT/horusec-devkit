package mageutils

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-github/v40/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
)

func newGithubMock(repositoryRelease *github.RepositoryRelease, repositoryTags []*github.RepositoryTag) *http.Client {
	return mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetReposReleasesLatestByOwnerByRepo,
			repositoryRelease,
		),
		mock.WithRequestMatch(
			mock.GetReposTagsByOwnerByRepo,
			repositoryTags,
		),
	)
}

type testCaseVersion struct {
	releaseType               string
	repositoryRelease         *github.RepositoryRelease
	repositoryTags            []*github.RepositoryTag
	expectedReleaseVersion    string
	expectedReleaseBranchName string
	expectedBetaVersion       string
	expectedRCVersion         string
}

func toStringPointer(value string) *string {
	return &value
}

func newMockRelease(tagName string) *github.RepositoryRelease {
	return &github.RepositoryRelease{TagName: toStringPointer(tagName)}
}

func newMockTags(tags []string) []*github.RepositoryTag {
	var repositoryTags []*github.RepositoryTag

	for _, tag := range tags {
		repositoryTags = append(repositoryTags, &github.RepositoryTag{Name: toStringPointer(tag)})
	}

	return repositoryTags
}

func TestVersion(t *testing.T) {
	t.Run("should updated versions with according the expected", func(t *testing.T) {
		testCases := []testCaseVersion{
			{
				releaseType:               "patch",
				repositoryRelease:         newMockRelease("v1.0.0"),
				repositoryTags:            newMockTags([]string{"v1.0.0", "v0.1.0", "v0.0.1"}),
				expectedReleaseVersion:    "v1.0.1",
				expectedReleaseBranchName: "release/v1.0",
				expectedBetaVersion:       "v1.0.1-beta.1",
				expectedRCVersion:         "v1.0.1-rc.1",
			},
			{
				releaseType:               "minor",
				repositoryRelease:         newMockRelease("v1.0.0"),
				repositoryTags:            newMockTags([]string{"v1.0.0", "v0.1.0", "v0.0.1"}),
				expectedReleaseVersion:    "v1.1.0",
				expectedReleaseBranchName: "release/v1.1",
				expectedBetaVersion:       "v1.1.0-beta.1",
				expectedRCVersion:         "v1.1.0-rc.1",
			},
			{
				releaseType:               "major",
				repositoryRelease:         newMockRelease("v1.0.0"),
				repositoryTags:            newMockTags([]string{"v1.0.0", "v0.1.0", "v0.0.1"}),
				expectedReleaseVersion:    "v2.0.0",
				expectedReleaseBranchName: "release/v2.0",
				expectedBetaVersion:       "v2.0.0-beta.1",
				expectedRCVersion:         "v2.0.0-rc.1",
			},
			{
				releaseType:               "major",
				repositoryRelease:         newMockRelease("v1.0.0"),
				repositoryTags:            newMockTags([]string{"v1.0.0", "v0.1.0-rc.1", "v0.1.0-beta.1"}),
				expectedReleaseVersion:    "v2.0.0",
				expectedReleaseBranchName: "release/v2.0",
				expectedBetaVersion:       "v2.0.0-beta.1",
				expectedRCVersion:         "v2.0.0-rc.1",
			},
			{
				releaseType:               "minor",
				repositoryRelease:         newMockRelease("v2.0.0"),
				repositoryTags:            newMockTags([]string{"v2.1.0-beta.1", "v2.1.0-rc.1", "v2.0.0", "v2.0.0-beta.1", "v2.0.0-rc.1"}),
				expectedReleaseVersion:    "v2.1.0",
				expectedReleaseBranchName: "release/v2.1",
				expectedBetaVersion:       "v2.1.0-beta.2",
				expectedRCVersion:         "v2.1.0-rc.2",
			},
			{
				releaseType:               "minor",
				repositoryRelease:         newMockRelease("v2.5.0"),
				repositoryTags:            newMockTags([]string{"v2.6.0-beta.1", "v2.5.0-rc.1", "v2.5.0", "v2.4.0"}),
				expectedReleaseVersion:    "v2.6.0",
				expectedReleaseBranchName: "release/v2.6",
				expectedBetaVersion:       "v2.6.0-beta.2",
				expectedRCVersion:         "v2.6.0-rc.1",
			},
			{
				releaseType:               "minor",
				repositoryRelease:         newMockRelease("v3.3.3"),
				repositoryTags:            newMockTags([]string{"v3.3.0-rc.1", "v3.3.0-beta.1", "v3.2.2", "v3.2.1", "v3.2.0", "v3.2.0-rc.2", "v3.2.0-rc.1", "v3.2.0-beta.3", "v3.2.0-beta.2", "v3.2.0-beta.1", "v3.1.1"}),
				expectedReleaseVersion:    "v3.4.0",
				expectedReleaseBranchName: "release/v3.4",
				expectedBetaVersion:       "v3.4.0-beta.1",
				expectedRCVersion:         "v3.4.0-rc.1",
			},
			{
				releaseType:               "patch",
				repositoryRelease:         newMockRelease("v3.2.2"),
				repositoryTags:            newMockTags([]string{"v3.2.2", "v3.2.1", "v3.2.0", "v3.2.0-rc.2", "v3.2.0-rc.1", "v3.2.0-beta.3", "v3.2.0-beta.2", "v3.2.0-beta.1", "v3.1.1"}),
				expectedReleaseVersion:    "v3.2.3",
				expectedReleaseBranchName: "release/v3.2",
				expectedBetaVersion:       "v3.2.3-beta.1",
				expectedRCVersion:         "v3.2.3-rc.1",
			},
			{
				releaseType:               "major",
				repositoryRelease:         newMockRelease("v5.5.5"),
				repositoryTags:            newMockTags([]string{"v5.5.0", "v5.5.0-rc.2", "v5.5.0-rc.1", "v5.5.0-beta.1"}),
				expectedReleaseVersion:    "v6.0.0",
				expectedReleaseBranchName: "release/v6.0",
				expectedBetaVersion:       "v6.0.0-beta.1",
				expectedRCVersion:         "v6.0.0-rc.1",
			},
			{
				releaseType:               "minor",
				repositoryRelease:         newMockRelease("v7.8.1"),
				repositoryTags:            newMockTags([]string{"v7.9.0-rc.2", "v7.9.0-rc.1", "v7.9.0-beta.3", "v7.9.0-beta.2", "v7.9.0-beta.1", "v7.8.1", "v7.8.0", "v7.8.0-rc.1", "v7.8.0-beta.1"}),
				expectedReleaseVersion:    "v7.9.0",
				expectedReleaseBranchName: "release/v7.9",
				expectedBetaVersion:       "v7.9.0-beta.4",
				expectedRCVersion:         "v7.9.0-rc.3",
			},
			{
				releaseType:               "major",
				repositoryRelease:         newMockRelease("v7.8.1"),
				repositoryTags:            newMockTags([]string{"v8.0.0-rc.2", "v8.0.0-rc.1", "v8.0.0-beta.3", "v8.0.0-beta.2", "v8.0.0-beta.1", "v7.8.1", "v7.8.0", "v7.8.0-rc.1", "v7.8.0-beta.1"}),
				expectedReleaseVersion:    "v8.0.0",
				expectedReleaseBranchName: "release/v8.0",
				expectedBetaVersion:       "v8.0.0-beta.4",
				expectedRCVersion:         "v8.0.0-rc.3",
			},
			{
				releaseType:               "patch",
				repositoryRelease:         newMockRelease("v7.8.1"),
				repositoryTags:            newMockTags([]string{"v7.8.2-rc.2", "vv7.8.2-rc.1", "vv7.8.2-beta.3", "vv7.8.2-beta.2", "vv7.8.2-beta.1", "v7.8.1", "v7.8.0", "v7.8.0-rc.1", "v7.8.0-beta.1"}),
				expectedReleaseVersion:    "v7.8.2",
				expectedReleaseBranchName: "release/v7.8",
				expectedBetaVersion:       "v7.8.2-beta.4",
				expectedRCVersion:         "v7.8.2-rc.3",
			},
		}

		for index, testCase := range testCases {
			version := &upVersions{
				githubClient: github.NewClient(newGithubMock(testCase.repositoryRelease, testCase.repositoryTags)),
				ctx:          context.Background(),
				releaseType:  testCase.releaseType,
				githubOrg:    "test",
				githubRepo:   "test",
			}

			fmt.Println("test case:", index+1)
			assert.NoError(t, version.setNextReleaseVersion())
			assert.NoError(t, version.setNextBetaAndRCVersion())
			assert.Equal(t, testCase.expectedReleaseVersion, version.nextReleaseVersion)
			assert.Equal(t, testCase.expectedReleaseBranchName, version.releaseBranchName)
			assert.Equal(t, testCase.expectedBetaVersion, version.nextBetaVersion)
			assert.Equal(t, testCase.expectedRCVersion, version.nextRCVersion)
		}
	})
}
