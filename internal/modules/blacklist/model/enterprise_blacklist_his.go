// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 企业黑名单操作纪录
type EnterpriseBlacklistHis struct {
	// ID
	ID int64 `json:"id"`
	// 黑名单记录ID
	HisID string `json:"his_id"`
	// 企业ID
	EnterpriseID string `json:"enterprise_id"`
	// 黑名单类别
	BlacklistType int `json:"blacklist_type"`
	// 操作类别
	Operate int `json:"operate"`
	// 备注
	Remarks *string `json:"remarks"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
}
