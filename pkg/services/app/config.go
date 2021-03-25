package app

import (
	"context"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/app/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type IConfig interface {
	IsBrokerDisabled() bool
	IsApplicationAdmEnabled() bool
	GetAuthorizationType() auth.AuthorizationType
}

type Config struct {
	authGRPC proto.AuthServiceClient
	context  context.Context
	*proto.GetAuthConfigResponse
}

func NewAppConfig(authGRPC proto.AuthServiceClient) IConfig {
	appConfig := &Config{
		authGRPC:              authGRPC,
		context:               context.Background(),
		GetAuthConfigResponse: &proto.GetAuthConfigResponse{},
	}

	return appConfig.getAuthConfig()
}

func (c *Config) getAuthConfig() IConfig {
	response, err := c.authGRPC.GetAuthConfig(c.context, &proto.GetAuthConfigData{})
	if err != nil {
		logger.LogPanic(enums.FailedToGetAuthConfigGRPCRequest, err)
	}

	c.GetAuthConfigResponse = response
	return c
}

func (c *Config) IsBrokerDisabled() bool {
	return c.GetDisableBroker()
}

func (c *Config) IsApplicationAdmEnabled() bool {
	return c.GetEnableApplicationAdmin()
}

func (c *Config) GetAuthorizationType() auth.AuthorizationType {
	return auth.AuthorizationType(c.GetAuthType())
}
