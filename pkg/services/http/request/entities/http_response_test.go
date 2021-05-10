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

package entities

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/http/request/enums"
)

func TestGetBodyBytes(t *testing.T) {
	t.Run("should success get request body", func(t *testing.T) {
		response := &HTTPResponse{Response: &http.Response{Body: ioutil.NopCloser(strings.NewReader("test"))}}

		bytes, err := response.GetBodyBytes()
		assert.NoError(t, err)
		assert.NotEmpty(t, bytes)
	})
}

func TestGetStatusCodeString(t *testing.T) {
	t.Run("should success get status code string", func(t *testing.T) {
		response := &HTTPResponse{Response: &http.Response{StatusCode: 200}}

		assert.Equal(t, http.StatusText(http.StatusOK), response.GetStatusCodeString())
	})
}

func TestGetContentType(t *testing.T) {
	t.Run("should success get content type", func(t *testing.T) {
		header := http.Header{}
		header.Add(enums.ContentType, "test")

		response := &HTTPResponse{Response: &http.Response{Header: header}}

		assert.Equal(t, "test", response.GetContentType())
	})
}

func TestGetStatusCode(t *testing.T) {
	t.Run("should success get status code", func(t *testing.T) {
		response := &HTTPResponse{Response: &http.Response{StatusCode: 200}}

		assert.Equal(t, http.StatusOK, response.GetStatusCode())
	})
}

func TestCloseBody(t *testing.T) {
	t.Run("should success close request body", func(t *testing.T) {
		response := &HTTPResponse{Response: &http.Response{Body: ioutil.NopCloser(strings.NewReader("test"))}}

		assert.NotPanics(t, func() {
			response.CloseBody()
		})
	})
}

func TestErrorByStatusCode(t *testing.T) {
	t.Run("should return server error when status code 500", func(t *testing.T) {
		response := &HTTPResponse{Response: &http.Response{StatusCode: http.StatusInternalServerError}}

		err := response.ErrorByStatusCode()
		assert.Error(t, err)
		assert.Equal(t, enums.ErrorRequestServerError, err)
	})

	t.Run("should return client error when status code is between 400 and 499", func(t *testing.T) {
		response := &HTTPResponse{Response: &http.Response{StatusCode: http.StatusBadRequest}}

		err := response.ErrorByStatusCode()
		assert.Error(t, err)
		assert.Equal(t, enums.ErrorRequestClientError, err)
	})

	t.Run("should return nil when success status code", func(t *testing.T) {
		response := &HTTPResponse{Response: &http.Response{StatusCode: http.StatusOK}}

		assert.NoError(t, response.ErrorByStatusCode())
	})
}
