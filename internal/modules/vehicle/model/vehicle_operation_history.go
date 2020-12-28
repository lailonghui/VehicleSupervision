// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆操作历史记录表
type VehicleOperationHistory struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	VehicleOperationHistoryID string `json:"vehicle_operation_history_id"`
	// vehicle_info 车辆信息表 的vehicle_id        (                            )
	VehicleID string `json:"vehicle_id"`
	// 备注                                            (                            )
	Remarks *string `json:"remarks"`
	// 操作类型  1.添加  2.删除                        (                            )
	OperationType *int `json:"operation_type"`
	// 操作人                                          ( system_user表的user_id )
	Operator *string `json:"operator"`
	// 审核状态  0.未审批  1.已审批                    (                            )
	ReviewStatus *int `json:"review_status"`
	// 地区                                            (                            )
	Area *string `json:"area"`
	// 审核人                                          ( system_user表的user_id )
	Reviewer *string `json:"reviewer"`
	// 创建时间                                        (                            )
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id )
	CreatedBy string `json:"created_by"`
	// 修改时间                                        (                            )
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间                                        (                            )
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id )
	DeletedBy *string `json:"deleted_by"`
}
