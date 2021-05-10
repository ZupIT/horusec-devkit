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

package analysis

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
)

func TestGetTableAnalysisVulnerabilities(t *testing.T) {
	t.Run("should success get database table name", func(t *testing.T) {
		analysisVulnerabilities := &AnalysisVulnerabilities{}

		assert.Equal(t, "analysis_vulnerabilities", analysisVulnerabilities.GetTable())
	})
}

func TestSetCreatedAtAnalysisVulnerabilities(t *testing.T) {
	t.Run("should success set created at", func(t *testing.T) {
		analysisVulnerabilities := &AnalysisVulnerabilities{}

		analysisVulnerabilities.SetCreatedAt()
		assert.NotEqual(t, time.Time{}, analysisVulnerabilities.CreatedAt)
	})
}

func TestSetVulnerabilityIDAnalysisVulnerabilities(t *testing.T) {
	t.Run("should success set vuln id in analysis vulnerabilities and vulnerability", func(t *testing.T) {
		analysisVulnerabilities := &AnalysisVulnerabilities{}

		analysisVulnerabilities.SetVulnerabilityID()
		assert.NotEqual(t, uuid.Nil, analysisVulnerabilities.VulnerabilityID)
		assert.NotEqual(t, uuid.Nil, analysisVulnerabilities.Vulnerability.VulnerabilityID)
	})
}

func TestSetAnalysisIDAnalysisVulnerabilities(t *testing.T) {
	t.Run("should success set vuln id in analysis vulnerabilities and vulnerability", func(t *testing.T) {
		analysisVulnerabilities := &AnalysisVulnerabilities{}

		analysisVulnerabilities.SetAnalysisID(uuid.New())
		assert.NotEqual(t, uuid.Nil, analysisVulnerabilities.AnalysisID)
	})
}

func TestGetAnalysisVulnerabilitiesWithoutVulnerability(t *testing.T) {
	t.Run("should success set vuln id in analysis vulnerabilities and vulnerability", func(t *testing.T) {
		analysisVulnerabilities := &AnalysisVulnerabilities{
			VulnerabilityID: uuid.New(),
			AnalysisID:      uuid.New(),
			CreatedAt:       time.Now(),
			Vulnerability:   vulnerability.Vulnerability{},
		}

		result := analysisVulnerabilities.GetAnalysisVulnerabilitiesWithoutVulnerability()
		assert.Empty(t, result.Vulnerability)
		assert.NotEmpty(t, result.VulnerabilityID)
		assert.NotEmpty(t, result.AnalysisID)
		assert.NotEmpty(t, result.CreatedAt)
	})
}
