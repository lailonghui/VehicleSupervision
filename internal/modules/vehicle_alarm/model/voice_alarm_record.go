// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 语音报警记录
type VoiceAlarmRecord struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	VioceAlarmRecordID string `json:"vioce_alarm_record_id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 报警时间
	AlarmTime *time.Time `json:"alarm_time"`
	// 报警类型
	AlarmType *string `json:"alarm_type"`
	// 提醒时间
	RemindTime *time.Time `json:"remind_time"`
	// 提醒内容
	RemindContent *string `json:"remind_content"`
	// 录入人                                          ( system_user表的user_id )
	InputPerson *string `json:"input_person"`
	// 录入时间
	InputTime *time.Time `json:"input_time"`
	// 是否成功
	IsSuccess *bool `json:"is_success"`
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
