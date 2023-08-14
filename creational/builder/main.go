package main

import (
	"fmt"
	"net/http"
	"strings"
)

type RequestBuilder struct {
	Method string
	URL    string
	Header http.Header
	Body   string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		Header: make(http.Header),
	}
}

func (b *RequestBuilder) SetMethod(method string) *RequestBuilder {
	b.Method = strings.ToUpper(method)
	return b
}

func (b *RequestBuilder) SetURL(url string) *RequestBuilder {
	b.URL = url
	return b
}

func (b *RequestBuilder) SetHeader(key, value string) *RequestBuilder {
	b.Header.Add(key, value)
	return b
}

func (b *RequestBuilder) SetBody(body string) *RequestBuilder {
	b.Body = body
	return b
}

func (b *RequestBuilder) Build() (*http.Request, error) {
	req, err := http.NewRequest(b.Method, b.URL, strings.NewReader(b.Body))
	if err != nil {
		return nil, err
	}

	req.Header = b.Header

	return req, nil
}

func main() {
	builder := NewRequestBuilder().
		SetMethod("POST").
		SetURL("http://example.com").
		SetHeader("Content-Type", "application/json").
		SetBody(`{"key": "value"}`)

	req, err := builder.Build()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", req) // This prints the request object
}
