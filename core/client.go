package core

import "net/http"

type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
}
