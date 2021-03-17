package health

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type CheckServer struct {
}

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
