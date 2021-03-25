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

func TestIsBrokerDisabled(t *testing.T) {
	t.Run("should return true for disabled broker", func(t *testing.T) {
		config := &Config{
			GetAuthConfigResponse: &proto.GetAuthConfigResponse{
				DisableBroker: true,
			},
		}

		assert.True(t, config.IsBrokerDisabled())
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

		assert.Equal(t, auth.Ldap, config.GetAuthorizationType())
	})
}
