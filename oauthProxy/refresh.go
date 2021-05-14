package oauthProxy

import (
	"encoding/json"
	"fmt"
)

func (p *Proxy) Refresh(authData string) (string, error) {
	jwt := p.Converter.FromFrontendToJWT(authData)

	var rToken string
	if rToken = ""; jwt["refresh_token"] != nil {
		rToken = jwt["refresh_token"].(string)
	}

	loginURL := p.BaseURL + p.OAuthURL

	dto := make(map[string]interface{})

	dto["grant_type"] = p.OAuthRefreshGrantType
	dto["username"] = ""
	dto["password"] = ""
	dto["domain"] = p.DomainURL
	dto["refresh_token"] = rToken

	result, _, statusCode := p.Client.Post(loginURL, dto, nil)
	if statusCode != 200 {
		return "", fmt.Errorf(string(result))
	}

	// From JWT To Frontend
	var newJWT map[string]interface{}
	json.Unmarshal(result, &newJWT)

	newAuthData := p.Converter.FromJWTToFrontend(newJWT)

	return newAuthData, nil
}
