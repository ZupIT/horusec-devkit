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
	patchAbbreviated = "p"
	minorAbbreviated = "m"
	majorAbbreviated = "M"
	patch            = "patch"
	minor            = "minor"
	major            = "major"
)

const (
	invalidReleaseType = "%s isn't a valid release type, please inform a valid release type (p or patch), " +
		"(m or minor), (M or major)"
	missingOrgAndRepositoryName = "missing github organization name and github repository name env vars, " +
		"please set and try again: (%s) - (%s)"
)

const (
	envRepositoryOrg  = "HORUSEC_REPOSITORY_ORG"
	envRepositoryName = "HORUSEC_REPOSITORY_NAME"
)

type upVersions struct {
	githubClient         *github.Client
	ctx                  context.Context
	githubOrg            string
	githubRepo           string
	releaseType          string
	actualReleaseVersion string
	nextReleaseVersion   string
	releaseBranchName    string
	actualRCVersion      string
	nextRCVersion        string
	actualBetaVersion    string
	nextBetaVersion      string
}

func newUpVersions(releaseType string) *upVersions {
	return &upVersions{
		githubClient: github.NewClient(nil),
		ctx:          context.Background(),
		releaseType:  releaseType,
		githubOrg:    os.Getenv(envRepositoryOrg),
		githubRepo:   os.Getenv(envRepositoryName),
	}
}

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

func (u *upVersions) isInvalidReleaseType() bool {
	return u.releaseType != patch && u.releaseType != minor && u.releaseType != major
}

func (u *upVersions) isInvalidGithubOrgAndRepo() bool {
	return u.githubOrg == "" || u.githubRepo == ""
}

func (u *upVersions) parseAbbreviatedReleaseTypeName() {
	switch u.releaseType {
	case majorAbbreviated:
		u.releaseType = major
	case minorAbbreviated:
		u.releaseType = minor
	case patchAbbreviated:
		u.releaseType = patch
	}
}

func (u *upVersions) setNextReleaseVersion() error {
	if err := u.getLatestRelease(); err != nil {
		return err
	}

	u.upVersionNextRelease()
	u.outputNextRelease()

	return nil
}

func (u *upVersions) getLatestRelease() error {
	release, resp, err := u.githubClient.Repositories.GetLatestRelease(u.ctx, u.githubOrg, u.githubRepo)
	if github.CheckResponse(resp.Response) != nil {
		return err
	}

	u.actualReleaseVersion = *release.TagName

	return nil
}

func (u *upVersions) upVersionNextRelease() {
	major, minor, patch := u.getSplittedVersionRelease()

	switch u.releaseType {
	case patch:
		u.nextReleaseVersion = u.setReleaseVersion(major, minor, u.upVersion(patch))
	case minor:
		u.nextReleaseVersion = u.setReleaseVersion(major, u.upVersion(minor), "0")
	case major:
		u.nextReleaseVersion = u.setReleaseVersion(u.upVersion(major), "0", "0")
	}
}

func (u *upVersions) getSplittedVersionRelease() (string, string, string) {
	splittedVersion := strings.Split(u.actualReleaseVersion, ".")
	major := strings.ReplaceAll(splittedVersion[0], "v", "")
	minor := splittedVersion[1]
	patch := splittedVersion[2]

	return major, minor, patch
}

func (u *upVersions) setReleaseVersion(major, minor, patch string) string {
	u.releaseBranchName = fmt.Sprintf("release/v%s.%s", major, minor)

	return fmt.Sprintf("%s%s.%s.%s", "v", major, minor, patch)
}

func (u *upVersions) upVersion(value string) string {
	valueInt, _ := strconv.Atoi(value)

	return strconv.Itoa(valueInt + 1)
}

func (u *upVersions) outputNextRelease() {
	fmt.Printf("::set-output name=releaseVersion::%s\n", u.nextReleaseVersion)

	fmt.Printf("::set-output name=strippedReleaseVersion::%s\n",
		strings.ReplaceAll(u.nextReleaseVersion, "v", ""))

	fmt.Printf("::set-output name=releaseBranchName::%s\n", u.releaseBranchName)
}

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

func (u *upVersions) getActualBetaAndRC() error {
	tags, err := u.listTags()
	if err != nil {
		return err
	}

	u.setLatestBetaAnRC(tags)

	return nil
}

func (u *upVersions) setLatestBetaAnRC(tags []*github.RepositoryTag) {
	for _, tag := range tags {
		u.searchAndSetBetaAndRC(tag)
		if u.actualRCVersion != "" && u.actualBetaVersion != "" {
			break
		}
	}
}

func (u *upVersions) searchAndSetBetaAndRC(tag *github.RepositoryTag) {
	if strings.Contains(tag.GetName(), "beta") && u.actualBetaVersion == "" {
		u.actualBetaVersion = tag.GetName()
	}

	if strings.Contains(tag.GetName(), "rc") && u.actualRCVersion == "" {
		u.actualRCVersion = tag.GetName()
	}
}

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

func (u *upVersions) setFirstRCOrBeta() {
	if u.actualBetaVersion == "" {
		u.nextBetaVersion = fmt.Sprintf("%s-beta.1", u.nextReleaseVersion)
	}

	if u.actualRCVersion == "" {
		u.nextRCVersion = fmt.Sprintf("%s-rc.1", u.nextReleaseVersion)
	}
}

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

func (u *upVersions) upVersionBeta() {
	major, minor, patch, beta := u.getSplittedVersionBetaOrRC(u.actualBetaVersion)

	u.nextBetaVersion = fmt.Sprintf("v%s.%s.%s-beta.%s", major, minor, patch, u.upVersion(beta))
}

func (u *upVersions) upVersionRC() {
	major, minor, patch, rc := u.getSplittedVersionBetaOrRC(u.actualRCVersion)

	u.nextRCVersion = fmt.Sprintf("v%s.%s.%s-rc.%s", major, minor, patch, u.upVersion(rc))
}

func (u *upVersions) isOlderThanActualRelease(version string) bool {
	actualRelease := u.getOnlyNumbersVersion(u.actualReleaseVersion)
	versionToCheck := u.getOnlyNumbersVersion(version)

	return actualRelease >= versionToCheck
}

func (u *upVersions) getOnlyNumbersVersion(version string) int {
	versionReplaced := u.removePrefixes(version)

	versionArray := strings.Split(versionReplaced, ".")

	intValue, _ := strconv.Atoi(versionArray[0] + versionArray[1] + versionArray[2])

	return intValue
}

func (u *upVersions) getSplittedVersionBetaOrRC(version string) (string, string, string, string) {
	splittedVersion := strings.Split(u.removePrefixes(version), ".")
	major := splittedVersion[0]
	minor := splittedVersion[1]
	patch := splittedVersion[2]
	betaRC := splittedVersion[3]

	return major, minor, patch, betaRC
}

func (u *upVersions) outputNextBetaAndRC() {
	fmt.Printf("::set-output name=betaVersion::%s\n", u.nextBetaVersion)

	fmt.Printf("::set-output name=rcVersion::%s\n", u.nextRCVersion)
}

func (u *upVersions) removePrefixes(version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(version, "-rc", ""),
		"-beta", ""), "v", "")
}
