package mocks

import (
	"bytes"
	"io"
	"net/http"
)

type MockHTTPClient struct {
	MockResponse *http.Response
	MockError    error
}

func (m *MockHTTPClient) DoRequest(method, url string, headers map[string]string, body interface{}) (*http.Response, error) {
	if m.MockError != nil {
		return nil, m.MockError
	}
	return m.MockResponse, nil
}

func NewMockHTTPResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewReader([]byte(body))),
		Header:     make(http.Header),
	}
}
