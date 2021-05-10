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

package auth

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/enums"
	"github.com/ZupIT/horusec-devkit/pkg/utils/env"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

func NewAuthGRPCConnection() grpc.ClientConnInterface {
	conn, err := makeConnection()
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToAuthGRPC, err)
	}

	return conn
}

func makeConnection() (grpc.ClientConnInterface, error) {
	if env.GetEnvOrDefaultBool(enums.HorusecGRPCConnectionUsesCerts, false) {
		return setupWithCerts()
	}

	return setupWithoutCerts()
}

func setupWithoutCerts() (grpc.ClientConnInterface, error) {
	return grpc.Dial(env.GetEnvOrDefault(enums.HorusecAuthURL, enums.HorusecDefaultAuthHost),
		grpc.WithInsecure())
}

func setupWithCerts() (grpc.ClientConnInterface, error) {
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
