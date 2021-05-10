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
