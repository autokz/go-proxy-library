package oauthProxy

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
	cryptedAuthData := p.Converter.FromJWTToFrontend(newAuthData)


	return cryptedAuthData, userPayload, nil
}
