// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆状态变更记录表
type VehicleStatusChangeLog struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	VehicleStatusChangeLogID string `json:"vehicle_status_change_log_id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 终端ID
	TerminalID *string `json:"terminal_id"`
	// 开始时间
	StartTime *time.Time `json:"start_time"`
	// 结束时间
	EndTime *time.Time `json:"end_time"`
	// 车辆状态类型(车厢状态,举升状态,载重状态)        ( 车辆状态类型字典       )
	VehicleStatusType *int `json:"vehicle_status_type"`
	// 值
	Value *string `json:"value"`
	// 是否完成
	IsCompleted *bool `json:"is_completed"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id )
	DeletedBy *string `json:"deleted_by"`
}