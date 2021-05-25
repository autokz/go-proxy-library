package oauthProxy

type UsernameType string
type PasswordType string
type DomainType string

func (p *Proxy) Login(username UsernameType, password PasswordType, domain DomainType) (string, map[string]interface{}, error) {
	dto := make(map[string]interface{})

	dto["username"] = username
	dto["password"] = password
	dto["domain"] = domain

	// Getting new JWT with User Payload
	newAuthData, userPayload, err := p.Client.Auth(dto, p.OAuthGrantType)
	if err != nil {
		return "", nil, err
	}

	// Converting JWT...
	cryptedAuthData := p.Converter.FromJWTToFrontend(newAuthData)


	return cryptedAuthData, userPayload, nil
}
