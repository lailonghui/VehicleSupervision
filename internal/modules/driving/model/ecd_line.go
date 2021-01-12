// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 工程审批路线表
type EcdLine struct {
	// ID
	ID int64 `json:"id"`
	// 工程审批路线表ID
	LineID string `json:"line_id"`
	// 路线名称
	LineName string `json:"line_name"`
	// 所属企业ID
	EnterpriseID *string `json:"enterprise_id"`
	// 路线描述
	LineDesc *string `json:"line_desc"`
	// 审批人
	AuditUserID *string `json:"audit_user_id"`
	// 审批时间
	AuditTime *time.Time `json:"audit_time"`
	// 是否解析描述
	IsResolveDesc bool `json:"is_resolve_desc"`
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
