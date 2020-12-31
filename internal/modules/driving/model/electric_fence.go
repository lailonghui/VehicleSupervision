// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 电子围栏
type ElectricFence struct {
	// ID
	ID int64 `json:"id"`
	// 电子围栏ID
	ElectricFenceID string `json:"electric_fence_id"`
	// 区域类型(1矩形 2圆形 3多边形)
	AreaType int `json:"area_type"`
	// 围栏类型(1消纳场 2工地 3工程 4重点区域)
	FenceType *int `json:"fence_type"`
	// 围栏名称
	FenceName *string `json:"fence_name"`
	// 地点
	Address *string `json:"address"`
	// 所属区域
	DistrictID *string `json:"district_id"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 更新时间
	UpdateAt *time.Time `json:"update_at"`
	// 更新人
	UpdateBy *string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 围栏数据
	GisData *string `json:"gis_data"`
}