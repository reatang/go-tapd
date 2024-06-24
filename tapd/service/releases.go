package service

import (
	"context"
	"net/http"

	jsondt "github.com/reatang/go-json-datetime"
)

type Releases struct {
	cmd Cmdable
}

func NewReleases(cmd Cmdable) *Releases {
	return &Releases{cmd: cmd}
}

type GetReleasesRequest struct {
	BaseRequest
	QueryFields

	ID          any    `json:"id,omitempty" url:"id,omitempty"`                   // 支持多ID查询
	Name        string `json:"name,omitempty" url:"name,omitempty"`               // 标题，支持模糊匹配
	Description string `json:"description,omitempty" url:"description,omitempty"` // 详细描述
	StartDate   string `json:"startdate,omitempty" url:"startdate,omitempty"`     // 开始时间
	EndDate     string `json:"enddate,omitempty" url:"enddate,omitempty"`         // 结束时间
	Creator     string `json:"creator,omitempty" url:"creator,omitempty"`         // 创建人
	Status      string `json:"status,omitempty" url:"status,omitempty"`           // 状态，枚举值为 'done' 或 'open'
}

type ReleaseData struct {
	ID          string           `json:"id,omitempty"`           // 发布计划的唯一标识符
	WorkspaceID string           `json:"workspace_id,omitempty"` // 工作空间ID
	Name        string           `json:"name,omitempty"`         // 发布计划的名称
	Description string           `json:"description,omitempty"`  // 发布计划的描述
	StartDate   *jsondt.Date     `json:"startdate,omitempty"`    // 发布计划的开始日期
	EndDate     *jsondt.Date     `json:"enddate,omitempty"`      // 发布计划的结束日期
	Creator     string           `json:"creator,omitempty"`      // 发布计划的创建者
	Created     *jsondt.DateTime `json:"created,omitempty"`      // 发布计划的创建时间
	Modified    *jsondt.DateTime `json:"modified,omitempty"`     // 发布计划的最后修改时间
	Status      string           `json:"status,omitempty"`       // 发布计划的状态
}

type GetReleaseResponse struct {
	BaseResponse
	Data []struct {
		Release ReleaseData `json:"release"`
	} `json:"data"`
}

// GetReleases
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_releases.html
func (r *Releases) GetReleases(ctx context.Context, req *GetReleasesRequest) (*GetReleaseResponse, error) {
	resp := &GetReleaseResponse{}
	err := r.cmd(ctx, http.MethodGet, "releases", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetReleasesCount
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_releases_count.html
func (r *Releases) GetReleasesCount(ctx context.Context, req *GetReleasesRequest) (*CountResponse, error) {
	resp := &CountResponse{}
	err := r.cmd(ctx, http.MethodGet, "releases/count", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
