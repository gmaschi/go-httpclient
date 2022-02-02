package gohttp

import (
	"net/http"
	"testing"
)

const (
	headerContentTypeKey   = "Content-Type"
	headerContentTypeValue = "application/json"
	headerUserAgentKey     = "User-Agent"
	headerUserAgentValue   = "cool-http-client"
	headerXRequestIdKey    = "X-Request-Id"
	headerXRequestIdValue  = "cool-http-client"
)

func TestGetRequestHeaders(t *testing.T) {
	// initialization
	commonHeaders := make(http.Header)
	commonHeaders.Set(headerContentTypeKey, headerContentTypeValue)
	commonHeaders.Set(headerUserAgentKey, headerUserAgentValue)
	client := client{}

	// execution
	requestHeaders := make(http.Header)
	requestHeaders.Set(headerXRequestIdKey, headerXRequestIdValue)

	finalHeaders := client.getRequestHeaders(requestHeaders)
	expectedHeaders := 3
	// TODO: validation with mocks
	if len(finalHeaders) != expectedHeaders {
		t.Errorf("length of final Headers expected to be %v, and got %v", expectedHeaders, len(finalHeaders))
	}

	if finalHeaders.Get(headerContentTypeKey) != headerContentTypeValue {
		t.Errorf(
			"invalid header value for key %v, expected %v, got %v",
			headerContentTypeKey,
			headerContentTypeValue,
			finalHeaders.Get(headerContentTypeKey),
		)
	}

	if finalHeaders.Get(headerUserAgentKey) != headerUserAgentValue {
		t.Errorf(
			"invalid header value for key %v, expected %v, got %v",
			headerUserAgentKey,
			headerUserAgentValue,
			finalHeaders.Get(headerUserAgentKey),
		)
	}

	if finalHeaders.Get(headerXRequestIdKey) != headerXRequestIdValue {
		t.Errorf(
			"invalid header value for key %v, expected %v, got %v",
			headerXRequestIdKey,
			headerXRequestIdValue,
			finalHeaders.Get(headerXRequestIdKey),
		)
	}
}

func TestGetRequestBody(t *testing.T) {
	client := client{}

	t.Run("nilBody", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Errorf("error while passing nil Body, expected no error")
		}

		if body != nil {
			t.Errorf("expected Body to be %v, instead got: %v", nil, body)
		}
	})

	t.Run("jsonBody", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)
		if err != nil {
			t.Errorf("error while marshaling to json, err: %v", err)
		}

		expectedResponse := `["one","two"]`
		if string(body) != expectedResponse {
			t.Errorf("unexpected Body. expected: %v, got: %v", expectedResponse, string(body))
		}
	})

}
