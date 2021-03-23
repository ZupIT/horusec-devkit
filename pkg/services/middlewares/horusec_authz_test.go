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

package middlewares

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	authGRPC "github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt"
	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/entities"
)

func testHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func createValidToken() string {
	token, _, _ := jwt.CreateToken(&entities.TokenData{
		Email:     "test@test.com",
		Username:  "test",
		AccountID: uuid.New(),
	}, nil)

	return token
}

func TestNewAuthzMiddleware(t *testing.T) {
	t.Run("should success create a new middleware service", func(t *testing.T) {
		assert.NotNil(t, NewAuthzMiddleware(&grpc.ClientConn{}))
	})
}

func TestIsCompanyMember(t *testing.T) {
	t.Run("should return 200 when valid request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to verify request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, errors.New("test"))

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 401 unauthorized request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 unauthorized request with invalid token", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", "test")

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestIsCompanyAdmin(t *testing.T) {
	t.Run("should return 200 when valid request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to verify request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, errors.New("test"))

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 401 unauthorized request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 unauthorized request with invalid token", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsCompanyAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", "test")

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestIsRepositoryMember(t *testing.T) {
	t.Run("should return 200 when valid request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to verify request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, errors.New("test"))

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 401 unauthorized request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 unauthorized request with invalid token", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryMember(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", "test")

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestIsRepositorySupervisor(t *testing.T) {
	t.Run("should return 200 when valid request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositorySupervisor(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to verify request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, errors.New("test"))

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositorySupervisor(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 401 unauthorized request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositorySupervisor(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 unauthorized request with invalid token", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositorySupervisor(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", "test")

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestIsRepositoryAdmin(t *testing.T) {
	t.Run("should return 200 when valid request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to verify request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, errors.New("test"))

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 401 unauthorized request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 unauthorized request with invalid token", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsRepositoryAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", "test")

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestIsApplicationAdmin(t *testing.T) {
	t.Run("should return 200 when valid request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: true}, nil)
		grpcMock.On("GetAuthConfig").Return(&authGRPC.
			GetAuthConfigResponse{AuthType: "test", EnableApplicationAdmin: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsApplicationAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("should return 500 when failed to verify request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, errors.New("test"))
		grpcMock.On("GetAuthConfig").Return(&authGRPC.
			GetAuthConfigResponse{AuthType: "test", EnableApplicationAdmin: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsApplicationAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("should return 401 unauthorized request", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)
		grpcMock.On("GetAuthConfig").Return(&authGRPC.
			GetAuthConfigResponse{AuthType: "test", EnableApplicationAdmin: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsApplicationAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", createValidToken())

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 401 unauthorized request with invalid token", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("IsAuthorized").
			Return(&authGRPC.IsAuthorizedResponse{IsAuthorized: false}, nil)
		grpcMock.On("GetAuthConfig").Return(&authGRPC.
			GetAuthConfigResponse{AuthType: "test", EnableApplicationAdmin: true}, nil)

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsApplicationAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Add("X-Horusec-Authorization", "test")

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("should return 500 when failed to get auth config", func(t *testing.T) {
		grpcMock := &authGRPC.Mock{}

		grpcMock.On("GetAuthConfig").Return(&authGRPC.GetAuthConfigResponse{}, errors.New("test"))

		middleware := AuthzMiddleware{
			grpcClient: grpcMock,
		}

		handler := middleware.IsApplicationAdmin(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
