package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
	Client http.Client
}

func (c *HttpClient) Get(URL string, requestHeader map[string]interface{}) ([]byte, http.Header, int) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Panicln(err)
		return nil, nil, 0
	}

	addHeaders(req, requestHeader)

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Panicln(err)
		return nil, nil, 0
	}
	defer resp.Body.Close()

	// Prepare Result
	responseBody, _ := ioutil.ReadAll(resp.Body)

	return responseBody, resp.Header, resp.StatusCode
}

func (c *HttpClient) Post(URL string, requestBody, requestHeader map[string]interface{}) ([]byte, http.Header, int) {

	bodyJson, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Panicln(err)
		return nil, nil, 0
	}

	addHeaders(req, requestHeader)

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Panicln(err)
		return nil, nil, 0
	}
	defer resp.Body.Close()

	// Prepare Result
	responseBody, _ := ioutil.ReadAll(resp.Body)

	return responseBody, resp.Header, resp.StatusCode
}

func addHeaders(req *http.Request, header map[string]interface{}) {
	req.Header.Add("Accept", `application/json`)

	for k, v := range header {
		req.Header.Add(k, v.(string))
	}
}
