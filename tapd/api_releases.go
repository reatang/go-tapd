package tapd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
	"github.com/reatang/go-tapd/pkg/hx"
)

type Releases struct {
	c *Client
}

func newReleases(c *Client) *Releases {
	return &Releases{c}
}

type GetReleasesRequest struct {
	ID          int       `json:"id,omitempty" url:"id,omitempty"`                   // 支持多ID查询
	WorkspaceID int       `json:"workspace_id" url:"workspace_id"`                   // 项目ID
	Name        string    `json:"name,omitempty" url:"name,omitempty"`               // 标题，支持模糊匹配
	Description string    `json:"description,omitempty" url:"description,omitempty"` // 详细描述
	StartDate   time.Time `json:"startdate,omitempty" url:"startdate,omitempty"`     // 开始时间
	EndDate     time.Time `json:"enddate,omitempty" url:"enddate,omitempty"`         // 结束时间
	Creator     string    `json:"creator,omitempty" url:"creator,omitempty"`         // 创建人
	Created     time.Time `json:"created,omitempty" url:"created,omitempty"`         // 创建时间，支持时间查询
	Modified    time.Time `json:"modified,omitempty" url:"modified,omitempty"`       // 最后修改时间，支持时间查询
	Status      string    `json:"status,omitempty" url:"status,omitempty"`           // 状态，枚举值为 'done' 或 'open'
	Limit       int       `json:"limit,omitempty" url:"limit,omitempty"`             // 设置返回数量限制，默认为30
	Page        int       `json:"page,omitempty" url:"page,omitempty"`               // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order       string    `json:"order,omitempty" url:"order,omitempty"`             // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode
	Fields      string    `json:"fields,omitempty" url:"fields,omitempty"`           // 设置获取的字段，多个字段间以','逗号隔开
}

type Release struct {
	ID          string    `json:"id"`           // 发布计划的唯一标识符
	WorkspaceID string    `json:"workspace_id"` // 工作空间ID
	Name        string    `json:"name"`         // 发布计划的名称
	Description string    `json:"description"`  // 发布计划的描述
	StartDate   time.Time `json:"startdate"`    // 发布计划的开始日期
	EndDate     time.Time `json:"enddate"`      // 发布计划的结束日期
	Creator     string    `json:"creator"`      // 发布计划的创建者
	Created     time.Time `json:"created"`      // 发布计划的创建时间
	Modified    time.Time `json:"modified"`     // 发布计划的最后修改时间
	Status      string    `json:"status"`       // 发布计划的状态
}

type GetReleaseResponse struct {
	BaseResponse
	Data struct {
		Release Release `json:"release"`
	} `json:"data"`
}

// GetReleases
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_releases.html
func (r *Releases) GetReleases(ctx context.Context, req *GetReleasesRequest) (*GetReleaseResponse, error) {
	p, _ := query.Values(req)

	uri := fmt.Sprintf("/releases?%s", p.Encode())
	request := r.c.newRequest(ctx, http.MethodGet, uri)

	response, err := DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "[http] Releases/GetReleases error")
	}

	body, err := hx.ReadResponseBody(response)
	if err != nil {
		return nil, errors.Wrap(err, "[http] Releases/GetReleases read response error")
	}

	resp := &GetReleaseResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, errors.Wrap(err, "[TapdClient] Releases/GetReleases json error")
	}

	return resp, nil
}
