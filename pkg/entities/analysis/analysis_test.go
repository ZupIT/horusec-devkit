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

package analysis

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	vulnerabilityEntities "github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
	analysisEnum "github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
)

func TestGetTableAnalysis(t *testing.T) {
	t.Run("should return analysis database table name", func(t *testing.T) {
		analysis := &Analysis{}

		assert.Equal(t, "analysis", analysis.GetTable())
	})
}

func TestToBytesAnalysis(t *testing.T) {
	t.Run("should parse analysis to bytes", func(t *testing.T) {
		analysis := &Analysis{}

		assert.NotEmpty(t, analysis.ToBytes())
	})
}

func TestGetID(t *testing.T) {
	t.Run("should success get analysis id", func(t *testing.T) {
		analysis := &Analysis{ID: uuid.New()}

		assert.NotEqual(t, uuid.Nil, analysis.GetID())
	})
}

func TestGetIDString(t *testing.T) {
	t.Run("should success get analysis id string", func(t *testing.T) {
		analysis := &Analysis{ID: uuid.New()}

		assert.NotEqual(t, uuid.Nil.String(), analysis.GetIDString())
	})
}

func TestToString(t *testing.T) {
	t.Run("should success parse analysis to string", func(t *testing.T) {
		analysis := &Analysis{}

		assert.NotEmpty(t, analysis.ToString())
	})
}

func TestMap(t *testing.T) {
	t.Run("should return a map of analysis", func(t *testing.T) {
		analysis := &Analysis{}

		assert.NotEmpty(t, analysis.Map())
	})
}

func TestSetFindOneFilter(t *testing.T) {
	t.Run("should set find one filter by analysis", func(t *testing.T) {
		analysis := &Analysis{}

		assert.NotEmpty(t, analysis.SetFindOneFilter())
	})
}

func TestSetError(t *testing.T) {
	t.Run("should success set one error", func(t *testing.T) {
		analysis := &Analysis{}

		analysis.SetError(errors.New("test"))
		assert.Equal(t, "test", analysis.Errors)
	})

	t.Run("should success set many errors", func(t *testing.T) {
		analysis := &Analysis{}

		analysis.SetError(errors.New("test"))
		assert.Equal(t, "test", analysis.Errors)

		analysis.SetError(errors.New("test"))
		assert.Equal(t, "test; test", analysis.Errors)

		analysis.SetError(errors.New("test"))
		assert.Equal(t, "test; test; test", analysis.Errors)
	})
}

func TestSetAllAnalysisVulnerabilitiesDefaultData(t *testing.T) {
	t.Run("should success set data for each vulnerability", func(t *testing.T) {
		analysis := &Analysis{
			ID: uuid.New(),
			AnalysisVulnerabilities: []RelationshipAnalysisVuln{
				{
					Vulnerability: vulnerabilityEntities.Vulnerability{},
				},
				{
					Vulnerability: vulnerabilityEntities.Vulnerability{},
				},
			},
		}

		analysis.SetAllAnalysisVulnerabilitiesDefaultData()
		for _, value := range analysis.AnalysisVulnerabilities {
			assert.NotEqual(t, time.Time{}, value.CreatedAt)
			assert.NotEqual(t, uuid.Nil, value.VulnerabilityID)
			assert.NotEqual(t, uuid.Nil, value.Vulnerability.VulnerabilityID)
			assert.Equal(t, analysis.ID, value.AnalysisID)
		}
	})
}

func TestSetWorkspaceName(t *testing.T) {
	t.Run("should success set workspace name", func(t *testing.T) {
		analysis := &Analysis{}

		analysis.SetWorkspaceName("test")
		assert.Equal(t, "test", analysis.WorkspaceName)
	})
}

func TestSetRepositoryName(t *testing.T) {
	t.Run("should success set repository name", func(t *testing.T) {
		analysis := &Analysis{}

		analysis.SetRepositoryName("test")
		assert.Equal(t, "test", analysis.RepositoryName)
	})
}

func TestSetRepositoryID(t *testing.T) {
	t.Run("should success set repository id", func(t *testing.T) {
		analysis := &Analysis{}

		analysis.SetRepositoryID(uuid.New())
		assert.NotEqual(t, uuid.Nil, analysis.RepositoryID)
	})
}

func TestSetFinishedData(t *testing.T) {
	t.Run("should success set finished analysis data with success status", func(t *testing.T) {
		analysis := &Analysis{}

		analysis.SetFinishedData()
		assert.NotEqual(t, time.Time{}, analysis.FinishedAt)
		assert.Equal(t, analysisEnum.Success, analysis.Status)
	})

	t.Run("should success set finished analysis data with error status", func(t *testing.T) {
		analysis := &Analysis{Errors: "test"}

		analysis.SetFinishedData()
		assert.NotEqual(t, time.Time{}, analysis.FinishedAt)
		assert.Equal(t, analysisEnum.Error, analysis.Status)
	})
}

func TestHasErrors(t *testing.T) {
	t.Run("should return true when the analysis have errors", func(t *testing.T) {
		analysis := &Analysis{Errors: "test"}

		assert.True(t, analysis.HasErrors())
	})

	t.Run("should return false when the analysis do not have errors", func(t *testing.T) {
		analysis := &Analysis{}

		assert.False(t, analysis.HasErrors())
	})
}

func TestGetTotalVulnerabilities(t *testing.T) {
	t.Run("should return a total of 2 vulnerabilities", func(t *testing.T) {
		analysis := &Analysis{
			AnalysisVulnerabilities: []RelationshipAnalysisVuln{
				{
					Vulnerability: vulnerabilityEntities.Vulnerability{},
				},
				{
					Vulnerability: vulnerabilityEntities.Vulnerability{},
				},
			},
		}

		assert.Equal(t, 2, analysis.GetTotalVulnerabilities())
	})
}

func TestGetDataWithoutVulnerabilities(t *testing.T) {
	t.Run("should return anything but analysis vulnerabilities ", func(t *testing.T) {
		analysis := &Analysis{
			ID:             uuid.New(),
			RepositoryID:   uuid.New(),
			RepositoryName: "test",
			WorkspaceID:    uuid.New(),
			WorkspaceName:  "test",
			Status:         "test",
			Errors:         "test",
			CreatedAt:      time.Now(),
			FinishedAt:     time.Now(),
			AnalysisVulnerabilities: []RelationshipAnalysisVuln{
				{
					Vulnerability: vulnerabilityEntities.Vulnerability{},
				},
			},
		}

		result := analysis.GetDataWithoutVulnerabilities()
		assert.Nil(t, result.AnalysisVulnerabilities)
		assert.NotEmpty(t, result.ID)
		assert.NotEmpty(t, result.RepositoryID)
		assert.NotEmpty(t, result.RepositoryName)
		assert.NotEmpty(t, result.WorkspaceID)
		assert.NotEmpty(t, result.WorkspaceName)
		assert.NotEmpty(t, result.Status)
		assert.NotEmpty(t, result.Errors)
		assert.NotEmpty(t, result.CreatedAt)
		assert.NotEmpty(t, result.FinishedAt)
	})
}
