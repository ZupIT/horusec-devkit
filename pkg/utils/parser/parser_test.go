package parser

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
)

func TestNewAccountCompanyFromReadCLoser(t *testing.T) {
	t.Run("should success parse body to entity with no errors", func(t *testing.T) {
		analysisData := cli.AnalysisData{RepositoryName: "test"}
		response := &cli.AnalysisData{}

		body := ioutil.NopCloser(strings.NewReader(string(analysisData.ToBytes())))

		assert.NoError(t, ParseBodyToEntity(body, response))
		assert.NotNil(t, response)
		assert.Equal(t, "test", response.RepositoryName)
	})

	t.Run("should return error when failed to parse body", func(t *testing.T) {
		response := &cli.AnalysisData{}

		body := ioutil.NopCloser(strings.NewReader(""))

		assert.Error(t, ParseBodyToEntity(body, response))
		assert.Empty(t, response)
	})
}

func TestParseEntityToIOReadCloser(t *testing.T) {
	t.Run("should success parse entity to io read closer", func(t *testing.T) {
		entity := &cli.AnalysisData{RepositoryName: "test"}

		bytes, err := ParseEntityToIOReadCloser(entity)
		assert.NoError(t, err)
		assert.NotEmpty(t, bytes)
	})

	t.Run("should return error when failed to parse entity to bytes", func(t *testing.T) {
		bytes, err := ParseEntityToIOReadCloser(make(chan string))
		assert.Error(t, err)
		assert.Nil(t, bytes)
	})
}
