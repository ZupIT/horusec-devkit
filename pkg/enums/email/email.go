package email

type Template string

const (
	AccountConfirmation Template = "account-confirmation"
	ResetPassword       Template = "reset-password"
	OrganizationInvite  Template = "organization-invite"
	RepositoryInvite    Template = "repository-invite"
)
