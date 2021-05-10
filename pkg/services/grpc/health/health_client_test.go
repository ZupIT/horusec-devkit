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
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestNewHealthCheckGrpcClient(t *testing.T) {
	t.Run("should success create a new health check service", func(t *testing.T) {
		service := NewHealthCheckGrpcClient(&grpc.ClientConn{})

		assert.NotNil(t, service)
	})
}

func TestIsAvailable(t *testing.T) {
	t.Run("should return true when healthy", func(t *testing.T) {
		connection, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
		assert.NoError(t, err)

		service := NewHealthCheckGrpcClient(connection)
		isAvailable, status := service.IsAvailable()

		assert.True(t, isAvailable)
		assert.NotEmpty(t, status)
	})

	t.Run("should return false when unhealthy", func(t *testing.T) {
		connection, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
		assert.NoError(t, err)

		err = connection.Close()
		assert.NoError(t, err)

		service := NewHealthCheckGrpcClient(connection)
		isAvailable, status := service.IsAvailable()

		assert.False(t, isAvailable)
		assert.NotEmpty(t, status)
	})
}
