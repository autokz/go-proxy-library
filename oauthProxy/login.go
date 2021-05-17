package oauthProxy

import (
	"encoding/json"
)

func (p *Proxy) Login(username, password string) (string, map[string]interface{}, error) {
	dto := make(map[string]interface{})

	dto["username"] = username
	dto["password"] = password

	// Getting new JWT with User Payload
	newAuthData, userPayload, err := p.Client.Auth(dto, p.OAuthGrantType)
	if err != nil {
		return "", nil, err
	}

	// Converting JWT...
	var jwt map[string]interface{}
	json.Unmarshal([]byte(newAuthData), &jwt)
	authData := p.Converter.FromJWTToFrontend(jwt)


	return authData, userPayload, nil
}
