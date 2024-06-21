package service

import "context"

type Cmdable func(ctx context.Context, method, uri string, req any, resp any) error

type BaseResponse struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

type CountResponse struct {
	BaseResponse
	Data struct {
		Count int `json:"count"`
	} `json:"data"`
}

type ISetWorkspaceID interface {
	GetWorkspaceID() int
	SetWorkspaceID(id int)
}

type BaseRequest struct {
	WorkspaceID int `json:"workspace_id" url:"workspace_id"` // 项目ID
}

func (r *BaseRequest) GetWorkspaceID() int {
	return r.WorkspaceID
}

func (r *BaseRequest) SetWorkspaceID(id int) {
	r.WorkspaceID = id
}

type QueryRequest struct {
	Limit  int    `json:"limit,omitempty" url:"limit,omitempty"`   // 设置返回数量限制，默认为30
	Page   int    `json:"page,omitempty" url:"page,omitempty"`     // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order  string `json:"order,omitempty" url:"order,omitempty"`   // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode
	Fields string `json:"fields,omitempty" url:"fields,omitempty"` // 设置获取的字段，多个字段间以','逗号隔开
}
