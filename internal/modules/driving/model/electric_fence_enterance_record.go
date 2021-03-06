// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 电子围栏进出记录表
type ElectricFenceEnteranceRecord struct {
	// ID
	ID int64 `json:"id"`
	// 电子围栏进出记录表ID
	RecordID string `json:"record_id"`
	// 车辆ID
	VehicleID string `json:"vehicle_id"`
	// 电子围栏ID
	ElectricFenceID string `json:"electric_fence_id"`
	// 进入时间
	InTime *time.Time `json:"in_time"`
	// 离开时间
	OutTime *time.Time `json:"out_time"`
	// 是否离开
	IsOut *bool `json:"is_out"`
	// 位置点
	Position *string `json:"position"`
	// 是否在线
	IsOnline *bool `json:"is_online"`
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
