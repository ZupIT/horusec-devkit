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

package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	t.Run("should return 6 valid queue values", func(t *testing.T) {
		assert.Len(t, Values(), 10)
	})
}

func TestIsInvalid(t *testing.T) {
	t.Run("should return false for valid queue value", func(t *testing.T) {
		assert.False(t, HorusecEmail.IsInvalid())
	})

	t.Run("should return true for invalid queue value", func(t *testing.T) {
		assert.True(t, Queue("test").IsInvalid())
	})
}

func TestIsValid(t *testing.T) {
	t.Run("should return true for valid value", func(t *testing.T) {
		assert.True(t, HorusecEmail.IsValid())
	})

	t.Run("should return false for valid value", func(t *testing.T) {
		assert.False(t, Queue("test").IsValid())
	})
}

func TestValueOf(t *testing.T) {
	t.Run("should return value of horusec email queue", func(t *testing.T) {
		assert.Equal(t, HorusecEmail, ValueOf("horusec-email"))
	})

	t.Run("should return unknown for invalid value", func(t *testing.T) {
		assert.Equal(t, Queue(""), ValueOf("test"))
	})
}

func TestIsEqual(t *testing.T) {
	t.Run("should return true for equal value", func(t *testing.T) {
		assert.True(t, HorusecEmail.IsEqual("horusec-email"))
	})

	t.Run("should return false for equal value", func(t *testing.T) {
		assert.False(t, HorusecEmail.IsEqual("test"))
	})
}

func TestToString(t *testing.T) {
	t.Run("should return parse queue to string", func(t *testing.T) {
		assert.Equal(t, "horusec-email", HorusecEmail.ToString())
	})
}
