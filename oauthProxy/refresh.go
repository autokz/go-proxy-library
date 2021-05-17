package oauthProxy

import (
	"encoding/json"
)

func (p *Proxy) Refresh(authData string) (string, error) {
	dto := p.Converter.FromFrontendToJWT(authData)

	// Getting new JWT with User Payload
	newAuthData, _, err := p.Client.Auth(dto, p.OAuthRefreshGrantType)
	if err != nil {
		return "",  err
	}

	// Converting JWT...
	var jwt map[string]interface{}
	json.Unmarshal([]byte(newAuthData), &jwt)
	result := p.Converter.FromJWTToFrontend(jwt)

	return result, nil
}
