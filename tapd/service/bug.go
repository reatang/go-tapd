package service

type Bug struct {
	cmd Cmdable
}

func NewBug(cmd Cmdable) *Bug {
	return &Bug{cmd: cmd}
}
