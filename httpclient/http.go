package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	// 当接口请求成功
	OnApiSuccess func() error
	// 当接口请求失败
	OnApiError func(errDesc string) error
	// 当接口请求完成后
	OnResponse func(method, url string, in, out interface{})
)

type Client struct {
	// 忽略调用OnApiSuccess、OnApiError等消息
	ignoreCallback bool

	client *http.Client
}

func NewClient() *Client {
	return newClient("")
}

func NewClientWithBearer(bearerToken string) *Client {
	return newClient(bearerToken)
}

func newClient(bearerToken string) *Client {
	var trans = http.DefaultTransport

	if len(bearerToken) > 0 {
		trans = &BearerToken{
			Base:  http.DefaultTransport,
			Token: bearerToken,
		}
	}

	client := &http.Client{
		Transport: trans,
		Timeout:   10 * time.Second,
	}

	return &Client{
		ignoreCallback: false,
		client:         client,
	}
}

func (c *Client) IgnoreOnCallback() {
	c.ignoreCallback = true
}

func (c *Client) Post(path string, body, out interface{}) (*http.Response, error) {
	return c.do("POST", path, body, out)
}

func (c *Client) Get(path string, out interface{}) (*http.Response, error) {
	return c.do("GET", path, nil, out)
}

func (c *Client) do(method, path string, in, out interface{}) (resp *http.Response, err error) {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	fullURL := path

	req, _ := http.NewRequest(method, fullURL, nil)
	if in != nil {
		buf := new(bytes.Buffer)
		_ = json.NewEncoder(buf).Encode(in)
		req.Header.Set("Content-Type", "application/json")
		req.Body = ioutil.NopCloser(buf)
	}

	resp, err = c.client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		_ = resp.Body.Close()
		c.onDone(method, fullURL, in, out)
	}()

	if out != nil {
		if w, ok := out.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
			return resp, nil
		}
		return resp, json.NewDecoder(resp.Body).Decode(out)
	}

	return resp, nil
}

func (c *Client) onDone(method, fullURL string, in, out interface{}) {
	c.onResponse(method, fullURL, in, out)

	if c.ignoreCallback {
		log.Printf("ignore, lark request %s\n", fullURL)
		return
	}
	c.onApiError(out)
	c.onApiSuccess(out)
}

func (c *Client) onResponse(method, url string, in, out interface{}) {
	if OnResponse == nil {
		return
	}
	OnResponse(method, url, in, out)
}

func (c *Client) onApiError(out interface{}) {
	if OnApiError == nil {
		return
	}
	// TODO 处理错误请求
	_ = OnApiError("")
}

func (c *Client) onApiSuccess(out interface{}) {
	if OnApiSuccess == nil {
		return
	}
	// TODO 处理成功请求
	_ = OnApiSuccess()
}
