// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 各县市区报警内容推送表
type DistrictAlarmContentPush struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// vehicle_alarm_data报警数据表的alarm_data_id
	AlarmDataID string `json:"alarm_data_id"`
	// 报警类型                                        ( 报警类型字典           )
	AlarmType *string `json:"alarm_type"`
	// 内容
	AlarmContent *string `json:"alarm_content"`
	// 省份ID                                          ( 省份表province_id      )
	ProvinceID *string `json:"province_id"`
	// 城市ID                                          ( 城市表city_id          )
	CityID *string `json:"city_id"`
	// 区ID                                            ( 区域表district_id      )
	DistrictID *string `json:"district_id"`
	// 是否删除                                        ( false                      )
	IsDeleted *bool `json:"is_deleted"`
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
