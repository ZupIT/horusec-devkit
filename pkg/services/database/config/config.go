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
	"net/url"
	"strings"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"

	validation "github.com/go-ozzo/ozzo-validation/v4"

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
	uri        string
	username   string
	password   string
	host       string
	port       string
	sslEnabled bool
	dbName     string
	logMode    bool
}

func NewDatabaseConfig() IConfig {
	config := &Config{}
	config.SetURI(env.GetEnvOrDefault(enums.EnvRelationalURI,
		"postgresql://root:root@localhost:5432/horusec_db?sslmode=disable"))
	config.SetLogMode(env.GetEnvOrDefaultBool(enums.EnvRelationalLogMode, false))
	config.username = env.GetEnvOrDefault(enums.EnvRelationalUsername, "root")
	config.password = env.GetEnvOrDefault(enums.EnvRelationalPassword, "root")
	config.host = env.GetEnvOrDefault(enums.EnvRelationalHost, "localhost")
	config.port = env.GetEnvOrDefault(enums.EnvRelationalPort, "5432")
	config.dbName = env.GetEnvOrDefault(enums.EnvRelationalDBName, "horusec_db")
	config.sslEnabled = env.GetEnvOrDefaultBool(enums.EnvRelationalSSLEnable, false)
	return config
}

func (c *Config) SetURI(uri string) {
	if uri != "" {
		c.mountEscapedURI(uri)
	} else {
		var sslEnabled string
		if c.sslEnabled {
			sslEnabled = "enable"
		} else {
			sslEnabled = "disable"
		}
		c.uri = strings.Join([]string{
			"postgresql://", url.PathEscape(c.username), ":", url.PathEscape(c.password), "@",
			c.host, ":", c.port, "/", c.dbName, "?", "sslmode=", sslEnabled},
			"")
	}
}

func (c *Config) mountEscapedURI(uri string) {
	holder := strings.Split(uri, "//")
	if len(holder) <= 1 {
		logger.LogPanic(enums.ErrorInvalidURI.Error(), enums.ErrorInvalidURI)
	}
	holder2 := strings.Split(holder[1], "@")
	userAndPassword := strings.Split(holder2[0], ":")
	c.password = url.PathEscape(userAndPassword[1])
	c.username = userAndPassword[0]
	result := strings.Join([]string{holder[0], "//", c.username, ":", c.password, "@", holder2[1]}, "")
	c.uri = result
}

func (c *Config) GetURI() string {
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
