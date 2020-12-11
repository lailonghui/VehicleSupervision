/*
@Time : 2020/12/11 14:38
@Author : lai
@Description :
@File : district_alarm_content_push
*/
package model

import "time"

// 各县市区报警内容推送表
//
//
// columns and relationships of "district_alarm_content_push"
type DistrictAlarmContentPush struct {
	// 内容
	AlarmContent *string `json:"alarm_content"`
	// vehicle_alarm_data报警数据表的alarm_data_id
	AlarmDataID string `json:"alarm_data_id"`
	// 报警类型字典
	AlarmType *string `json:"alarm_type"`
	// 城市ID
	CityID *string `json:"city_id"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 区ID
	DistrictID *string `json:"district_id"`
	// 主键
	ID int64 `json:"id"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 省份ID
	ProvinceID *string `json:"province_id"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
}

func (DistrictAlarmContentPush) TableName() string {
	return "district_alarm_content_push"
}
