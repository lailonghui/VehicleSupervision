package model

import "time"

// 行车日志
//
//
// columns and relationships of "driving_log"
type DrivingLog struct {
	// 事由
	Cause *string `json:"cause"`
	// 审核机构级别
	CheckOrganizationLevel int `json:"check_organization_level"`
	// 审核状态
	CheckState int `json:"check_state"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 驾驶员ID
	DriverID string `json:"driver_id"`
	// 用车结束日期
	DrivingEndTime time.Time `json:"driving_end_time"`
	// 行车日志ID
	DrivingLogID string `json:"driving_log_id"`
	// 用车起始日期
	DrivingStartTime time.Time `json:"driving_start_time"`
	// 结束时间
	EndTime time.Time `json:"end_time"`
	// ID
	ID int64 `json:"id"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 是否补录
	IsFill bool `json:"is_fill"`
	// 登记时间
	RegisterAt *time.Time `json:"register_at"`
	// 登记人
	RegisterBy *string `json:"register_by"`
	// 备注
	Remarks *string `json:"remarks"`
	// 路线
	Route *string `json:"route"`
	// 开始时间
	StartTime time.Time `json:"start_time"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 车辆ID
	VehicleID string `json:"vehicle_id"`
}

func (d DrivingLog) TableName() string {
	return "driving_log"
}
