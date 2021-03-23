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
	"time"

	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/vulnerability"
)

type RelationshipAnalysisVuln struct {
	VulnerabilityID uuid.UUID                   `json:"vulnerabilityID" gorm:"Column:vulnerability_id"`
	AnalysisID      uuid.UUID                   `json:"analysisID" gorm:"Column:analysis_id"`
	CreatedAt       time.Time                   `json:"createdAt" gorm:"Column:created_at"`
	Vulnerability   vulnerability.Vulnerability `json:"vulnerabilities" gorm:"foreignKey:VulnerabilityID;references:VulnerabilityID"` //nolint:lll // notations need more than 130 characters
}

func (a *RelationshipAnalysisVuln) GetTable() string {
	return "analysis_vulnerabilities"
}

func (a *RelationshipAnalysisVuln) SetCreatedAt() {
	a.CreatedAt = time.Now()
}

func (a *RelationshipAnalysisVuln) SetVulnerabilityID() {
	a.Vulnerability.GenerateID()
	a.VulnerabilityID = a.Vulnerability.VulnerabilityID
}

func (a *RelationshipAnalysisVuln) SetAnalysisID(id uuid.UUID) {
	a.AnalysisID = id
}

func (a *RelationshipAnalysisVuln) GetAnalysisVulnerabilitiesWithoutVulnerability() *RelationshipAnalysisVuln {
	return &RelationshipAnalysisVuln{
		VulnerabilityID: a.VulnerabilityID,
		AnalysisID:      a.AnalysisID,
		CreatedAt:       a.CreatedAt,
	}
}
