/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : dispute_violation_record_log
*/
package model

import "time"

// 违章争议审批日志表
//
//
// columns and relationships of "dispute_violation_record_log"
type DisputeViolationRecordLog struct {
	// 审批人
	Approver *string `json:"approver"`
	// 违章争议记录表id(dispute_violation_record的dispute_violation_id)
	DisputeViolationID *string `json:"dispute_violation_id"`
	// 联合主键
	DisputeViolationLogID string `json:"dispute_violation_log_id"`
	// 主键
	ID int64 `json:"id"`
	// 审核动作名称
	ReviewActionName *string `json:"review_action_name"`
	// 审核意见
	ReviewOpinion *string `json:"review_opinion"`
	// 审核结果
	ReviewResult *string `json:"review_result"`
	// 审核时间
	ReviewTime *time.Time `json:"review_time"`
	// 审核人
	Reviewer *string `json:"reviewer"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
}

func (DisputeViolationRecordLog) TableName() string {
	return "dispute_violation_record_log"
}
