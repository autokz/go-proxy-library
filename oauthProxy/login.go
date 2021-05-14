package oauthProxy

import (
	"encoding/json"
)

func (p *Proxy) Login(username, password string) (authData, userID string, err error) {
	loginURL := p.BaseURL + p.OAuthURL

	dto := make(map[string]interface{})

	dto["grant_type"] = p.OAuthGrantType
	dto["username"] = username
	dto["password"] = password
	dto["domain"] = p.DomainURL
	dto["refresh_token"] = ""

	result, headers, statusCode := p.Client.Post(loginURL, dto, nil)
	if statusCode != 200 {
		authData = string(result)
		return
	}

	// JWT
	var jwt map[string]interface{}
	json.Unmarshal(result, &jwt)

	// User Payload
	uPayload := make(map[string]string)
	json.Unmarshal([]byte(headers.Get("X-User-Payload")), &uPayload)

	authData = p.Converter.FromJWTToFrontend(jwt)
	userID = uPayload["id"]
	return
}
