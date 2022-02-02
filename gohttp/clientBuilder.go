package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	disableTimeouts    bool
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	client             *http.Client
	userAgent          string
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
	SetUserAgent(agent string) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) Build() Client {
	client := client{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(agent string) ClientBuilder {
	c.userAgent = agent
	return c
}
