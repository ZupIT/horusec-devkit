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

package app

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
	"github.com/ZupIT/horusec-devkit/pkg/services/grpc/auth/proto"
)

func TestSetConnections(t *testing.T) {
	t.Run("should panic when failed to create a new app config", func(t *testing.T) {
		assert.Panics(t, func() {
			NewAppConfig(proto.NewAuthServiceClient(&grpc.ClientConn{}))
		})
	})
}

func TestGetAuthConfig(t *testing.T) {
	t.Run("should success get and set auth config", func(t *testing.T) {
		mock := &proto.Mock{}

		mock.On("GetAuthConfig").Return(&proto.GetAuthConfigResponse{}, nil)

		config := &Config{
			authGRPC: mock,
		}

		config.getAuthConfig()
		assert.NotNil(t, config.GetAuthConfigResponse)
	})

	t.Run("should panic when failed to get auth config", func(t *testing.T) {
		mock := &proto.Mock{}

		mock.On("GetAuthConfig").Return(&proto.GetAuthConfigResponse{}, errors.New("test"))

		config := &Config{
			authGRPC: mock,
		}

		assert.Panics(t, func() {
			config.getAuthConfig()
		})
	})
}

func TestIsEmailsDisabled(t *testing.T) {
	t.Run("should return true when emails are not enabled", func(t *testing.T) {
		config := &Config{
			GetAuthConfigResponse: &proto.GetAuthConfigResponse{
				DisableEmails: true,
			},
		}

		assert.True(t, config.IsEmailsDisabled())
	})
}

func TestIsApplicationAdmEnabled(t *testing.T) {
	t.Run("should return true for application admin enabled", func(t *testing.T) {
		config := &Config{
			GetAuthConfigResponse: &proto.GetAuthConfigResponse{
				EnableApplicationAdmin: true,
			},
		}

		assert.True(t, config.IsApplicationAdmEnabled())
	})
}

func TestGetAuthorizationType(t *testing.T) {
	t.Run("should return ldap auth type", func(t *testing.T) {
		config := &Config{
			GetAuthConfigResponse: &proto.GetAuthConfigResponse{
				AuthType: auth.Ldap.ToString(),
			},
		}

		assert.Equal(t, auth.Ldap, config.GetAuthenticationType())
	})
}
