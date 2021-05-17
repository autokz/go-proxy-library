package oauthProxy

import (
	"github.com/autokz/go-proxy-library/client"
	"github.com/autokz/go-proxy-library/converter"
)

type Proxy struct {
	OAuthGrantType        string
	OAuthRefreshGrantType string

	Client    client.OAuthClientInterface
	Converter converter.JwtConverterInterface
}