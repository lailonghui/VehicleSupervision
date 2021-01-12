// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 工程跨区表
type EcdFileDistrict struct {
	// ID
	ID int64 `json:"id"`
	// 工程跨区表ID
	FileDistrictID string `json:"file_district_id"`
	// 报备主表ID
	FileMainID string `json:"file_main_id"`
	// 区域ID
	DistrictID string `json:"district_id"`
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
}
