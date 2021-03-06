// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 行车日志信息
type DrivingLogInfo struct {
	// 按指定方法生成                                  ( 主键                           )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                       )
	DrivingLogInfoID string `json:"driving_log_info_id"`
	// 车辆id                                          ( vehicle_info表的vehicle_id )
	VehicleID *string `json:"vehicle_id"`
	// 驾驶员id                                        ( driver_info表的driver_id   )
	DriverID *string `json:"driver_id"`
	// 用车起始日期
	DrivingStartTime *time.Time `json:"driving_start_time"`
	// 用车结束日期
	DrivingEndTime *time.Time `json:"driving_end_time"`
	// 事由
	Cause *string `json:"cause"`
	// 路线
	Route *string `json:"route"`
	// 备注
	Remarks *string `json:"remarks"`
	// 开始时间
	StartTime *string `json:"start_time"`
	// 结束时间
	EndTime *string `json:"end_time"`
	// 审核状态
	ReviewStatus *int `json:"review_status"`
	// 审核机构级别
	ReviewAgecyLevel *int `json:"review_agecy_level"`
	// 是否补录
	IsMarkup *bool `json:"is_markup"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id     )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id     )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id     )
	DeletedBy *string `json:"deleted_by"`
}
