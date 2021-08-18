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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
)

func TestNewConfig(t *testing.T) {
	t.Run("should success create config with default values", func(t *testing.T) {
		databaseConfig := NewDatabaseConfig()

		assert.Equal(t, false, databaseConfig.GetLogMode())
		assert.Equal(t, "postgresql://root:root@localhost:5432/horusec_db?sslmode=disable",
			databaseConfig.GetURI())
	})

	t.Run("should success create config with custom values", func(t *testing.T) {
		_ = os.Setenv(enums.EnvRelationalURI, "postgresql://root123:root321@localhost:1432/horusec_db?sslmode=enable")
		_ = os.Setenv(enums.EnvRelationalLogMode, "true")

		databaseConfig := NewDatabaseConfig()

		assert.Equal(t, true, databaseConfig.GetLogMode())
		assert.Equal(t, "postgresql://root123:root321@localhost:1432/horusec_db?sslmode=enable", databaseConfig.GetURI())
	})
}

func TestGetAndSetURI(t *testing.T) {
	t.Run("should success set and get dialect", func(t *testing.T) {
		databaseConfig := NewDatabaseConfig()
		databaseConfig.SetURI("postgresql://root123:root321@localhost:1432/horusec_db?sslmode=enable")

		assert.Equal(t, "postgresql://root123:root321@localhost:1432/horusec_db?sslmode=enable", databaseConfig.GetURI())
	})
	t.Run("should panic when set an invalid URI", func(t *testing.T) {
		databaseConfig := NewDatabaseConfig()

		assert.Panics(t, func() {
			databaseConfig.SetURI("test")
		})
	})
	t.Run("should success return default values when set URI with empty value", func(t *testing.T) {
		databaseConfig := NewDatabaseConfig()
		databaseConfig.SetURI("")
		assert.Equal(t, "postgresql://root:root@localhost:5432/horusec_db?sslmode=disable",
			databaseConfig.GetURI())
	})
	t.Run("should success return default values whith ssl enable when set URI with sslEnabled", func(t *testing.T) {
		_ = os.Setenv(enums.EnvRelationalSSLEnable, "true")
		databaseConfig := NewDatabaseConfig()
		databaseConfig.SetURI("")
		assert.Equal(t, "postgresql://root:root@localhost:5432/horusec_db?sslmode=enable",
			databaseConfig.GetURI())
	})
}

func TestGetAndSetLogMode(t *testing.T) {
	t.Run("should success set and get dialect", func(t *testing.T) {
		databaseConfig := NewDatabaseConfig()
		databaseConfig.SetLogMode(true)

		assert.Equal(t, true, databaseConfig.GetLogMode())
	})
}

func TestValidate(t *testing.T) {
	t.Run("should return no error when valid config", func(t *testing.T) {
		databaseConfig := NewDatabaseConfig()

		assert.NoError(t, databaseConfig.Validate())
	})

	t.Run("should return error when invalid config", func(t *testing.T) {
		databaseConfig := &Config{}

		assert.Error(t, databaseConfig.Validate())
	})
}
