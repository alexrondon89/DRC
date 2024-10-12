package mock

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	ExpectedResponse   interface{}
	ExpectedStatusCode int
}

func NewHttpMock(expectedStatusCode int, expectedResponse interface{}) HttpClient {
	return HttpClient{
		ExpectedResponse:   expectedResponse,
		ExpectedStatusCode: expectedStatusCode,
	}
}

func (cli HttpClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: cli.ExpectedStatusCode,
		Body:       io.NopCloser(bytes.NewBufferString(fmt.Sprintf("%v", cli.ExpectedResponse))),
	}, nil
}
