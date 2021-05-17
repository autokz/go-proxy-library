package oauthProxy

import (
	"github.com/autokz/go-proxy-library/client"
)

func (p *Proxy) LogoutAll(authData string)  error {
	jwt := p.Converter.FromFrontendToJWT(authData)

	var accessToken client.AccessToken
	if accessToken = ""; jwt["access_token"] != nil {
		accessToken = jwt["access_token"].(client.AccessToken)
	}

	err := p.Client.LogoutAll(accessToken)
	if err != nil {
		return err
	}

	return nil
}
