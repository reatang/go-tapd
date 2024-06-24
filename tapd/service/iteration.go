package service

import (
	"context"
	"net/http"

	jsondt "github.com/reatang/go-json-datetime"
)

type Iteration struct {
	cmd Cmdable
}

func NewIteration(cmd Cmdable) *Iteration {
	return &Iteration{cmd: cmd}
}

// GetIterationsRequest 结构体定义了获取迭代信息的请求参数
type GetIterationsRequest struct {
	BaseRequest
	QueryFields

	ID             any    `json:"id,omitempty" url:"id,omitempty"`                             // 支持多ID查询
	Name           string `json:"name,omitempty" url:"name,omitempty"`                         // 标题，支持模糊匹配
	Description    string `json:"description,omitempty" url:"description,omitempty"`           // 详细描述
	StartDate      string `json:"startdate,omitempty" url:"startdate,omitempty"`               // 开始时间，支持时间查询
	EndDate        string `json:"enddate,omitempty" url:"enddate,omitempty"`                   // 结束时间，支持时间查询
	WorkItemTypeID int    `json:"workitem_type_id,omitempty" url:"workitem_type_id,omitempty"` // 迭代类别
	PlanAppID      int    `json:"plan_app_id,omitempty" url:"plan_app_id,omitempty"`           // 计划应用ID
	Status         string `json:"status,omitempty" url:"status,omitempty"`                     // 状态（系统状态open/done，自定义状态可传中文）
	Creator        string `json:"creator,omitempty" url:"creator,omitempty"`                   // 创建人
	Completed      string `json:"completed,omitempty" url:"completed,omitempty"`               // 完成时间
}

type IterationData struct {
	ID          string           `json:"id,omitempty"`           // 迭代ID
	Name        string           `json:"name,omitempty"`         // 迭代名称
	WorkspaceID string           `json:"workspace_id,omitempty"` // 项目ID
	StartDate   string           `json:"startdate,omitempty"`    // 开始时间
	EndDate     string           `json:"enddate,omitempty"`      // 结束时间
	Status      string           `json:"status,omitempty"`       // 状态
	ReleaseID   string           `json:"release_id,omitempty"`   // 发布ID，可能为空
	Description string           `json:"description,omitempty"`  // 描述
	Creator     string           `json:"creator,omitempty"`      // 创建人
	Created     *jsondt.DateTime `json:"created,omitempty"`      // 创建时间
	Modified    *jsondt.DateTime `json:"modified,omitempty"`     // 修改时间
	Completed   *jsondt.DateTime `json:"completed,omitempty"`    // 完成时间，可能为空
}

type GetIterationsResponse struct {
	BaseResponse
	Data []struct {
		Iteration IterationData `json:"iteration"`
	} `json:"data"`
}

// GetIterations 调用示例，用于获取迭代信息
// doc: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations.html
func (r *Iteration) GetIterations(ctx context.Context, req *GetIterationsRequest) (*GetIterationsResponse, error) {
	resp := &GetIterationsResponse{}
	err := r.cmd(ctx, http.MethodGet, "iterations", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetIterationsCount
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations_count.html
func (r *Iteration) GetIterationsCount(ctx context.Context, req *GetIterationsRequest) (*CountResponse, error) {
	resp := &CountResponse{}
	err := r.cmd(ctx, http.MethodGet, "iterations/count", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
