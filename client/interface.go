package client

import "net/http"

type NetClientInterface interface {
	Get(URL string, header map[string]interface{}) (result []byte, headers http.Header, statusCode int)
	Post(URL string, body, header map[string]interface{}) (result []byte, headers http.Header, statusCode int)
}
