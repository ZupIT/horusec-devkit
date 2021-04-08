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
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/enums/analysis"
)

type Analysis struct {
	ID                      uuid.UUID                  `json:"id" gorm:"Column:analysis_id"`
	RepositoryID            uuid.UUID                  `json:"repositoryID" gorm:"Column:repository_id"`
	RepositoryName          string                     `json:"repositoryName" gorm:"Column:repository_name"`
	WorkspaceID             uuid.UUID                  `json:"workspaceID" gorm:"Column:workspace_id"`
	WorkspaceName           string                     `json:"workspaceName" gorm:"Column:workspace_name"`
	Status                  analysis.Status            `json:"status" gorm:"Column:status"`
	Errors                  string                     `json:"errors" gorm:"Column:errors"`
	CreatedAt               time.Time                  `json:"createdAt" gorm:"Column:created_at"`
	FinishedAt              time.Time                  `json:"finishedAt" gorm:"Column:finished_at"`
	AnalysisVulnerabilities []RelationshipAnalysisVuln `json:"analysisVulnerabilities" gorm:"foreignKey:AnalysisID;references:ID;polymorphicValue:analysis_vulnerabilities;polymorphic:analysis_vulnerabilities"` //nolint:lll // notations need more than 130 characters
}

func (a *Analysis) GetTable() string {
	return "analysis"
}

func (a *Analysis) ToBytes() []byte {
	bytes, _ := json.Marshal(a)
	return bytes
}

func (a *Analysis) GetID() uuid.UUID {
	return a.ID
}

func (a *Analysis) GetIDString() string {
	return a.ID.String()
}

func (a *Analysis) ToString() string {
	return string(a.ToBytes())
}

func (a *Analysis) Map() map[string]interface{} {
	return map[string]interface{}{
		"id":                      a.ID,
		"createdAt":               a.CreatedAt,
		"repositoryID":            a.RepositoryID,
		"repositoryName":          a.RepositoryName,
		"workspaceName":           a.WorkspaceName,
		"workspaceID":             a.WorkspaceID,
		"status":                  a.Status,
		"errors":                  a.Errors,
		"finishedAt":              a.FinishedAt,
		"analysisVulnerabilities": a.AnalysisVulnerabilities,
	}
}

func (a *Analysis) SetFindOneFilter() map[string]interface{} {
	return map[string]interface{}{"id": a.GetID()}
}

func (a *Analysis) SetError(err error) {
	if err != nil {
		toAppend := ""
		if len(a.Errors) > 0 {
			a.Errors += "; " + err.Error()
			return
		}

		a.Errors += toAppend + err.Error()
	}
}

func (a *Analysis) SetAllAnalysisVulnerabilitiesDefaultData() {
	for key := range a.AnalysisVulnerabilities {
		a.AnalysisVulnerabilities[key].SetCreatedAt()
		a.AnalysisVulnerabilities[key].SetAnalysisID(a.ID)
		a.AnalysisVulnerabilities[key].SetVulnerabilityID()
	}
}

func (a *Analysis) SetWorkspaceName(workspaceName string) {
	a.WorkspaceName = workspaceName
}

func (a *Analysis) SetRepositoryName(repositoryName string) {
	a.RepositoryName = repositoryName
}

func (a *Analysis) SetRepositoryID(repositoryID uuid.UUID) {
	a.RepositoryID = repositoryID
}

func (a *Analysis) SetFinishedData() {
	a.FinishedAt = time.Now()

	if a.HasErrors() {
		a.Status = analysis.Error
		return
	}

	a.Status = analysis.Success
}

func (a *Analysis) HasErrors() bool {
	return len(a.Errors) > 0
}

func (a *Analysis) GetTotalVulnerabilities() int {
	return len(a.AnalysisVulnerabilities)
}

func (a *Analysis) GetDataWithoutVulnerabilities() *Analysis {
	return &Analysis{
		ID:             a.ID,
		RepositoryID:   a.RepositoryID,
		RepositoryName: a.RepositoryName,
		WorkspaceID:    a.WorkspaceID,
		WorkspaceName:  a.WorkspaceName,
		Status:         a.Status,
		Errors:         a.Errors,
		CreatedAt:      a.CreatedAt,
		FinishedAt:     a.FinishedAt,
	}
}
