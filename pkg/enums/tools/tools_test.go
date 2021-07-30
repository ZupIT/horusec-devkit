// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	t.Run("should success parse all tools to string", func(t *testing.T) {
		assert.Equal(t, "HorusecEngine", HorusecEngine.ToString())
		assert.Equal(t, "GoSec", GoSec.ToString())
		assert.Equal(t, "SecurityCodeScan", SecurityCodeScan.ToString())
		assert.Equal(t, "Brakeman", Brakeman.ToString())
		assert.Equal(t, "Safety", Safety.ToString())
		assert.Equal(t, "Bandit", Bandit.ToString())
		assert.Equal(t, "NpmAudit", NpmAudit.ToString())
		assert.Equal(t, "YarnAudit", YarnAudit.ToString())
		assert.Equal(t, "GitLeaks", GitLeaks.ToString())
		assert.Equal(t, "TfSec", TfSec.ToString())
		assert.Equal(t, "Semgrep", Semgrep.ToString())
		assert.Equal(t, "Flawfinder", Flawfinder.ToString())
		assert.Equal(t, "PhpCS", PhpCS.ToString())
		assert.Equal(t, "MixAudit", MixAudit.ToString())
		assert.Equal(t, "Sobelow", Sobelow.ToString())
		assert.Equal(t, "ShellCheck", ShellCheck.ToString())
		assert.Equal(t, "BundlerAudit", BundlerAudit.ToString())
		assert.Equal(t, "OwaspDependencyCheck", OwaspDependencyCheck.ToString())
		assert.Equal(t, "DotnetCli", DotnetCli.ToString())
		assert.Equal(t, "Nancy", Nancy.ToString())
		assert.Equal(t, "Trivy", Trivy.ToString())
	})
}

func TestToLowerCase(t *testing.T) {
	t.Run("should success parse to lower camel", func(t *testing.T) {
		assert.Equal(t, "horusecEngine", HorusecEngine.ToLowerCamel())
		assert.Equal(t, "goSec", GoSec.ToLowerCamel())
		assert.Equal(t, "securityCodeScan", SecurityCodeScan.ToLowerCamel())
		assert.Equal(t, "brakeman", Brakeman.ToLowerCamel())
		assert.Equal(t, "safety", Safety.ToLowerCamel())
		assert.Equal(t, "bandit", Bandit.ToLowerCamel())
		assert.Equal(t, "npmAudit", NpmAudit.ToLowerCamel())
		assert.Equal(t, "yarnAudit", YarnAudit.ToLowerCamel())
		assert.Equal(t, "gitLeaks", GitLeaks.ToLowerCamel())
		assert.Equal(t, "tfSec", TfSec.ToLowerCamel())
		assert.Equal(t, "semgrep", Semgrep.ToLowerCamel())
		assert.Equal(t, "flawfinder", Flawfinder.ToLowerCamel())
		assert.Equal(t, "phpCs", PhpCS.ToLowerCamel())
		assert.Equal(t, "mixAudit", MixAudit.ToLowerCamel())
		assert.Equal(t, "sobelow", Sobelow.ToLowerCamel())
		assert.Equal(t, "shellCheck", ShellCheck.ToLowerCamel())
		assert.Equal(t, "bundlerAudit", BundlerAudit.ToLowerCamel())
		assert.Equal(t, "owaspDependencyCheck", OwaspDependencyCheck.ToLowerCamel())
		assert.Equal(t, "dotnetCli", DotnetCli.ToLowerCamel())
		assert.Equal(t, "nancy", Nancy.ToLowerCamel())
		assert.Equal(t, "trivy", Trivy.ToLowerCamel())
	})
}

func TestValues(t *testing.T) {
	t.Run("should return 21 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 21)
	})
}
