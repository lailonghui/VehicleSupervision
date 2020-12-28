// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 视频平台报警类型字典
type VideoPlatformAlarmType struct {
	// 按指定方法生成                                  ( 主键                                                      )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                                                  )
	VideoPlatformAlarmTypeID string `json:"video_platform_alarm_type_id"`
	// 报警数据表id                                    ( vehicle_alarm_data  报警数据表的vehicle_alarm_data_id )
	VehicleAlarmDataID *int `json:"vehicle_alarm_data_id"`
	// 报警类型
	AlarmType *string `json:"alarm_type"`
	// 报警来源
	AlarmSource *string `json:"alarm_source"`
	// 报警分类                                        ( Adas报警字典                                          )
	AlarmClassify *string `json:"alarm_classify"`
	// 报警代码
	AlarmCode *string `json:"alarm_code"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id                                )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id                                )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id                                )
	DeletedBy *string `json:"deleted_by"`
}
