package service

type Task struct {
	cmd Cmdable
}

func NewTask(cmd Cmdable) *Task {
	return &Task{cmd: cmd}
}
