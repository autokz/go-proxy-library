package client

type AccessToken string

type OAuthClientInterface interface {
	Check(aToken AccessToken) error
	Auth(accessData map[string]interface{}, authType string) (newAuthData string, userPayload map[string]interface{}, err error)
	Logout(aToken AccessToken) error
	LogoutAll(aToken AccessToken) error
}
