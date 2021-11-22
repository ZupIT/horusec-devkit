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
	"fmt"
	"log"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// CreateAlphaTag executes "git", "tag", "-f", "-s", "v0.0.0-alpha", "-m", "v0.0.0-alpha"
func CreateAlphaTag() error {
	mg.Deps(isGitExistent)

	return sh.RunV("git", "tag", "-f", "-s", "v0.0.0-alpha", "-m", "v0.0.0-alpha")
}

// GitPushAlpha executes "git", "push", "origin", "-f", "v0.0.0-alpha"
func GitPushAlpha() error {
	mg.Deps(isGitExistent)

	return sh.RunV("git", "push", "origin", "-f", "v0.0.0-alpha")
}

// CreateLocalTag executes "git", "tag", "-s", tag, "-m", "release "+tag
func CreateLocalTag(tag string) (err error) {
	mg.Deps(isGitExistent)
	mg.Deps(mg.F(isValidTag, tag))

	return sh.RunV("git", "tag", "-s", tag, "-m", "release "+tag)
}

// CheckoutReleaseBranch creates if not exists a release branch and then checkout
func CheckoutReleaseBranch(branchName string) error {
	mg.Deps(isGitExistent)

	if err := sh.RunV("git", "checkout", branchName); err != nil {
		log.Printf("First %s release, creating release branch", branchName)

		return sh.RunV("git", "checkout", "-b", branchName)
	}

	return nil
}

// GitPushAll executes "git", "push", "--all" and "git", "push", "--tags"
func GitPushAll() error {
	mg.Deps(isGitExistent)

	if err := sh.RunV("git", "push", "--all"); err != nil {
		return err
	}

	return sh.RunV("git", "push", "--tags")
}

// GitConfig configures global email and user for git
func GitConfig(email, name string) error {
	mg.Deps(isGitExistent)

	if err := sh.RunV("git", "config", "--global", "user.email", email); err != nil {
		return err
	}

	return sh.RunV("git", "config", "--global", "user.name", name)
}

// DefaultGitConfig sets horusec as global git user and horusec@zup.com.br as global git email
func DefaultGitConfig() error {
	return GitConfig("horusec@zup.com.br", "horusec")
}

func isGitExistent() error {
	return sh.RunV("git", "version")
}

func isValidTag(tag string) error {
	validTagLength := 4
	// TODO(ian) : we should make a more assertive check on passed tags
	if len(tag) < validTagLength {
		return fmt.Errorf("invalid tag format: %s", tag)
	}

	return nil
}
