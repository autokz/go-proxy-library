package oauthProxy

import (
	"github.com/autokz/go-proxy-library/client"
	"github.com/autokz/go-proxy-library/converter"
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
	if p.Client == nil {
		p.Client = &client.HttpClient{
			Client: http.Client{Timeout: time.Duration(1) * time.Second},
		}
	}

	if p.Converter == nil {
		p.Converter = &converter.Converter{}
	}
}
