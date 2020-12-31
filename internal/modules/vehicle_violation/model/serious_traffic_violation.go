// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 严重交通违法行为表
type SeriousTrafficViolation struct {
	// 按指定方法生成                                  ( 主键                            )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                        )
	SeriousTrafficViolationID string `json:"serious_traffic_violation_id"`
	// 违章车辆id                                      ( vehicle_info表的vehicle_id  )
	VehicleID *string `json:"vehicle_id"`
	// 违法代码                                        ( VIO_CODEWFDM 违法描述字典表 )
	IllegalCode *string `json:"illegal_code"`
	// 违法日期
	IllegalTime *time.Time `json:"illegal_time"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id      )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id      )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id      )
	DeletedBy *string `json:"deleted_by"`
}