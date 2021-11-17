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

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// CherryPick Check for the dependencies than exec the cherry pick. When the workflow it's triggered to run on main,
// the commit it's going to be made directly to the main branch, so it's not needed to run this.
// Output:
// cherryPickBranchName: name of the branch that the updating versioning files was cherry picked.
// commitShaToPick: commit sha that was cherry picked to the new branch.
func CherryPick() error {
	mg.Deps(isGitExistent)

	if err := DefaultGitConfig(); err != nil {
		return err
	}

	actualBranch, err := getActualBranch()
	if err != nil || actualBranch == "main" {
		return err
	}

	return execCherryPick()
}

// getLatestCommitSha get the latest commit sha made in the workflow, witch in the workflow context should be the
// updating versioning files commit.
func getLatestCommitSha() (string, error) {
	githubSha, err := sh.Output("git", "log", "-1", "--format=%H")
	if err != nil {
		return "", err
	}

	return githubSha, nil
}

// nolint:forbidigo,lll  // this isn't a debug statement, link pass the line length
// createBranchFromOriginMain create a branch from the origin main and print the output to be used in the GitHub actions.
// https://docs.github.com/pt/actions/learn-github-actions/workflow-commands-for-github-actions#setting-an-output-paramete
func createBranchFromOriginMain(githubSha string) (string, error) {
	branchName := fmt.Sprintf("automatic/updating-versioning-files-%s", githubSha)
	fmt.Printf("::set-output name=cherryPickBranchName::%s\n", branchName)
	fmt.Printf("::set-output name=commitShaToPick::%s\n", githubSha)

	if err := sh.RunV("git", "checkout", "-b", branchName, "origin/main"); err != nil {
		return "", err
	}

	return branchName, nil
}

// cherryPickAndPush Cherry pick the release branch commit updating versioning files into the new branch.
// The cherry pick strategy always accept incoming to avoid conflicts, and store tha sha of the commit made in the
// release branch (-x). After that the branch is pushed to the GitHub.
func cherryPickAndPush(githubSha, branchName string) error {
	if err := sh.RunV("git", "cherry-pick", githubSha, "-x", "--strategy-option", "theirs"); err != nil {
		return err
	}

	return sh.RunV("git", "push", "-u", "origin", branchName)
}

// getActualBranch get the actual branch that the workflow it's running.
func getActualBranch() (string, error) {
	return sh.Output("git", "rev-parse", "--abbrev-ref", "HEAD")
}

// execCherryPick get the latest commit made in the workflow and cherry pick it to a new branch.
func execCherryPick() error {
	githubSha, err := getLatestCommitSha()
	if err != nil {
		return err
	}

	branchName, err := createBranchFromOriginMain(githubSha)
	if err != nil {
		return err
	}

	return cherryPickAndPush(githubSha, branchName)
}
