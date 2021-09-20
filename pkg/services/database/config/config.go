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
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
)

type IConfig interface {
	SetURI(uri string)
	GetURI() string
	SetLogMode(logMode bool)
	GetLogMode() bool
	Validate() error
}

type Config struct {
	uri     string
	logMode bool
}

func NewDatabaseConfig() IConfig {
	config := &Config{}
	config.SetURI(env.GetEnvOrDefault(enums.EnvRelationalURI,
		"postgresql://root:root@localhost:5432/horusec_db?sslmode=disable"))
	config.SetLogMode(env.GetEnvOrDefaultBool(enums.EnvRelationalLogMode, false))

	return config
}

func (c *Config) SetURI(uri string) {
	c.uri = uri
}

func (c *Config) GetURI() string {
	if strings.Contains(c.uri, enums.DefaultUsernameAndPassword) {
		logger.LogWarn(enums.MessageWarningDefaultDatabaseConnection)
	}
	return c.uri
}

func (c *Config) SetLogMode(logMode bool) {
	c.logMode = logMode
}

func (c *Config) GetLogMode() bool {
	return c.logMode
}

func (c *Config) Validate() error {
	fieldRules := []*validation.FieldRules{
		validation.Field(&c.uri, validation.Required),
	}

	return validation.ValidateStruct(c, fieldRules...)
}
