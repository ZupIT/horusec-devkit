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

package http

import (
	"compress/flate"
	"net/http"
	"time"

	"github.com/go-chi/cors"

	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
)

type Server struct {
	port       string
	timeout    time.Duration
	corsConfig *cors.Options
}

func NewServerConfigService(defaultPort string, corsConfig *cors.Options) *Server {
	return &Server{
		port:       env.GetEnvOrDefault("HORUSEC_PORT", defaultPort),
		corsConfig: corsConfig,
	}
}

func (s *Server) Cors(next http.Handler) http.Handler {
	return cors.New(*s.corsConfig).Handler(next)
}

func (s *Server) SetTimeout(timeInSeconds int) {
	s.timeout = time.Duration(timeInSeconds) * time.Second
}

func (s *Server) GetTimeout() time.Duration {
	return s.timeout
}

func (s *Server) GetCompression() int {
	return flate.BestCompression
}

func (s *Server) GetPort() string {
	return s.port
}

func (s *Server) SetPort(port string) {
	s.port = port
}
