package tapd

type Iteration struct {
	c *Client
}

func newIteration(c *Client) *Iteration {
	return &Iteration{c: c}
}
