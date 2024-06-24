package service

import (
	"context"
	"strings"
	"time"
)

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

type QueryFields struct {
	Created  string `json:"created,omitempty" url:"created,omitempty"`   // 创建时间，支持时间查询
	Modified string `json:"modified,omitempty" url:"modified,omitempty"` // 最后修改时间，支持时间查询

	Limit  int    `json:"limit,omitempty" url:"limit,omitempty"`   // 设置返回数量限制，默认为30
	Page   int    `json:"page,omitempty" url:"page,omitempty"`     // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order  string `json:"order,omitempty" url:"order,omitempty"`   // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode
	Fields string `json:"fields,omitempty" url:"fields,omitempty"` // 设置获取的字段，多个字段间以','逗号隔开
}

func (q *QueryFields) Select(field ...string) {
	q.Fields = strings.Join(field, ",")
}

func (q *QueryFields) WhereCreatedGt(t time.Time) {
	q.Created = ">" + t.Format("2006-01-02 15:04:05")
}

func (q *QueryFields) WhereCreatedLt(t time.Time) {
	q.Created = "<" + t.Format("2006-01-02 15:04:05")
}

func (q *QueryFields) WhereCreatedBetween(st time.Time, et time.Time) {
	q.Created = st.Format("2006-01-02 15:04:05") + "~" + et.Format("2006-01-02 15:04:05")
}

func (q *QueryFields) WhereModifiedGt(t time.Time) {
	q.Modified = ">" + t.Format("2006-01-02 15:04:05")
}

func (q *QueryFields) WhereModifiedLt(t time.Time) {
	q.Created = "<" + t.Format("2006-01-02 15:04:05")
}

func (q *QueryFields) WhereModifiedBetween(st time.Time, et time.Time) {
	q.Modified = st.Format("2006-01-02 15:04:05") + "~" + et.Format("2006-01-02 15:04:05")
}

func (q *QueryFields) OrderByCreatedAsc() {
	q.Order = "created asc"
}

func (q *QueryFields) OrderByCreatedDesc() {
	q.Order = "created desc"
}

func (q *QueryFields) OrderByModifiedAsc() {
	q.Order = "modified asc"
}

func (q *QueryFields) OrderByModifiedDesc() {
	q.Order = "modified desc"
}
