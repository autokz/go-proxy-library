package oauthProxy

import (
	"fmt"
)

func (p *Proxy) Logout(authData string) (bool, error) {
	jwt := p.Converter.FromFrontendToJWT(authData)

	var aToken string
	if aToken = ""; jwt["access_token"] != nil {
		aToken = jwt["access_token"].(string)
	}

	logoutURL := p.BaseURL + p.LogoutURL

	headers := map[string]interface{}{
		"Authorization": fmt.Sprintf("Bearer %s", aToken),
	}

	result, _, statusCode := p.Client.Post(logoutURL, nil, headers)
	if statusCode != 200 {
		return false, fmt.Errorf(string(result))
	}

	return true, nil
}
