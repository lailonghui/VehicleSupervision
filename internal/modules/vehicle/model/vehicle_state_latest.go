// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 新型车辆最新状态表
type VehicleStateLatest struct {
	// 按指定方法生成                                  ( 主键                                 )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                             )
	VehicleStateLatestID string `json:"vehicle_state_latest_id"`
	// vehicle_info 车辆信息表 的vehicle_id        (                                      )
	VehicleID string `json:"vehicle_id"`
	// 操作类型（lock.锁车 speed.限速 lift.限举）      (                                      )
	OperationType *string `json:"operation_type"`
	// 操作人                                          ( system_user表的user_id           )
	Operator *string `json:"operator"`
	// 操作人单位                                      ( enterprise_info表的enterprise_id )
	OperatorInstitution *string `json:"operator_institution"`
	// 状态                                            (                                      )
	Status *string `json:"status"`
	// 限速值                                          (                                      )
	SpeedLimit *string `json:"speed_limit"`
	// 创建时间                                        (                                      )
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id           )
	CreatedBy string `json:"created_by"`
	// 修改时间                                        (                                      )
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id           )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间                                        (                                      )
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id           )
	DeletedBy *string `json:"deleted_by"`
}
