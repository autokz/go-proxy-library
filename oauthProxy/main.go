package oauthProxy

import (
	"gateway/proxy/client"
	"gateway/proxy/converter"
	"net/http"
	"time"
)

type Proxy struct {
	BaseURL      string
	CheckURL     string
	OAuthURL     string
	LogoutURL    string
	LogoutAllURL string
	DomainURL    string

	OAuthGrantType        string
	OAuthRefreshGrantType string

	Client    client.NetClientInterface
	Converter converter.JwtConverterInterface
}

func (p *Proxy) Init() {
	p.Client = &client.HttpClient{
		Client: http.Client{Timeout: time.Duration(1) * time.Second},
	}

	p.Converter = &converter.Converter{}
}
