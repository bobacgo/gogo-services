package uhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Get[T any](url string, query url.Values) (*T, error) {
	if query != nil {
		url += "?" + query.Encode()
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return ReadResponse[T](resp)
}

func Post[T any](url string, body any) (*T, error) {
	req, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, MIMEJSON, bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}
	return ReadResponse[T](resp)
}

func ReadResponse[T any](resp *http.Response) (*T, error) {
	if resp == nil {
		return nil, errors.New("nil response")
	}
	defer GracefulClose(resp)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(T)
	ct := resp.Header.Get(HeaderContentType)
	switch {
	case strings.Contains(ct, MIMEJSON):
		if err = json.Unmarshal(res, data); err != nil {
			return nil, err
		}
	case strings.Contains(ct, MIMEXML):
		if err = xml.Unmarshal(res, data); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported content type")
	}
	return data, nil
}

func GracefulClose(resp *http.Response) {
	if resp == nil || resp.Body == nil {
		return
	}
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
}

type HttpClient struct {
	Endpoint string
	Method   string
	Header   http.Header
	Query    map[string]string
	Body     []byte
	Timeout  time.Duration
}

func NewHttpClient(url string, method string, options ...HttpClientOption) *HttpClient {
	return &HttpClient{
		Endpoint: url,
		Method:   method,
		Header:   make(http.Header),
		Query:    make(map[string]string),
		Body:     []byte{},
		Timeout:  10 * time.Second,
	}
}

func (c *HttpClient) Do(ctx context.Context) any {
	//req, err := http.NewRequestWithContext(ctx, c.Method, c.Endpoint, nil)
	//http.Get()
	//
	//url.ParseQuery()
	return nil
}

type HttpClientOption func(c *HttpClient)

func WithTimeout(timeout time.Duration) HttpClientOption {
	return func(c *HttpClient) {
		c.Timeout = timeout
	}
}

func WithBody(data any) HttpClientOption {
	return func(c *HttpClient) {

	}
}
