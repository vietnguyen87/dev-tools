package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type HTTPClient struct {
	client  http.Client
	timeout time.Duration
	url     string
}

type HTTPRequest struct {
	Method       string
	URI          string
	HeaderParams url.Values
	BodyData     []byte
}

func (h *HTTPClient) SendHTTPContextRequest(ctx context.Context, method, uri string, params url.Values, body ...byte) ([]byte, error) {
	var err error
	var fullUrl *url.URL
	var req *http.Request
	var res []byte
	fullUrl, err = url.Parse(h.url)
	if err != nil {
		return nil, err
	}
	fullUrl.Path = path.Join(fullUrl.Path, uri)

	fullUrl.RawQuery = params.Encode()
	fullPath := fullUrl.String()

	req, err = http.NewRequestWithContext(ctx, method, fullPath, bytes.NewReader(body))
	if err != nil {
		err = fmt.Errorf("%v %v", fullUrl.String(), err)
		return nil, err
	}

	res, err = h.sendHTTPRequest(req)
	if err != nil {
		err = fmt.Errorf("%v %v", fullUrl.String(), err)
	}
	return res, err
}

func (h *HTTPClient) SendHTTPRequest(ctx context.Context, httpRequest HTTPRequest) ([]byte, error) {
	var err error
	var req *http.Request
	var res []byte

	fullPath := h.url + httpRequest.URI

	req, err = http.NewRequestWithContext(ctx, httpRequest.Method, fullPath, bytes.NewReader(httpRequest.BodyData))
	if err != nil {
		err = fmt.Errorf("%v %v", fullPath, err)
		return nil, err
	}

	SetHeader(req, httpRequest.HeaderParams)
	res, err = h.sendHTTPRequest(req)
	if err != nil {
		err = fmt.Errorf("%v %v", fullPath, err)
	}

	return res, err
}

func (h *HTTPClient) sendHTTPRequest(req *http.Request) ([]byte, error) {

	resp, err := h.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func NewHTTPClient() *HTTPClient {
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 3 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   4 * time.Second,
		ResponseHeaderTimeout: 6 * time.Second,
		ExpectContinueTimeout: 4 * time.Second,
		DisableKeepAlives:     false,
		MaxIdleConnsPerHost:   1024,
		MaxConnsPerHost:       2048,
	}
	return &HTTPClient{
		client: http.Client{
			Transport: trans,
			Timeout:   6 * time.Second,
		},
	}
}

func (h *HTTPClient) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

func (h *HTTPClient) GetTimeout() time.Duration {
	if h.timeout == 0 {
		return 3 * time.Second
	}
	return h.timeout
}

func (h *HTTPClient) IsDisabled() bool {
	return false
}

func (h *HTTPClient) SetURL(u string) *HTTPClient {
	h.url = u
	return h
}

func (h *HTTPClient) GetURL() string {
	return h.url
}

func SetHeader(req *http.Request, header url.Values) {
	for k, v := range header {
		value := strings.Join(v, ",")
		req.Header.Set(k, value)
	}
}
