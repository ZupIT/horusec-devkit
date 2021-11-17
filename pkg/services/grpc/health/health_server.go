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

package health

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type CheckServer struct{}

func NewHealthCheckGrpcServer() *CheckServer {
	return &CheckServer{}
}

func (c *CheckServer) Check(context.Context,
	*grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	logger.LogInfo("sending the grpc check server request for health check")

	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (c *CheckServer) Watch(_ *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	logger.LogInfo("sending the grpc watch request for health check")

	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}
