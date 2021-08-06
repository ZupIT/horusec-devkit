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

package tools

import "github.com/iancoleman/strcase"

type Tool string

const (
	HorusecEngine        Tool = "HorusecEngine"
	GoSec                Tool = "GoSec"
	SecurityCodeScan     Tool = "SecurityCodeScan"
	Brakeman             Tool = "Brakeman"
	Safety               Tool = "Safety"
	Bandit               Tool = "Bandit"
	NpmAudit             Tool = "NpmAudit"
	YarnAudit            Tool = "YarnAudit"
	GitLeaks             Tool = "GitLeaks"
	TfSec                Tool = "TfSec"
	Checkov              Tool = "Checkov"
	Semgrep              Tool = "Semgrep"
	Flawfinder           Tool = "Flawfinder"
	PhpCS                Tool = "PhpCS"
	MixAudit             Tool = "MixAudit"
	Sobelow              Tool = "Sobelow"
	ShellCheck           Tool = "ShellCheck"
	BundlerAudit         Tool = "BundlerAudit"
	OwaspDependencyCheck Tool = "OwaspDependencyCheck"
	DotnetCli            Tool = "DotnetCli"
	Nancy                Tool = "Nancy"
	Trivy                Tool = "Trivy"
)

func (t Tool) ToString() string {
	return string(t)
}

func (t Tool) ToLowerCamel() string {
	return strcase.ToLowerCamel(strcase.ToSnake(t.ToString()))
}

//nolint:funlen // method need to have more then 15 lines
func Values() []Tool {
	return []Tool{
		HorusecEngine,
		GoSec,
		SecurityCodeScan,
		Brakeman,
		Safety,
		Bandit,
		NpmAudit,
		YarnAudit,
		GitLeaks,
		TfSec,
		Checkov,
		Semgrep,
		Flawfinder,
		PhpCS,
		MixAudit,
		Sobelow,
		ShellCheck,
		BundlerAudit,
		OwaspDependencyCheck,
		DotnetCli,
		Nancy,
		Trivy,
	}
}
