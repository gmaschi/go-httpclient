package examples

import (
	"errors"
	"github.com/gmaschi/go-httpclient/gohttpMock"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	// TODO: validation with mocks
	t.Run("FetchingError", func(t *testing.T) {
		url := "https://api.github.com"
		mock := gohttpMock.Mock{
			Method: http.MethodGet,
			Url:    url,
			Error:  errors.New("request error"),
		}
		endpoints, err := Get()
	})

	t.Run("UnmarshalError", func(t *testing.T) {
		endpoints, err := Get()
	})

	t.Run("NoError", func(t *testing.T) {
		endpoints, err := Get()
	})
}
