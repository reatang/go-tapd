package service

import (
	"context"
	"net/http"

	jsondt "github.com/reatang/go-json-datetime"
)

type Task struct {
	cmd Cmdable
}

func NewTask(cmd Cmdable) *Task {
	return &Task{cmd: cmd}
}

type GetTaskRequest struct {
	BaseRequest
	QueryFields

	ID              any     `json:"id,omitempty" url:"id,omitempty"`                             // id，支持多个ID查询及模糊匹配
	Name            string  `json:"name,omitempty" url:"name,omitempty"`                         // 任务标题，支持模糊匹配
	Description     string  `json:"description,omitempty" url:"description,omitempty"`           // 任务详细描述
	Creator         string  `json:"creator,omitempty" url:"creator,omitempty"`                   // 创建者，支持多个人员查询
	Status          string  `json:"status,omitempty" url:"status,omitempty"`                     // 状态，支持枚举查询
	Label           string  `json:"label,omitempty" url:"label,omitempty"`                       // 标签查询，支持枚举查询
	Owner           string  `json:"owner,omitempty" url:"owner,omitempty"`                       // 当前任务处理人，支持模糊匹配
	CC              string  `json:"cc,omitempty" url:"cc,omitempty"`                             // 抄送，抄送人
	Begin           string  `json:"begin,omitempty" url:"begin,omitempty"`                       // 预计开始，支持时间查询
	Due             string  `json:"due,omitempty" url:"due,omitempty"`                           // 预计结束，支持时间查询
	StoryID         any     `json:"story_id,omitempty" url:"story_id,omitempty"`                 // 相关需求ID，支持多个ID查询
	IterationID     int     `json:"iteration_id,omitempty" url:"iteration_id,omitempty"`         // 所属迭代ID，支持枚举查询
	Priority        string  `json:"priority,omitempty" url:"priority,omitempty"`                 // 优先级，若需要自定义优先级兼容，请使用priority_label
	PriorityLabel   string  `json:"priority_label,omitempty" url:"priority_label,omitempty"`     // 建议优先级字段
	Progress        int     `json:"progress,omitempty" url:"progress,omitempty"`                 // 进度
	Completed       string  `json:"completed,omitempty" url:"completed,omitempty"`               // 完成时间，支持时间查询
	EffortCompleted string  `json:"effort_completed,omitempty" url:"effort_completed,omitempty"` // 已完成工作量
	Exceed          float64 `json:"exceed,omitempty" url:"exceed,omitempty"`                     // 超出工作量
	Remain          float64 `json:"remain,omitempty" url:"remain,omitempty"`                     // 剩余工作量
	Effort          string  `json:"effort,omitempty" url:"effort,omitempty"`

	QueryTaskCustomFields
}

type QueryTaskCustomFields struct {
	CustomFieldOne   *string `json:"custom_field_one,omitempty" url:"custom_field_one,omitempty"`
	CustomFieldTwo   *string `json:"custom_field_two,omitempty" url:"custom_field_two,omitempty"`
	CustomFieldThree *string `json:"custom_field_three,omitempty" url:"custom_field_three,omitempty"`
	CustomFieldFour  *string `json:"custom_field_four,omitempty" url:"custom_field_four,omitempty"`
	CustomFieldFive  *string `json:"custom_field_five,omitempty" url:"custom_field_five,omitempty"`
	CustomFieldSix   *string `json:"custom_field_six,omitempty" url:"custom_field_six,omitempty"`
	CustomFieldSeven *string `json:"custom_field_seven,omitempty" url:"custom_field_seven,omitempty"`
	CustomFieldEight *string `json:"custom_field_eight,omitempty" url:"custom_field_eight,omitempty"`
	CustomField9     *string `json:"custom_field_9,omitempty" url:"custom_field_9,omitempty"`
	CustomField10    *string `json:"custom_field_10,omitempty" url:"custom_field_10,omitempty"`
}

type TaskData struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	WorkspaceID      string           `json:"workspace_id"`
	Creator          string           `json:"creator"`
	Created          jsondt.DateTime  `json:"created"`
	Modified         jsondt.DateTime  `json:"modified"`
	Status           string           `json:"status"`
	Owner            string           `json:"owner"`
	CC               string           `json:"cc"`
	Begin            *jsondt.Date     `json:"begin"`
	Due              *jsondt.Date     `json:"due"`
	StoryID          string           `json:"story_id"`
	IterationID      string           `json:"iteration_id"`
	Priority         string           `json:"priority"`
	Progress         string           `json:"progress"`
	Completed        *jsondt.DateTime `json:"completed"`
	EffortCompleted  string           `json:"effort_completed"`
	Exceed           string           `json:"exceed"`
	Remain           string           `json:"remain"`
	Effort           string           `json:"effort"`
	HasAttachment    string           `json:"has_attachment"`
	ReleaseID        string           `json:"release_id"`
	Label            *string          `json:"label"`
	CustomFieldOne   *string          `json:"custom_field_one"`
	CustomFieldTwo   *string          `json:"custom_field_two"`
	CustomFieldThree *string          `json:"custom_field_three"`
	CustomFieldFour  *string          `json:"custom_field_four"`
	CustomFieldFive  *string          `json:"custom_field_five"`
	CustomFieldSix   *string          `json:"custom_field_six"`
	CustomFieldSeven *string          `json:"custom_field_seven"`
	CustomFieldEight *string          `json:"custom_field_eight"`
	CustomField9     *string          `json:"custom_field_9"`
	CustomField10    *string          `json:"custom_field_10"`
}

type GetTaskResponse struct {
	BaseResponse
	Data []struct {
		Task TaskData `json:"task"`
	} `json:"data"`
}

// GetTasks
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks.html
func (r *Task) GetTasks(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error) {
	resp := &GetTaskResponse{}
	err := r.cmd(ctx, http.MethodGet, "tasks", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTasksCount
// Doc：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_releases_count.html
func (r *Task) GetTasksCount(ctx context.Context, req *GetTaskRequest) (*CountResponse, error) {
	resp := &CountResponse{}
	err := r.cmd(ctx, http.MethodGet, "tasks/count", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
