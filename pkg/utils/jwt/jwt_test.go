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

package jwt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/utils/jwt/entities"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func TestCreateToken(t *testing.T) {
	t.Run("should success create a signed token with no errors", func(t *testing.T) {
		account := &entities.TokenData{
			Email:     "test@test.com",
			Username:  "test",
			AccountID: uuid.New(),
		}

		token, _, err := CreateToken(account, nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}

func TestDecodeToken(t *testing.T) {
	t.Run("should success decode token", func(t *testing.T) {
		account := &entities.TokenData{
			Email:     "test@test.com",
			Username:  "test",
			AccountID: uuid.New(),
		}

		token, _, err := CreateToken(account, nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := DecodeToken(token)
		assert.NoError(t, err)
		assert.NoError(t, claims.Valid())
	})

	t.Run("should return error invalid signature", func(t *testing.T) {
		account := &entities.TokenData{
			Email:     "test@test.com",
			Username:  "test",
			AccountID: uuid.New(),
		}

		token, _, err := CreateToken(account, nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		_ = os.Setenv("HORUSEC_JWT_SECRET_KEY", "test")

		_, err = DecodeToken(token)
		assert.Error(t, err)
		assert.Equal(t, "signature is invalid", err.Error())
	})
}

func TestAuthMiddleware(t *testing.T) {
	t.Run("should return 200 when valid token", func(t *testing.T) {
		handler := AuthMiddleware(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		token, _, _ := CreateToken(&entities.TokenData{
			AccountID: uuid.New(),
			Email:     "test@test.com",
			Username:  "test",
		}, nil)

		req.Header.Set("Authorization", "Bearer "+token)

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 401 when invalid jwt token", func(t *testing.T) {
		handler := AuthMiddleware(http.HandlerFunc(testHandler))

		req, _ := http.NewRequest("GET", "http://test", nil)

		req.Header.Set("X-Horusec-Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxM"+
			"jM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}

func TestGetAccountIDByJWTToken(t *testing.T) {
	t.Run("should success return a account ID with no error", func(t *testing.T) {
		token, _, _ := CreateToken(&entities.TokenData{
			AccountID: uuid.New(),
			Email:     "test@test.com",
			Username:  "test",
		}, nil)

		accountID, err := GetAccountIDByJWTToken(token)
		assert.NoError(t, err)
		assert.NotEmpty(t, accountID)
	})

	t.Run("should success return a account ID with token containing bearer", func(t *testing.T) {
		token, _, _ := CreateToken(&entities.TokenData{
			AccountID: uuid.New(),
			Email:     "test@test.com",
			Username:  "test",
		}, nil)

		accountID, err := GetAccountIDByJWTToken(fmt.Sprintf("Bearer %s", token))
		assert.NoError(t, err)
		assert.NotEmpty(t, accountID)
	})

	t.Run("should return error parsing invalid token", func(t *testing.T) {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QiLCJ1c2VybmFtZSI6InRlc3QiLCJzdWIiOiJ0" +
			"ZXN0In0.TMmUwiXEHkgtEX3Fxudu_GP1f9ZmRnkJfzlSuKaaH-o"

		accountID, err := GetAccountIDByJWTToken(token)
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, accountID)
	})
}

func TestCreateRefreshToken(t *testing.T) {
	t.Run("should success create a refresh token", func(t *testing.T) {
		refreshToken := CreateRefreshToken()
		assert.NotEmpty(t, refreshToken)
	})
}
