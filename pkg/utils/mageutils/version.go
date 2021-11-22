// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mageutils

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v40/github"
)

// Possible releases types that can be passed to trigger the UpVersions
const (
	patchAbbreviatedType = "p"
	minorAbbreviatedType = "m"
	majorAbbreviatedType = "M"
	patchType            = "patch"
	minorType            = "minor"
	majorType            = "major"
)

// messages or errors
const (
	invalidReleaseType = "%s isn't a valid release type, please inform a valid release type (p or patch), " +
		"(m or minor), (M or major)"
	missingOrgAndRepositoryName = "missing github organization name and github repository name env vars, " +
		"please set and try again: (%s) - (%s)"
)

// env vars
const (
	envRepositoryOrg  = "HORUSEC_REPOSITORY_ORG"
	envRepositoryName = "HORUSEC_REPOSITORY_NAME"
)

// UpVersions data struct
type upVersions struct {
	githubClient               *github.Client
	ctx                        context.Context
	githubOrg                  string
	githubRepo                 string
	releaseType                string
	actualReleaseVersion       string
	nextReleaseVersion         string
	nextReleaseVersionStripped string
	nextReleaseBranchName      string
	actualRCVersion            string
	nextRCVersion              string
	actualBetaVersion          string
	nextBetaVersion            string
}

// newUpVersions instantiates a new UpVersions command struct.
func newUpVersions(releaseType string) *upVersions {
	return &upVersions{
		githubClient: github.NewClient(nil),
		ctx:          context.Background(),
		releaseType:  releaseType,
		githubOrg:    os.Getenv(envRepositoryOrg),
		githubRepo:   os.Getenv(envRepositoryName),
	}
}

// UpVersions command to up latest version of the repository to the next, including the beta and rc next versions.
// Outputs:
// actualReleaseVersion: represents the actual version of the repository (v1.0.0).
// nextReleaseVersion: represents the next release version (v1.1.0).
// nextReleaseVersionStripped: represents the next release version without the v prefix (1.1.0).
// nextReleaseBranchName: represents the next release branch name (release/v1.1).
// nextBetaVersion: represents the next beta release tag name (v1.1.0-beta.1).
// nextRCVersion: represents the next rc release tag name (v1.1.0-rc.1).
// actualBetaVersion: represents the actual beta release tag name (v1.0.0-beta.1).
// actualRCVersion: represents the actual rc release tag name (v1.0.0-rc.1).
func UpVersions(releaseType string) error {
	version := newUpVersions(releaseType)
	if err := version.isValidVersionCommand(); err != nil {
		return err
	}

	if err := version.setNextReleaseVersion(); err != nil {
		return err
	}

	return version.setNextBetaAndRCVersion()
}

// isValidVersionCommand parse abbreviated release type and check for invalid or missing data.
func (u *upVersions) isValidVersionCommand() error {
	u.parseAbbreviatedReleaseTypeName()

	if u.isInvalidReleaseType() {
		return fmt.Errorf(invalidReleaseType, u.releaseType)
	}

	if u.isInvalidGithubOrgAndRepo() {
		return fmt.Errorf(missingOrgAndRepositoryName, envRepositoryOrg, envRepositoryName)
	}

	return nil
}

// isInvalidReleaseType check for invalid release type.
func (u *upVersions) isInvalidReleaseType() bool {
	return u.releaseType != patchType && u.releaseType != minorType && u.releaseType != majorType
}

// isInvalidGithubOrgAndRepo check for empty GitHub info.
func (u *upVersions) isInvalidGithubOrgAndRepo() bool {
	return u.githubOrg == "" || u.githubRepo == ""
}

// parseAbbreviatedReleaseTypeName parse abbreviated release type to the full name.
func (u *upVersions) parseAbbreviatedReleaseTypeName() {
	switch u.releaseType {
	case majorAbbreviatedType:
		u.releaseType = majorType
	case minorAbbreviatedType:
		u.releaseType = minorType
	case patchAbbreviatedType:
		u.releaseType = patchType
	}
}

// setNextReleaseVersion get the latest release from the GitHub and up the version according to
// the requested release type.
func (u *upVersions) setNextReleaseVersion() error {
	if err := u.getLatestRelease(); err != nil {
		return err
	}

	u.upVersionNextRelease()
	u.outputNextRelease()

	return nil
}

// getLatestRelease get the latest release launched from the GitHub.
func (u *upVersions) getLatestRelease() error {
	release, resp, err := u.githubClient.Repositories.GetLatestRelease(u.ctx, u.githubOrg, u.githubRepo)
	if github.CheckResponse(resp.Response) != nil {
		return err
	}

	u.actualReleaseVersion = *release.TagName

	return nil
}

// upVersionNextRelease increase version according to the requested release type, follows the semantic versioning.
// https://semver.org
func (u *upVersions) upVersionNextRelease() {
	major, minor, patch := u.getSplittedVersionRelease()

	switch u.releaseType {
	case patchType:
		u.nextReleaseVersion = u.setReleaseVersion(major, minor, u.upVersion(patch))
	case minorType:
		u.nextReleaseVersion = u.setReleaseVersion(major, u.upVersion(minor), "0")
	case majorType:
		u.nextReleaseVersion = u.setReleaseVersion(u.upVersion(major), "0", "0")
	}
}

// getSplittedVersionRelease remove all the prefixes amd split the version in the dot character, returning major,
// minor, patch.
func (u *upVersions) getSplittedVersionRelease() (string, string, string) {
	splittedVersion := strings.Split(u.removePrefixes(u.actualReleaseVersion), ".")
	major := splittedVersion[0]
	minor := splittedVersion[1]
	patch := splittedVersion[2]

	return major, minor, patch
}

// setReleaseVersion set the next release branch name and version.
func (u *upVersions) setReleaseVersion(major, minor, patch string) string {
	u.nextReleaseBranchName = fmt.Sprintf("release/v%s.%s", major, minor)
	u.nextReleaseVersionStripped = fmt.Sprintf("%s.%s.%s", major, minor, patch)

	return fmt.Sprintf("%s%s.%s.%s", "v", major, minor, patch)
}

// upVersion parse string to int sum +1 and return it back as a string
func (u *upVersions) upVersion(value string) string {
	valueInt, _ := strconv.Atoi(value)

	return strconv.Itoa(valueInt + 1)
}

// nolint:forbidigo,lll // this isn't a debug statement, link pass the line length
// outputNextRelease set the release branch, version and stripped release output to be available in GitHub actions.
// https://docs.github.com/pt/actions/learn-github-actions/workflow-commands-for-github-actions#setting-an-output-parameter
func (u *upVersions) outputNextRelease() {
	fmt.Printf("::set-output name=actualReleaseVersion::%s\n", u.actualReleaseVersion)

	fmt.Printf("::set-output name=nextReleaseVersion::%s\n", u.nextReleaseVersion)

	fmt.Printf("::set-output name=nextReleaseVersionStripped::%s\n", u.nextReleaseVersionStripped)

	fmt.Printf("::set-output name=nextReleaseBranchName::%s\n", u.nextReleaseBranchName)
}

// listTags list last 50 github tags
func (u *upVersions) listTags() ([]*github.RepositoryTag, error) {
	listOptions := &github.ListOptions{
		Page:    1,
		PerPage: 50,
	}

	tags, resp, err := u.githubClient.Repositories.ListTags(u.ctx, u.githubOrg, u.githubRepo, listOptions)
	if github.CheckResponse(resp.Response) != nil {
		return nil, err
	}

	return tags, nil
}

// getActualBetaAndRC will list last 50 GitHub tags and search for the latest beta/rc, GitHub will return it ordered
// to the newest to the latest.
func (u *upVersions) getActualBetaAndRC() error {
	tags, err := u.listTags()
	if err != nil {
		return err
	}

	u.setLatestBetaAnRC(tags)

	return nil
}

// setLatestBetaAnRC search into returned GitHub tags for the latest rc/beta.
func (u *upVersions) setLatestBetaAnRC(tags []*github.RepositoryTag) {
	for _, tag := range tags {
		u.searchAndSetBetaAndRC(tag)

		if u.actualRCVersion != "" && u.actualBetaVersion != "" {
			break
		}
	}
}

// searchAndSetBetaAndRC for each tag check if it's a beta of rc, ignore if the actual it's already set.
func (u *upVersions) searchAndSetBetaAndRC(tag *github.RepositoryTag) {
	if strings.Contains(tag.GetName(), "beta") && u.actualBetaVersion == "" {
		u.actualBetaVersion = tag.GetName()
	}

	if strings.Contains(tag.GetName(), "rc") && u.actualRCVersion == "" {
		u.actualRCVersion = tag.GetName()
	}
}

// setNextBetaAndRCVersion get the latest beta/rc versions, check for not found beta/rc,
// increasing the actual beta/tc or create a new beta/rc 1.
func (u *upVersions) setNextBetaAndRCVersion() error {
	if err := u.getActualBetaAndRC(); err != nil {
		return err
	}

	if u.actualBetaVersion == "" || u.actualRCVersion == "" {
		u.setFirstRCOrBeta()
	}

	u.upVersionNextBetaAndRc()
	u.outputNextBetaAndRC()

	return nil
}

// setFirstRCOrBeta in case of the repository don't ever had a rc or beta tag create for the fist time ignoring
// the following upVersionNextBetaAndRc function.
func (u *upVersions) setFirstRCOrBeta() {
	if u.actualBetaVersion == "" {
		u.nextBetaVersion = fmt.Sprintf("%s-beta.1", u.nextReleaseVersion)
	}

	if u.actualRCVersion == "" {
		u.nextRCVersion = fmt.Sprintf("%s-rc.1", u.nextReleaseVersion)
	}
}

// upVersionNextBetaAndRc check if the next rc and beta are already defined, if not, check if it is a new version, and
// it is necessary to create a new beta / rc or is it a new version of an existing beta.
func (u *upVersions) upVersionNextBetaAndRc() {
	if u.nextRCVersion == "" {
		if u.isOlderThanActualRelease(u.actualRCVersion) {
			u.nextRCVersion = fmt.Sprintf("%s-rc.1", u.nextReleaseVersion)
		} else {
			u.upVersionRC()
		}
	}

	if u.nextBetaVersion == "" {
		if u.isOlderThanActualRelease(u.actualBetaVersion) {
			u.nextBetaVersion = fmt.Sprintf("%s-beta.1", u.nextReleaseVersion)
		} else {
			u.upVersionBeta()
		}
	}
}

// upVersionBeta increase the actual beta number by 1 and set next beta version, for majors also check for new beta cont
// by verifying if the actual beta it's lower than the actual release.
func (u *upVersions) upVersionBeta() {
	if u.releaseType == majorType && u.isOlderThanActualRelease(u.actualBetaVersion) {
		u.nextBetaVersion = fmt.Sprintf("%s-beta.0", u.nextReleaseVersion)

		return
	}

	major, minor, patch, beta := u.getSplittedVersionBetaOrRC(u.actualBetaVersion)
	u.nextBetaVersion = fmt.Sprintf("v%s.%s.%s-beta.%s", major, minor, patch, u.upVersion(beta))
}

// upVersionRC increase the actual rc number by 1 and set next rc version, for majors also check for new rc cont
// by verifying if the actual RC it's lower than the actual release.
func (u *upVersions) upVersionRC() {
	if u.releaseType == majorType && u.isOlderThanActualRelease(u.actualRCVersion) {
		u.nextRCVersion = fmt.Sprintf("%s-rc.0", u.nextReleaseVersion)

		return
	}

	major, minor, patch, rc := u.getSplittedVersionBetaOrRC(u.actualRCVersion)
	u.nextRCVersion = fmt.Sprintf("v%s.%s.%s-rc.%s", major, minor, patch, u.upVersion(rc))
}

// isOlderThanActualRelease check if the last beta and rc refer to the last release or if it is necessary to
// start a new one. If the current release is greater than or equal, it means that it is necessary to create a new beta,
// otherwise just increase the number of the current version.
func (u *upVersions) isOlderThanActualRelease(version string) bool {
	actualRelease := u.getOnlyNumbersVersion(u.actualReleaseVersion)
	versionToCheck := u.getOnlyNumbersVersion(version)

	return actualRelease >= versionToCheck
}

// getOnlyNumbersVersion remove all prefixes and return version only numbers from the release,
// also ignore the rc and beta integer. Ex: (v1.0.0-rc.1 -> 100)
func (u *upVersions) getOnlyNumbersVersion(version string) int {
	versionReplaced := u.removePrefixes(version)

	versionArray := strings.Split(versionReplaced, ".")

	intValue, _ := strconv.Atoi(versionArray[0] + versionArray[1] + versionArray[2])

	return intValue
}

// getSplittedVersionBetaOrRC remove all the prefixes amd split the version in the dot character, returning major,
// minor, patch beta or rc according to the version informed.
func (u *upVersions) getSplittedVersionBetaOrRC(version string) (string, string, string, string) {
	splittedVersion := strings.Split(u.removePrefixes(version), ".")
	major := splittedVersion[0]
	minor := splittedVersion[1]
	patch := splittedVersion[2]
	betaRC := splittedVersion[3]

	return major, minor, patch, betaRC
}

// nolint:forbidigo,lll // this isn't a debug statement, link pass the line length
// outputNextBetaAndRC set the beta and rc release output to be available in GitHub actions.
// https://docs.github.com/pt/actions/learn-github-actions/workflow-commands-for-github-actions#setting-an-output-paramete
func (u *upVersions) outputNextBetaAndRC() {
	fmt.Printf("::set-output name=actualBetaVersion::%s\n", u.actualBetaVersion)

	fmt.Printf("::set-output name=nextBetaVersion::%s\n", u.nextBetaVersion)

	fmt.Printf("::set-output name=actualRCVersion::%s\n", u.actualRCVersion)

	fmt.Printf("::set-output name=nextRCVersion::%s\n", u.nextRCVersion)
}

// removePrefixes remove all the prefixes from the version, including -rc, -beta, v.
func (u *upVersions) removePrefixes(version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(version, "-rc", ""),
		"-beta", ""), "v", "")
}
