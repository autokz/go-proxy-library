package oauthProxy

import (
	"github.com/autokz/go-proxy-library/client"
)

func (p *Proxy) Check(authData string) error {
	jwt := p.Converter.FromFrontendToJWT(authData)

	var accessToken string
	if accessToken = ""; jwt["access_token"] != nil {
		accessToken = jwt["access_token"].(string)
	}

	err := p.Client.Check(client.AccessToken(accessToken))
	if err != nil {
		return err
	}

	return nil
}
