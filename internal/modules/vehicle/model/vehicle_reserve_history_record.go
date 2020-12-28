// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆预备库历史记录表
type VehicleReserveHistoryRecord struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	VehicleReserveHistoryRecordID string `json:"vehicle_reserve_history_record_id"`
	// vehicle_info 车辆信息表 的vehicle_id        (                            )
	VehicleID string `json:"vehicle_id"`
	// 操作 1.加入预备库  2.移出预备库                 (                            )
	Operation *int `json:"operation"`
	// 操作用户                                        ( system_user表的user_id )
	OperationUser *string `json:"operation_user"`
	// 操作时间                                        (                            )
	OperationTime *time.Time `json:"operation_time"`
	// 操作来源 1.车辆  2.驾驶员                       (                            )
	OperationSource *int `json:"operation_source"`
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
