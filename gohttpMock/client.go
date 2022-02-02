package gohttpMock

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct {
}

func (cm *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	reqBody, err := req.GetBody()
	if err != nil {
		return nil, err
	}
	defer reqBody.Close()

	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	var response http.Response
	mock := MockupServer.mocks[MockupServer.getMockKey(req.Method, req.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		response.StatusCode = mock.ResponseStatusCode
		response.Body = ioutil.NopCloser(strings.NewReader(mock.ResponseBody))
		response.ContentLength = int64(len(mock.ResponseBody))
		response.Request = req
		return &response, nil
	}

	return nil, errors.New(fmt.Sprintf("no mock matching %s from %s with given body", req.Method, req.URL.String()))
}
