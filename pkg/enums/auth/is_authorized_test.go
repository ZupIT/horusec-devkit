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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToStringIsAuthorizedType(t *testing.T) {
	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsApplicationAdmin", ApplicationAdmin.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsWorkspaceAdmin", WorkspaceAdmin.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsWorkspaceMember", WorkspaceMember.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsRepositoryAdmin", RepositoryAdmin.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsRepositorySupervisor", RepositorySupervisor.ToString())
	})

	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "IsRepositoryMember", RepositoryMember.ToString())
	})
}

func TestValuesIsAuthorizedType(t *testing.T) {
	t.Run("should return 6 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 6)
	})
}
