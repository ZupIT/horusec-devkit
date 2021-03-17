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
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
)

type IConfig interface {
	Validate() error
	GetHost() string
	SetHost(host string)
	GetPort() string
	SetPort(port string)
	GetUsername() string
	SetUsername(username string)
	GetPassword() string
	SetPassword(password string)
	GetConnectionString() string
}

type Config struct {
	host     string
	port     string
	username string
	password string
}

func NewBrokerConfig() IConfig {
	config := &Config{}
	config.SetHost(env.GetEnvOrDefault(enums.EnvBrokerHost, "127.0.0.1"))
	config.SetPort(env.GetEnvOrDefault(enums.EnvBrokerPort, "5672"))
	config.SetUsername(env.GetEnvOrDefault(enums.EnvBrokerUsername, "guest"))
	config.SetPassword(env.GetEnvOrDefault(enums.EnvBrokerPassword, "guest"))

	return config
}

func (c *Config) Validate() error {
	fieldRules := []*validation.FieldRules{
		validation.Field(&c.host, validation.Required),
		validation.Field(&c.port, validation.Required),
		validation.Field(&c.username, validation.Required),
		validation.Field(&c.password, validation.Required),
	}

	return validation.ValidateStruct(c, fieldRules...)
}

func (c *Config) GetHost() string {
	return c.host
}

func (c *Config) SetHost(host string) {
	c.host = host
}

func (c *Config) GetPort() string {
	return c.port
}

func (c *Config) SetPort(port string) {
	c.port = port
}

func (c *Config) GetUsername() string {
	return c.username
}

func (c *Config) SetUsername(username string) {
	c.username = username
}

func (c *Config) GetPassword() string {
	return c.password
}

func (c *Config) SetPassword(password string) {
	c.password = password
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		c.GetUsername(),
		c.GetPassword(),
		c.GetHost(),
		c.GetPort(),
	)
}
