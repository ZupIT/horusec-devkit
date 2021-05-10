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
)

func TestNewHealthCheckGrpcServer(t *testing.T) {
	t.Run("should success create a new service", func(t *testing.T) {
		assert.NotNil(t, NewHealthCheckGrpcServer())
	})
}

func TestCheck(t *testing.T) {
	t.Run("should success get check result", func(t *testing.T) {
		service := NewHealthCheckGrpcServer()

		response, err := service.Check(nil, nil)

		assert.NotNil(t, response)
		assert.NoError(t, err)
	})
}

func TestWatch(t *testing.T) {
	t.Run("should panic when no active server", func(t *testing.T) {
		service := NewHealthCheckGrpcServer()

		assert.Panics(t, func() {
			_ = service.Watch(nil, nil)
		})
	})
}
