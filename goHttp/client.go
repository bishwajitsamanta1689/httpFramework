package goHttp

import (
	"net/http"
)

type httpClient struct {
}

type ClientInt interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func New() ClientInt {
	client := &httpClient{}
	return client
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, body)
}
func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, body)
}
func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, "https://api.github.com", headers, nil)
}
