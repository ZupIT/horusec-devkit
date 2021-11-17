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

package parser

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-devkit/pkg/utils/parser/enums"
)

func TestParseBodyToEntity(t *testing.T) {
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

func TestCheckParseBodyToEntityError(t *testing.T) {
	t.Run("should success return eof error", func(t *testing.T) {
		assert.Equal(t, enums.ErrorBodyEmpty, checkParseBodyToEntityError(errors.New("eof")))
	})

	t.Run("should success return eof error", func(t *testing.T) {
		assert.Equal(t, enums.ErrorBodyInvalid, checkParseBodyToEntityError(errors.New("invalid character")))
	})

	t.Run("should success return generic error", func(t *testing.T) {
		assert.Equal(t, errors.New("test"), checkParseBodyToEntityError(errors.New("test")))
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

func TestParseStringToUUID(t *testing.T) {
	t.Run("should success parse string to uuid", func(t *testing.T) {
		id := uuid.New()

		assert.Equal(t, id, ParseStringToUUID(id.String()))
	})
}

func TestParsePacketToEntity(t *testing.T) {
	t.Run("should success parse packet to entity", func(t *testing.T) {
		pkg := packet.NewPacket(&amqp.Delivery{})
		pkg.SetBody((&cli.AnalysisData{RepositoryName: "test"}).ToBytes())

		entity := &cli.AnalysisData{}
		assert.NoError(t, ParsePacketToEntity(pkg, &entity))
		assert.Equal(t, "test", entity.RepositoryName)
	})

	t.Run("should error parsing invalid data", func(t *testing.T) {
		pkg := packet.NewPacket(&amqp.Delivery{})

		entity := &cli.AnalysisData{}
		assert.Error(t, ParsePacketToEntity(pkg, &entity))
	})
}
