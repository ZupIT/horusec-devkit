package auth

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

func NewAuthGRPCConnection() *grpc.ClientConn {
	conn, err := makeConnection()
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToAuthGRPC, err)
	}

	return conn
}

func makeConnection() (*grpc.ClientConn, error) {
	if env.GetEnvOrDefaultBool(enums.HorusecGRPCConnectionUsesCerts, false) {
		return setupWithCerts()
	}

	return setupWithoutCerts()
}

func setupWithoutCerts() (*grpc.ClientConn, error) {
	return grpc.Dial(env.GetEnvOrDefault(enums.HorusecAuthURL, enums.HorusecDefaultAuthHost),
		grpc.WithInsecure())
}

func setupWithCerts() (*grpc.ClientConn, error) {
	return grpc.Dial(env.GetEnvOrDefault(enums.HorusecAuthURL, enums.HorusecDefaultAuthHost),
		grpc.WithTransportCredentials(getCredentials()))
}

func getCredentials() credentials.TransportCredentials {
	cred, err := credentials.NewClientTLSFromFile(
		env.GetEnvOrDefault(enums.HorusecGRPCCertificatePath, ""), "")
	if err != nil {
		logger.LogError(enums.MessageFailedToGetGRPCCerts, err)
	}

	return cred
}
