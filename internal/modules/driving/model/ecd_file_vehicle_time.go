// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 工程报备-车辆运行时间表
type EcdFileVehicleTime struct {
	// ID
	ID int64 `json:"id"`
	// 工程审批-车辆运行时间表ID
	VehicleTimeID string `json:"vehicle_time_id"`
	// 报备主表ID
	FileMainID string `json:"file_main_id"`
	// 车辆ID
	VechileID string `json:"vechile_id"`
	// 运行起始日期
	StartDate *time.Time `json:"start_date"`
	// 运行截止日期
	EndDate *time.Time `json:"end_date"`
	// 运行起始时间
	StartTime *time.Time `json:"start_time"`
	// 运行截止时间
	EndTime *time.Time `json:"end_time"`
	// 审批状态 0未审批 1正常 2废弃
	CheckStatus *int `json:"check_status"`
	// 是否审批完成
	IsCheck bool `json:"is_check"`
	// 审批时间
	CheckTime *time.Time `json:"check_time"`
	// 路线ID
	LineID *string `json:"line_id"`
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
