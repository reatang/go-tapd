package hx

import "net/http"

type LimitTransport struct {
	limit int
}

func (l *LimitTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

// NewLimitTransport 构造一个每分钟限制访问次数的Transport，应对
// TODO 未实现
func NewLimitTransport(limit int) http.RoundTripper {
	return nil
}
