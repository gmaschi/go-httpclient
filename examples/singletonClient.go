package examples

import (
	"github.com/gmaschi/go-httpclient/gohttp"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetResponseTimeout(2 * time.Second).
		SetConnectionTimeout(2 * time.Second).
		Build()

	return client
}
