package client

type AccessToken string
type UserPayload map[string]interface{}
type AuthData map[string]interface{}

type OAuthClientInterface interface {
	Check(aToken AccessToken) error
	Auth(accessData map[string]interface{}, authType string) (newAuthData AuthData, userPayload UserPayload, err error)
	Logout(aToken AccessToken) error
	LogoutAll(aToken AccessToken) error
}
