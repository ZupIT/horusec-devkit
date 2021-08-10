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

package tracer

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer/enums"
)

func TestMiddleware(t *testing.T) {

	handler200 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	handler400 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}
	handler500 := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	handlerPanic := func(w http.ResponseWriter, r *http.Request) {
		panic("error")
	}

	err := os.Setenv(enums.HorusecJaegerName, "test")
	assert.NoError(t, err)
	err = os.Setenv(enums.HorusecJaegerLogInfo, "true")
	assert.NoError(t, err)
	err = os.Setenv(enums.HorusecJaegerLogError, "true")
	assert.NoError(t, err)

	j, err := NewJaeger()
	assert.NoError(t, err)
	closer, err := j.Config(true)
	assert.NoError(t, err)
	defer func() {
		err := closer.Close()
		assert.NoError(t, err)
	}()
	t.Run("should get a status 200 ", func(t *testing.T) {
		handler := Tracer(http.HandlerFunc(handler200))

		req := httptest.NewRequest("GET", "/", nil)
		ctx := AddOperationName(req.Context(), "TestOperationName")
		req = req.WithContext(ctx)
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})
	t.Run("should get a status 400 ", func(t *testing.T) {
		handler := Tracer(http.HandlerFunc(handler400))

		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("should get a status 500 ", func(t *testing.T) {
		handler := Tracer(http.HandlerFunc(handler500))

		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
	t.Run("should get a status 200 and skip trace when swagger is present on path ", func(t *testing.T) {
		handler := Tracer(http.HandlerFunc(handler200))

		req := httptest.NewRequest("GET", "/swagger", nil)
		res := httptest.NewRecorder()

		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})
	t.Run("should panic", func(t *testing.T) {
		handler := Tracer(http.HandlerFunc(handlerPanic))

		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

}
