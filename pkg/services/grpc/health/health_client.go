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
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type ICheckClient interface {
	IsAvailable() (bool, string)
}

type CheckClient struct {
	grpcCon *grpc.ClientConn
}

func NewHealthCheckGrpcClient(grpcCon *grpc.ClientConn) ICheckClient {
	return &CheckClient{
		grpcCon: grpcCon,
	}
}

func (c *CheckClient) IsAvailable() (bool, string) {
	if state := c.grpcCon.GetState(); state != connectivity.Idle && state != connectivity.Ready {
		return false, c.grpcCon.GetState().String()
	}

	return true, c.grpcCon.GetState().String()
}
