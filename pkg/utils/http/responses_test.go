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

package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/http/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

func TestStatusOK(t *testing.T) {
	t.Run("should return status code 200", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		StatusOK(w, "ok")

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestStatusNoContent(t *testing.T) {
	t.Run("should return status code 204", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodDelete, "/test", nil)
		w := httptest.NewRecorder()

		StatusNoContent(w)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}

func TestStatusBadRequest(t *testing.T) {
	t.Run("should return status code 400", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		StatusBadRequest(w, errors.New("test"))

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestStatusNotFound(t *testing.T) {
	t.Run("should return status code 404", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		StatusNotFound(w, errors.New("test"))

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestStatusConflict(t *testing.T) {
	t.Run("should return status code 409", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusConflict(w, errors.New("test"))

		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

func TestStatusInternalServerError(t *testing.T) {
	t.Run("should return status code 500", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusInternalServerError(w, errors.New("test"))

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
	t.Run("should return status code 500 with body of invalid health broker", func(t *testing.T) {
		output := bytes.NewBufferString("")
		_ = logger.LogSetOutput(output, nil)
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusInternalServerError(w, enums.ErrorBrokerIsNotHealth)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, output.String(), enums.ErrorBrokerIsNotHealth.Error())
	})
	t.Run("should return status code 500 with body of invalid health database", func(t *testing.T) {
		output := bytes.NewBufferString("")
		_ = logger.LogSetOutput(output, nil)
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusInternalServerError(w, enums.ErrorDatabaseIsNotHealth)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, output.String(), enums.ErrorDatabaseIsNotHealth.Error())
	})
	t.Run("should return status code 500 with body of invalid health grpc", func(t *testing.T) {
		output := bytes.NewBufferString("")
		_ = logger.LogSetOutput(output, nil)
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusInternalServerError(w, enums.ErrorGrpcIsNotHealth)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, output.String(), enums.ErrorGrpcIsNotHealth.Error())
	})
}

func TestStatusCreated(t *testing.T) {
	t.Run("should return status code 201", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusCreated(w, "ok")

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestStatusUnauthorized(t *testing.T) {
	t.Run("should return status code 401", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusUnauthorized(w, errors.New("test"))

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestStatusForbidden(t *testing.T) {
	t.Run("should return status code 403", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusForbidden(w, errors.New("test"))

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}

func TestStatusMethodNotAllowed(t *testing.T) {
	t.Run("should return status code 405", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusMethodNotAllowed(w, errors.New("test"))

		assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
	})
}

func TestGetErrorMessage(t *testing.T) {
	t.Run("should return error message message", func(t *testing.T) {
		result := getErrorMessage(errors.New("test error"))
		assert.NotEmpty(t, result)
		assert.Equal(t, "test error", result)
	})

	t.Run("should return empty string when error is nil", func(t *testing.T) {
		result := getErrorMessage(nil)
		assert.Empty(t, result)
	})
}

func TestStatusUnprocessableEntity(t *testing.T) {
	t.Run("should return status code 422", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodPost, "/test", nil)
		w := httptest.NewRecorder()

		StatusUnprocessableEntity(w, errors.New("test"))

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
}
