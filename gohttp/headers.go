package gohttp

import "net/http"

const (
	userAgentHeaderKey = "User-Agent"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *client) getRequestHeaders(requestHeaders http.Header) http.Header {
	resultHeaders := make(http.Header)

	// Common Headers
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			resultHeaders.Set(header, value[0])
		}
	}

	// Custom Headers
	for header, value := range requestHeaders {
		if len(value) > 0 {
			resultHeaders.Set(header, value[0])
		}
	}

	if c.builder.userAgent != "" {
		if resultHeaders.Get(userAgentHeaderKey) != "" {
			return resultHeaders
		}
		resultHeaders.Set(userAgentHeaderKey, c.builder.userAgent)
	}

	return resultHeaders
}
