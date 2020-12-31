// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 同一车辆报警次数记录表
type VehicleAlarmTimesRecord struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	VehicleAlarmTimesRecordID string `json:"vehicle_alarm_times_record_id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 报警类型                                        ( 报警类型字典           )
	AlarmType *string `json:"alarm_type"`
	// 处置措施
	DisposalMeasure *string `json:"disposal_measure"`
	// 处置时间
	DisposalTime *time.Time `json:"disposal_time"`
	// 处置结果
	DisposalResult *string `json:"disposal_result"`
	// 是否处置
	IsDisposal *bool `json:"is_disposal"`
	// 处置方式                                        ( 处置方式字典           )
	DisposalMethod *int `json:"disposal_method"`
	// 值班人
	DutyPerson *string `json:"duty_person"`
	// 报警次数
	AlarmTimes *string `json:"alarm_times"`
	// 备注
	Remarks *string `json:"remarks"`
	// 记录时间
	RecordTime *time.Time `json:"record_time"`
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