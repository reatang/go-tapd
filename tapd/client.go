package tapd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/reatang/go-tapd/pkg/hx"
	"github.com/reatang/go-tapd/tapd/service"
)

const DefaultEndpoint = "https://api.tapd.cn"

var DefaultClient *http.Client

func init() {
	DefaultClient = http.DefaultClient
}

type Client struct {
	opts Options

	Story     *service.Story
	Releases  *service.Releases
	Iteration *service.Iteration
	Task      *service.Task
	Bug       *service.Bug
}

func NewTAPDClient(opts ...Option) *Client {
	c := &Client{opts: Options{
		endpoint: DefaultEndpoint,
	}}

	for _, opt := range opts {
		opt(&c.opts)
	}

	// init
	c.Story = service.NewStory(c.Do)
	c.Releases = service.NewReleases(c.Do)
	c.Iteration = service.NewIteration(c.Do)
	c.Task = service.NewTask(c.Do)
	c.Bug = service.NewBug(c.Do)

	return c
}

func (c *Client) newRequest(ctx context.Context, method string, uri string) *http.Request {
	return c.newRequestWithBody(ctx, method, uri, nil)
}

func (c *Client) newRequestWithBody(ctx context.Context, method string, uri string, body io.ReadCloser) *http.Request {
	req, _ := http.NewRequestWithContext(ctx, method, c.opts.endpoint+uri, body)

	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "go-tapd/"+Version())
	req.SetBasicAuth(c.opts.username, c.opts.password)

	return req
}

// Do 如果有未实现，或者希望有自己调用检查，则可以使用调用该接口发起请求
func (c *Client) Do(ctx context.Context, method, uri string, req any, resp any) error {
	if _req, ok := req.(service.ISetWorkspaceID); ok && _req.GetWorkspaceID() == 0 {
		_req.SetWorkspaceID(c.opts.workspaceID)
	}

	var u string
	var request *http.Request
	if method == http.MethodGet {
		p, _ := query.Values(req)
		u = fmt.Sprintf("/%s?%s", uri, p.Encode())
		request = c.newRequest(ctx, method, u)
	} else if method == http.MethodPost {
		u = fmt.Sprintf("%s", uri)
		j, err := json.Marshal(req)
		if err != nil {
			return err
		}
		request, _ = http.NewRequestWithContext(ctx, method, u, bytes.NewBuffer(j))
		request.Header.Set("Content-Type", "application/json")
	} else {
		return errors.Errorf("unsupported method: %s", method)
	}

	response, err := DefaultClient.Do(request)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("[http] %s error", u))
	}

	body, err := hx.ReadResponseBody(response)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("[http] %s read response error", u))
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("[TapdClient] %s json error", u))
	}

	return nil
}
