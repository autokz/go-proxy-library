package oauthProxy

import (
	"fmt"
)

func (p *Proxy) Check(authData string) error {
	jwt := p.Converter.FromFrontendToJWT(authData)

	var aToken string
	if aToken = ""; jwt["access_token"] != nil {
		aToken = jwt["access_token"].(string)
	}

	checkURL := p.BaseURL + p.CheckURL
	headers := map[string]interface{}{
		"Authorization": fmt.Sprintf("Bearer %s", aToken),
	}

	result, _, statusCode := p.Client.Get(checkURL, headers)
	if statusCode != 200 {
		return fmt.Errorf(string(result))
	}

	return nil
}
