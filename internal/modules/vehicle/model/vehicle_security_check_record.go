// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆安全检查记录表
type VehicleSecurityCheckRecord struct {
	// 按指定方法生成                                  ( 主键                                 )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                             )
	VehicleSecurityCheckRecordID string `json:"vehicle_security_check_record_id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 所在企业id                                      ( enterprise_info表的enterprise_id )
	EnterpriseID *string `json:"enterprise_id"`
	// 刹车
	Brake *int `json:"brake"`
	// 轮胎
	Tire *int `json:"tire"`
	// 螺丝
	Screw *int `json:"screw"`
	// 液压油
	HydraulicOil *int `json:"hydraulic_oil"`
	// 机油
	EngineOil *int `json:"engine_oil"`
	// 水
	Water *int `json:"water"`
	// 大灯
	Headlight *int `json:"headlight"`
	// 尾灯
	Taillight *int `json:"taillight"`
	// 转向灯
	TurnSignal *int `json:"turn_signal"`
	// 刹车灯
	BrakeLight *int `json:"brake_light"`
	// 最后检查时间
	LastCheckTime *time.Time `json:"last_check_time"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id           )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id           )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id           )
	DeletedBy *string `json:"deleted_by"`
}
