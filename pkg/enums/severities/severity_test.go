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

package severities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	t.Run("should success parse severity to string", func(t *testing.T) {
		assert.Equal(t, "CRITICAL", Critical.ToString())
		assert.Equal(t, "HIGH", High.ToString())
		assert.Equal(t, "MEDIUM", Medium.ToString())
		assert.Equal(t, "LOW", Low.ToString())
		assert.Equal(t, "UNKNOWN", Unknown.ToString())
		assert.Equal(t, "INFO", Info.ToString())
	})
}

func TestIsValid(t *testing.T) {
	t.Run("should success check if severity is valid", func(t *testing.T) {
		assert.True(t, Critical.IsValid())
		assert.True(t, High.IsValid())
		assert.True(t, Medium.IsValid())
		assert.True(t, Low.IsValid())
		assert.True(t, Info.IsValid())
		assert.True(t, Unknown.IsValid())
	})
}

func TestMap(t *testing.T) {
	t.Run("should success parse to map", func(t *testing.T) {
		mapValues := Map()
		assert.NotEmpty(t, mapValues)

		assert.Equal(t, Critical, mapValues["CRITICAL"])
		assert.Equal(t, High, mapValues["HIGH"])
		assert.Equal(t, Medium, mapValues["MEDIUM"])
		assert.Equal(t, Low, mapValues["LOW"])
		assert.Equal(t, Unknown, mapValues["UNKNOWN"])
		assert.Equal(t, Info, mapValues["INFO"])
	})
}

func TestGetSeverityByString(t *testing.T) {
	t.Run("should success get severity value by string", func(t *testing.T) {
		assert.Equal(t, Critical, GetSeverityByString("CRITICAL"))
		assert.Equal(t, High, GetSeverityByString("HIGH"))
		assert.Equal(t, Medium, GetSeverityByString("MEDIUM"))
		assert.Equal(t, Low, GetSeverityByString("LOW"))
		assert.Equal(t, Unknown, GetSeverityByString("UNKNOWN"))
		assert.Equal(t, Info, GetSeverityByString("INFO"))
	})

	t.Run("should return unknown when invalid string value", func(t *testing.T) {
		assert.Equal(t, Unknown, GetSeverityByString("test"))
	})
}

func TestContains(t *testing.T) {
	t.Run("should return true when contains value", func(t *testing.T) {
		assert.True(t, Contains("CRITICAL"))
		assert.True(t, Contains("HIGH"))
		assert.True(t, Contains("MEDIUM"))
		assert.True(t, Contains("LOW"))
		assert.True(t, Contains("INFO"))
	})

	t.Run("should return false when not contains value", func(t *testing.T) {
		assert.False(t, Contains("false"))
	})
}

func TestValues(t *testing.T) {
	t.Run("should return 6 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 6)
	})
}
