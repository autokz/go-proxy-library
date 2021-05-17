package oauthProxy

func (p *Proxy) Refresh(authData string) (string, error) {
	dto := p.Converter.FromFrontendToJWT(authData)

	// Getting new JWT with User Payload
	newAuthData, _, err := p.Client.Auth(dto, p.OAuthRefreshGrantType)
	if err != nil {
		return "",  err
	}

	// Converting JWT...
	cryptedAuthData := p.Converter.FromJWTToFrontend(newAuthData)

	return cryptedAuthData, nil
}
