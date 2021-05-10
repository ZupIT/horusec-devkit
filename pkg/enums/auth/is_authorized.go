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

type AuthorizationType string

const (
	ApplicationAdmin     AuthorizationType = "IsApplicationAdmin"
	WorkspaceAdmin       AuthorizationType = "IsWorkspaceAdmin"
	WorkspaceMember      AuthorizationType = "IsWorkspaceMember"
	RepositoryAdmin      AuthorizationType = "IsRepositoryAdmin"
	RepositorySupervisor AuthorizationType = "IsRepositorySupervisor"
	RepositoryMember     AuthorizationType = "IsRepositoryMember"
)

func (i AuthorizationType) ToString() string {
	return string(i)
}

func Values() []AuthorizationType {
	return []AuthorizationType{
		ApplicationAdmin,
		WorkspaceAdmin,
		WorkspaceMember,
		RepositoryAdmin,
		RepositorySupervisor,
		RepositoryMember,
	}
}
