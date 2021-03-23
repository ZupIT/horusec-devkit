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

package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("should success return a hash of password with no errors", func(t *testing.T) {
		hash, err := HashPasswordBcrypt("test")

		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("should return true for valid password", func(t *testing.T) {
		result := CheckPasswordHashBcrypt("test",
			"$2a$10$CY6dyOjKD6rG.PxA6QlrLeUaHR.SD5VWLbkvc4YJM1ZT39geAIZQG")

		assert.True(t, result)
	})

	t.Run("should return false for invalid password", func(t *testing.T) {
		result := CheckPasswordHashBcrypt("invalid",
			"$2a$10$CY6dyOjKD6rG.PxA6QlrLeUaHR.SD5VWLbkvc4YJM1ZT39geAIZQG")

		assert.False(t, result)
	})
}

func TestGenerateSHA256(t *testing.T) {
	t.Run("should generate a sha256 without errors", func(t *testing.T) {
		hash := GenerateSHA256("test")

		assert.NotEmpty(t, hash)
	})

	t.Run("should generate a hash with many strings without errors", func(t *testing.T) {
		hash := GenerateSHA256("test1", "test2", "test3")

		assert.NotEmpty(t, hash)
	})

	t.Run("should generate two equal hashes without errors", func(t *testing.T) {
		hash := GenerateSHA256("test")
		hashExpected := GenerateSHA256("test")

		assert.Equal(t, hashExpected, hash)
	})

	t.Run("should generate two different hashes without errors", func(t *testing.T) {
		hash := GenerateSHA256("test1")
		hashExpected := GenerateSHA256("test2")

		assert.NotEqual(t, hashExpected, hash)
	})
}
