package auth

type IsAuthorizedType string

const (
	ApplicationAdmin     IsAuthorizedType = "IsApplicationAdmin"
	WorkspaceAdmin       IsAuthorizedType = "IsWorkspaceAdmin"
	WorkspaceMember      IsAuthorizedType = "IsWorkspaceMember"
	RepositoryAdmin      IsAuthorizedType = "IsRepositoryAdmin"
	RepositorySupervisor IsAuthorizedType = "IsRepositorySupervisor"
	RepositoryMember     IsAuthorizedType = "IsRepositoryMember"
)

func (i IsAuthorizedType) ToString() string {
	return string(i)
}

func Values() []IsAuthorizedType {
	return []IsAuthorizedType{
		ApplicationAdmin,
		WorkspaceAdmin,
		WorkspaceMember,
		RepositoryAdmin,
		RepositorySupervisor,
		RepositoryMember,
	}
}
