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
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid claim", func(t *testing.T) {
		claims := &JWTClaims{
			Email:    "test@test.com",
			Username: "test",
			StandardClaims: jwt.StandardClaims{
				Subject: uuid.New().String(),
			},
		}

		assert.NoError(t, claims.Validate())
	})

	t.Run("should return error when invalid", func(t *testing.T) {
		claims := &JWTClaims{}

		assert.Error(t, claims.Validate())
	})
}
