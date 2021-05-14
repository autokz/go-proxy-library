package converter

import (
	"encoding/json"
)

type Converter struct {
}

func (c *Converter) FromFrontendToJWT(authData string) map[string]interface{} {
	m := make(map[string]interface{})
	json.Unmarshal([]byte(authData), &m)

	return m
}

func (c *Converter) FromJWTToFrontend(jwt map[string]interface{}) string {
	jsonAuthData, _ := json.Marshal(jwt)
	return string(jsonAuthData)
}
