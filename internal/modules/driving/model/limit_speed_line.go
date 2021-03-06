// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 限速路线
type LimitSpeedLine struct {
	// ID
	ID int64 `json:"id"`
	// 限速路线ID
	LimitSpeedLineID string `json:"limit_speed_line_id"`
	// 路线名称
	LineName *string `json:"line_name"`
	// 路线类型(1路线 2区域)
	LineType *int `json:"line_type"`
	// 区域ID
	DistrictID *string `json:"district_id"`
	// 管控等级
	ControlLevel *int `json:"control_level"`
	// 申请原因
	ApplyReason *string `json:"apply_reason"`
	// 审核状态
	CheckStatus *int `json:"check_status"`
	// 审核时间
	CheckTime *time.Time `json:"check_time"`
	// 审核人ID
	CheckUserID *string `json:"check_user_id"`
	// 退回原因
	RejectReason *string `json:"reject_reason"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 更新时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 更新人
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
	// 线路数据
	GisData *string `json:"gis_data"`
}
