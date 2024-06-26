package service

import (
	"context"
	"net/http"

	jsondt "github.com/reatang/go-json-datetime"
)

type Story struct {
	cmd Cmdable
}

func NewStory(cmd Cmdable) *Story {
	return &Story{cmd: cmd}
}

type GetStoriesRequest struct {
	BaseRequest
	QueryFields

	ID              any     `json:"id,omitempty" url:"id,omitempty"`
	Name            string  `json:"name,omitempty" url:"name,omitempty"`
	Priority        string  `json:"priority,omitempty" url:"priority,omitempty"`
	BusinessValue   any     `json:"business_value,omitempty" url:"business_value,omitempty"`
	Status          string  `json:"status,omitempty" url:"status,omitempty"`
	VStatus         string  `json:"v_status,omitempty" url:"v_status,omitempty"`
	WithVStatus     string  `json:"with_v_status,omitempty" url:"with_v_status,omitempty"`
	Label           string  `json:"label,omitempty" url:"label,omitempty"`
	WorkItemTypeID  any     `json:"workitem_type_id,omitempty" url:"workitem_type_id,omitempty"`
	Version         string  `json:"version,omitempty" url:"version,omitempty"`
	Module          string  `json:"module,omitempty" url:"module,omitempty"`
	Feature         string  `json:"feature,omitempty" url:"feature,omitempty"`
	TestFocus       string  `json:"test_focus,omitempty" url:"test_focus,omitempty"`
	Size            any     `json:"size,omitempty" url:"size,omitempty"`
	Owner           string  `json:"owner,omitempty" url:"owner,omitempty"`
	CC              string  `json:"cc,omitempty" url:"cc,omitempty"`
	Creator         string  `json:"creator,omitempty" url:"creator,omitempty"`
	Developer       string  `json:"developer,omitempty" url:"developer,omitempty"`
	Begin           string  `json:"begin,omitempty" url:"begin,omitempty"`
	Due             string  `json:"due,omitempty" url:"due,omitempty"`
	Completed       string  `json:"completed,omitempty" url:"completed,omitempty"`
	IterationID     any     `json:"iteration_id,omitempty" url:"iteration_id,omitempty"`
	Effort          string  `json:"effort,omitempty" url:"effort,omitempty"`
	EffortCompleted string  `json:"effort_completed,omitempty" url:"effort_completed,omitempty"`
	Remain          float64 `json:"remain,omitempty" url:"remain,omitempty"`
	Exceed          float64 `json:"exceed,omitempty" url:"exceed,omitempty"`
	CategoryID      any     `json:"category_id,omitempty" url:"category_id,omitempty"`
	ReleaseID       any     `json:"release_id,omitempty" url:"release_id,omitempty"`
	Source          string  `json:"source,omitempty" url:"source,omitempty"`
	Type            string  `json:"type,omitempty" url:"type,omitempty"`
	ParentID        any     `json:"parent_id,omitempty" url:"parent_id,omitempty"`
	ChildrenID      any     `json:"children_id,omitempty" url:"children_id,omitempty"`
	Description     string  `json:"description,omitempty" url:"description,omitempty"`
}

type StoryData struct {
	ID              string           `json:"id,omitempty"`               // uint64
	WorkItemTypeID  string           `json:"workitem_type_id,omitempty"` // uint64
	Name            string           `json:"name,omitempty"`
	Description     string           `json:"description,omitempty"`
	WorkspaceID     string           `json:"workspace_id,omitempty"` // uint64
	Creator         string           `json:"creator,omitempty"`
	Created         *jsondt.DateTime `json:"created,omitempty"`
	Modified        *jsondt.DateTime `json:"modified,omitempty"`
	Status          string           `json:"status,omitempty"`
	Owner           string           `json:"owner,omitempty"`
	CC              string           `json:"cc,omitempty"`
	Begin           *jsondt.Date     `json:"begin,omitempty"`
	Due             *jsondt.Date     `json:"due,omitempty"`
	Size            string           `json:"size,omitempty"`
	Priority        string           `json:"priority,omitempty"`
	Developer       string           `json:"developer,omitempty"`
	IterationID     string           `json:"iteration_id,omitempty"` // uint64 or ""
	TestFocus       string           `json:"test_focus,omitempty"`
	Type            string           `json:"type,omitempty"`
	Source          string           `json:"source,omitempty"`
	Module          string           `json:"module,omitempty"`
	Version         string           `json:"version,omitempty"`
	Completed       *jsondt.DateTime `json:"completed,omitempty"`
	CategoryID      string           `json:"category_id,omitempty"` // uint64 or ""
	Path            string           `json:"path,omitempty"`        // id::id:
	ParentID        string           `json:"parent_id,omitempty"`   // uint64 or == AncestorID
	ChildrenID      string           `json:"children_id,omitempty"` // ||id|id
	AncestorID      string           `json:"ancestor_id,omitempty"` // uint64
	BusinessValue   string           `json:"business_value,omitempty"`
	Effort          string           `json:"effort,omitempty"`
	EffortCompleted string           `json:"effort_completed,omitempty"`
	Exceed          string           `json:"exceed,omitempty"`
	Remain          string           `json:"remain,omitempty"`
	ReleaseID       string           `json:"release_id,omitempty"`
	Label           string           `json:"label,omitempty"`
}

type GetStoriesResponse struct {
	BaseResponse
	Data []struct {
		Story StoryData `json:"story"`
	} `json:"data"`
}

// GetStories
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories.html
func (r *Story) GetStories(ctx context.Context, req *GetStoriesRequest) (*GetStoriesResponse, error) {
	resp := &GetStoriesResponse{}
	err := r.cmd(ctx, http.MethodGet, "stories", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetStoriesCount
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories_count.html
func (r *Story) GetStoriesCount(ctx context.Context, req *GetStoriesRequest) (*CountResponse, error) {
	resp := &CountResponse{}
	err := r.cmd(ctx, http.MethodGet, "stories/count", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
