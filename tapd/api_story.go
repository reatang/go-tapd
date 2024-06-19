package tapd

type Story struct {
	c *Client
}

func newStory(c *Client) *Story {
	return &Story{c: c}
}
