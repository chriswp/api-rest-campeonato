package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	DoRequest(method, url string, headers map[string]string, body interface{}) (*http.Response, error)
}

type httpClient struct {
	client *http.Client
}

func NewHTTPClient(timeout time.Duration) HTTPClient {
	return &httpClient{
		client: &http.Client{Timeout: timeout},
	}
}

func (h *httpClient) DoRequest(method, url string, headers map[string]string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return h.client.Do(req)
}
