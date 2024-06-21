package service

type Iteration struct {
	cmd Cmdable
}

func NewIteration(cmd Cmdable) *Iteration {
	return &Iteration{cmd: cmd}
}
