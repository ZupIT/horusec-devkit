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
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// WriteCosignKeyToFile executes "echo "$COSIGN_KEY" > $COSIGN_KEY_LOCATION"
func WriteCosignKeyToFile() error {
	mg.Deps(isCosignInstalled, hasAllNecessaryEnvs)

	file, err := os.Create(os.Getenv("COSIGN_KEY_LOCATION"))
	if err != nil {
		return err
	}

	_, err = file.WriteString(os.Getenv("COSIGN_KEY"))

	return err
}

func isCosignInstalled() error {
	return sh.RunV("cosign", "version")
}

//nolint:funlen //this function only has more than 15 lines because of empty lines lint check
func hasAllNecessaryEnvs() error {
	envs := map[string]string{
		"COSIGN_KEY":          os.Getenv("COSIGN_KEY"),
		"COSIGN_KEY_LOCATION": os.Getenv("COSIGN_KEY_LOCATION"),
	}

	var result []string

	for k, v := range envs {
		if v == "" {
			result = append(result, k)
		}
	}

	if len(result) != 0 {
		return fmt.Errorf("missing some env var: %v", result)
	}

	return nil
}
