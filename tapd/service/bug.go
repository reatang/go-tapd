package service

import (
	"context"
	"net/http"
)

type Bug struct {
	cmd Cmdable
}

func NewBug(cmd Cmdable) *Bug {
	return &Bug{cmd: cmd}
}

type GetBugsRequest struct {
	BaseRequest
	QueryFields

	ID             []int  `json:"id,omitempty" url:"id,omitempty"`                             // ID
	Title          string `json:"title,omitempty" url:"title,omitempty"`                       // 标题
	Priority       string `json:"priority,omitempty" url:"priority,omitempty"`                 // 优先级
	PriorityLabel  string `json:"priority_label,omitempty" url:"priority_label,omitempty"`     // 优先级标签
	Severity       string `json:"severity,omitempty" url:"severity,omitempty"`                 // 严重程度
	Status         string `json:"status,omitempty" url:"status,omitempty"`                     // 状态
	VStatus        string `json:"v_status,omitempty" url:"v_status,omitempty"`                 // 状态(中文)
	Label          string `json:"label,omitempty" url:"label,omitempty"`                       // 标签
	IterationID    string `json:"iteration_id,omitempty" url:"iteration_id,omitempty"`         // 迭代
	Module         string `json:"module,omitempty" url:"module,omitempty"`                     // 模块
	ReleaseID      int    `json:"release_id,omitempty" url:"release_id,omitempty"`             // 发布计划
	VersionReport  string `json:"version_report,omitempty" url:"version_report,omitempty"`     // 发现版本
	VersionTest    string `json:"version_test,omitempty" url:"version_test,omitempty"`         // 验证版本
	VersionFix     string `json:"version_fix,omitempty" url:"version_fix,omitempty"`           // 合入版本
	VersionClose   string `json:"version_close,omitempty" url:"version_close,omitempty"`       // 关闭版本
	BaselineFind   string `json:"baseline_find,omitempty" url:"baseline_find,omitempty"`       // 发现基线
	BaselineJoin   string `json:"baseline_join,omitempty" url:"baseline_join,omitempty"`       // 合入基线
	BaselineTest   string `json:"baseline_test,omitempty" url:"baseline_test,omitempty"`       // 验证基线
	BaselineClose  string `json:"baseline_close,omitempty" url:"baseline_close,omitempty"`     // 关闭基线
	Feature        string `json:"feature,omitempty" url:"feature,omitempty"`                   // 特性
	CurrentOwner   string `json:"current_owner,omitempty" url:"current_owner,omitempty"`       // 处理人
	CC             string `json:"cc,omitempty" url:"cc,omitempty"`                             // 抄送人
	Reporter       string `json:"reporter,omitempty" url:"reporter,omitempty"`                 // 创建人
	Participator   string `json:"participator,omitempty" url:"participator,omitempty"`         // 参与人
	TE             string `json:"te,omitempty" url:"te,omitempty"`                             // 测试人员
	DE             string `json:"de,omitempty" url:"de,omitempty"`                             // 开发人员
	Auditer        string `json:"auditer,omitempty" url:"auditer,omitempty"`                   // 审核人
	Confirmer      string `json:"confirmer,omitempty" url:"confirmer,omitempty"`               // 验证人
	Fixer          string `json:"fixer,omitempty" url:"fixer,omitempty"`                       // 修复人
	Closer         string `json:"closer,omitempty" url:"closer,omitempty"`                     // 关闭人
	LastModify     string `json:"lastmodify,omitempty" url:"lastmodify,omitempty"`             // 最后修改人
	InProgressTime string `json:"in_progress_time,omitempty" url:"in_progress_time,omitempty"` // 接受处理时间
	Resolved       string `json:"resolved,omitempty" url:"resolved,omitempty"`                 // 解决时间
	VerifyTime     string `json:"verify_time,omitempty" url:"verify_time,omitempty"`           // 验证时间
	Closed         string `json:"closed,omitempty" url:"closed,omitempty"`                     // 关闭时间
	RejectTime     string `json:"reject_time,omitempty" url:"reject_time,omitempty"`           // 拒绝时间
	Begin          string `json:"begin,omitempty" url:"begin,omitempty"`                       // 预计开始
	Due            string `json:"due,omitempty" url:"due,omitempty"`                           // 预计结束
	Deadline       string `json:"deadline,omitempty" url:"deadline,omitempty"`                 // 解决期限
	OS             string `json:"os,omitempty" url:"os,omitempty"`                             // 操作系统
	Platform       string `json:"platform,omitempty" url:"platform,omitempty"`                 // 软件平台
	TestMode       string `json:"testmode,omitempty" url:"testmode,omitempty"`                 // 测试方式
	TestPhase      string `json:"testphase,omitempty" url:"testphase,omitempty"`               // 测试阶段
	TestType       string `json:"testtype,omitempty" url:"testtype,omitempty"`                 // 测试类型
	Source         string `json:"source,omitempty" url:"source,omitempty"`                     // 缺陷根源
	BugType        string `json:"bugtype,omitempty" url:"bugtype,omitempty"`                   // 缺陷类型
	Frequency      string `json:"frequency,omitempty" url:"frequency,omitempty"`               // 重现规律
	OriginPhase    string `json:"originphase,omitempty" url:"originphase,omitempty"`           // 发现阶段
	SourcePhase    string `json:"sourcephase,omitempty" url:"sourcephase,omitempty"`           // 引入阶段
	Resolution     string `json:"resolution,omitempty" url:"resolution,omitempty"`             // 解决方法
	Estimate       int    `json:"estimate,omitempty" url:"estimate,omitempty"`                 // 预计解决时间
	Description    string `json:"description,omitempty" url:"description,omitempty"`           // 详细描述
}

type BugData struct {
	ID               string `json:"id"`                // 缺陷ID
	Title            string `json:"title"`             // 缺陷标题
	Description      string `json:"description"`       // 缺陷描述
	Priority         string `json:"priority"`          // 优先级
	Severity         string `json:"severity"`          // 严重性
	Module           string `json:"module"`            // 模块
	Status           string `json:"status"`            // 状态
	Reporter         string `json:"reporter"`          // 报告人
	Deadline         string `json:"deadline"`          // 截止日期
	Created          string `json:"created"`           // 创建时间
	BugType          string `json:"bugtype"`           // 缺陷类型
	Resolved         string `json:"resolved"`          // 解决时间
	Closed           string `json:"closed"`            // 关闭时间
	Modified         string `json:"modified"`          // 修改时间
	LastModify       string `json:"lastmodify"`        // 最后修改人
	Auditer          string `json:"auditer"`           // 审核人
	DE               string `json:"de"`                // 设计错误
	Fixer            string `json:"fixer"`             // 解决者
	VersionTest      string `json:"version_test"`      // 测试版本
	VersionReport    string `json:"version_report"`    // 报告版本
	VersionClose     string `json:"version_close"`     // 关闭版本
	VersionFix       string `json:"version_fix"`       // 修复版本
	BaselineFind     string `json:"baseline_find"`     // 发现基线
	BaselineJoin     string `json:"baseline_join"`     // 加入基线
	BaselineClose    string `json:"baseline_close"`    // 关闭基线
	BaselineTest     string `json:"baseline_test"`     // 测试基线
	SourcePhase      string `json:"sourcephase"`       // 引入阶段
	TE               string `json:"te"`                // 测试工程师
	CurrentOwner     string `json:"current_owner"`     // 当前负责人
	IterationID      string `json:"iteration_id"`      // 迭代ID
	Resolution       string `json:"resolution"`        // 解决方案
	Source           string `json:"source"`            // 来源
	OriginPhase      string `json:"originphase"`       // 发现阶段
	Confirmer        string `json:"confirmer"`         // 确认人
	Milestone        string `json:"milestone"`         // 里程碑
	Participator     string `json:"participator"`      // 参与者
	Closer           string `json:"closer"`            // 关闭者
	Platform         string `json:"platform"`          // 平台
	OS               string `json:"os"`                // 操作系统
	TestType         string `json:"testtype"`          // 测试类型
	TestPhase        string `json:"testphase"`         // 测试阶段
	Frequency        string `json:"frequency"`         // 频率
	CC               string `json:"cc"`                // 抄送
	RegressionNumber string `json:"regression_number"` // 回归次数
	Flows            string `json:"flows"`             // 流程
	Feature          string `json:"feature"`           // 特性
	TestMode         string `json:"testmode"`          // 测试模式
	Estimate         string `json:"estimate"`          // 预估时间
	IssueID          string `json:"issue_id"`          // 问题ID
	CreatedFrom      string `json:"created_from"`      // 创建来源
	InProgressTime   string `json:"in_progress_time"`  // 进行中时间
	VerifyTime       string `json:"verify_time"`       // 验证时间
	RejectTime       string `json:"reject_time"`       // 拒绝时间
	ReopenTime       string `json:"reopen_time"`       // 重新打开时间
	AuditTime        string `json:"audit_time"`        // 审计时间
	SuspendTime      string `json:"suspend_time"`      // 挂起时间
	Due              string `json:"due"`               // 到期时间
	Begin            string `json:"begin"`             // 开始时间
	ReleaseID        string `json:"release_id"`        // 发布ID
	Label            string `json:"label"`             // 标签
}

// GetBugsResponse 定义查询bug的响应结构
type GetBugsResponse struct {
	BaseResponse
	Data []struct {
		Bug BugData `json:"bug"`
	} `json:"data"`
}

// GetBugs 查询bug列表
// doc: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs.html
func (r *Bug) GetBugs(ctx context.Context, req *GetBugsRequest) (*GetBugsResponse, error) {
	resp := &GetBugsResponse{}
	err := r.cmd(ctx, http.MethodGet, "bugs", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetBugsCount 查询bug数量
// doc: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs_count.html
func (r *Bug) GetBugsCount(ctx context.Context, req *GetBugsRequest) (*CountResponse, error) {
	resp := &CountResponse{}
	err := r.cmd(ctx, http.MethodGet, "bugs/count", req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
