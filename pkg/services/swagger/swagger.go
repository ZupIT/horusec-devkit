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

package swagger

import (
	"fmt"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/ZupIT/horusec-devkit/pkg/services/swagger/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type ISwagger interface {
	SetupSwagger()
	GetSwaggerHost() string
}

type Swagger struct {
	router *chi.Mux
	port   string
	host   string
}

func NewSwagger(router *chi.Mux, defaultPort string) ISwagger {
	return &Swagger{
		router: router,
		port:   env.GetEnvOrDefault(enums.EnvHorusecPort, defaultPort),
		host:   env.GetEnvOrDefault(enums.EnvHorusecSwaggerHost, "localhost"),
	}
}

func (s *Swagger) SetupSwagger() {
	s.routerSwagger()

	logger.LogInfo(fmt.Sprintf(enums.MessageSwaggerURL, s.GetSwaggerHost()))
}

func (s *Swagger) routerSwagger() {
	swaggerConfig := httpSwagger.URL(fmt.Sprintf(enums.SwaggerURL, s.host, s.port))

	s.router.Get(enums.SwaggerRoute, httpSwagger.Handler(swaggerConfig))
}

func (s *Swagger) GetSwaggerHost() string {
	return fmt.Sprintf("%s:%s", s.host, s.port)
}
