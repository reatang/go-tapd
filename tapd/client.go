package tapd

import (
	"context"
	"io"
	"net/http"
)

const DefaultEndpoint = "https://api.tapd.cn"

var DefaultClient *http.Client

func init() {
	DefaultClient = http.DefaultClient
}

type Client struct {
	opts Options
}

func NewTAPDClient(opts ...Option) *Client {
	c := &Client{opts: Options{
		endpoint: DefaultEndpoint,
	}}

	for _, opt := range opts {
		opt(&c.opts)
	}

	return c
}

func (c *Client) newRequest(ctx context.Context, method string, uri string) *http.Request {
	return c.newRequestWithBody(ctx, method, uri, nil)
}

func (c *Client) newRequestWithBody(ctx context.Context, method string, uri string, body io.ReadCloser) *http.Request {
	req, _ := http.NewRequestWithContext(ctx, method, c.opts.endpoint+uri, body)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("User-Agent", "go-tapd/"+Version())
	req.SetBasicAuth(c.opts.username, c.opts.password)

	return req
}

func (c *Client) Story() *Story {
	return newStory(c)
}

func (c *Client) Releases() *Releases {
	return newReleases(c)
}

func (c *Client) Iteration() *Iteration {
	return newIteration(c)
}

func (c *Client) Task() *Task {
	return newTask(c)
}

func (c *Client) Bug() *Bug {
	return newBug(c)
}
