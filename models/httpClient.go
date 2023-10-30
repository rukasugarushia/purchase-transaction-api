package models

import (
	"io"
	"net/http"
)

type HTTPClient struct {
	baseURL string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{baseURL: baseURL}
}

func (c *HTTPClient) MakeGETRequest(endpoint string) ([]byte, error) {
	fullURL := c.baseURL + "?" + endpoint

	response, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
