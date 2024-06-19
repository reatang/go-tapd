package tapd

type Bug struct {
	c *Client
}

func newBug(c *Client) *Bug {
	return &Bug{c: c}
}
