package user

type RegistrationService interface {
	SignUp(userName string, password string) Dto
	SignIn(userName string, password string) string
	ReadToken(token string) User
}
