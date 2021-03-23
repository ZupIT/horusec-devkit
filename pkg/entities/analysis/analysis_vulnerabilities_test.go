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
		analysisVulnerabilities := &RelationshipAnalysisVuln{}

		assert.Equal(t, "analysis_vulnerabilities", analysisVulnerabilities.GetTable())
	})
}

func TestSetCreatedAtAnalysisVulnerabilities(t *testing.T) {
	t.Run("should success set created at", func(t *testing.T) {
		analysisVulnerabilities := &RelationshipAnalysisVuln{}

		analysisVulnerabilities.SetCreatedAt()
		assert.NotEqual(t, time.Time{}, analysisVulnerabilities.CreatedAt)
	})
}

func TestSetVulnerabilityIDAnalysisVulnerabilities(t *testing.T) {
	t.Run("should success set vuln id in analysis vulnerabilities and vulnerability", func(t *testing.T) {
		analysisVulnerabilities := &RelationshipAnalysisVuln{}

		analysisVulnerabilities.SetVulnerabilityID()
		assert.NotEqual(t, uuid.Nil, analysisVulnerabilities.VulnerabilityID)
		assert.NotEqual(t, uuid.Nil, analysisVulnerabilities.Vulnerability.VulnerabilityID)
	})
}

func TestSetAnalysisIDAnalysisVulnerabilities(t *testing.T) {
	t.Run("should success set vuln id in analysis vulnerabilities and vulnerability", func(t *testing.T) {
		analysisVulnerabilities := &RelationshipAnalysisVuln{}

		analysisVulnerabilities.SetAnalysisID(uuid.New())
		assert.NotEqual(t, uuid.Nil, analysisVulnerabilities.AnalysisID)
	})
}

func TestGetAnalysisVulnerabilitiesWithoutVulnerability(t *testing.T) {
	t.Run("should success set vuln id in analysis vulnerabilities and vulnerability", func(t *testing.T) {
		analysisVulnerabilities := &RelationshipAnalysisVuln{
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
