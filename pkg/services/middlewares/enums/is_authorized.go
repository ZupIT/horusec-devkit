package enums

type IsAuthorizedType string

const (
	ApplicationAdmin     IsAuthorizedType = "IsApplicationAdmin"
	CompanyAdmin         IsAuthorizedType = "IsCompanyAdmin"
	CompanyMember        IsAuthorizedType = "IsCompanyMember"
	RepositoryAdmin      IsAuthorizedType = "IsRepositoryAdmin"
	RepositorySupervisor IsAuthorizedType = "IsRepositorySupervisor"
	RepositoryMember     IsAuthorizedType = "IsRepositoryMember"
)

func (i IsAuthorizedType) ToString() string {
	return string(i)
}
