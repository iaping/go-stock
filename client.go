package stock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type Request func(*fasthttp.Request) error
type Response func(*fasthttp.Response) error

type Client struct {
	client *fasthttp.Client
}

func New(client *fasthttp.Client) *Client {
	return &Client{client: client}
}

func Default() *Client {
	c := &fasthttp.Client{
		ReadTimeout:              5 * time.Second,
		WriteTimeout:             5 * time.Second,
		NoDefaultUserAgentHeader: true,
	}
	return New(c)
}

func (c *Client) Json(reqopt Request, respopt Response, v interface{}) error {
	data, err := c.Do(reqopt, respopt)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (c *Client) Do(reqopt Request, respopt Response) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	var err error
	if err = reqopt(req); err != nil {
		return nil, err
	}

	if err = c.client.Do(req, resp); err != nil {
		return nil, err
	}
	if code := resp.StatusCode(); code != fasthttp.StatusOK {
		return nil, fmt.Errorf("http code: %d", code)
	}
	if respopt != nil {
		if err = respopt(resp); err != nil {
			return nil, err
		}
	}

	return bytes.Clone(resp.Body()), nil
}
