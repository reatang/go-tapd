package tapd

type Task struct {
	c *Client
}

func newTask(c *Client) *Task {
	return &Task{c: c}
}
