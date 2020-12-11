/*
@Time : 2020/12/11 14:38
@Author : lai
@Description :
@File : vehicle_alarm_data
*/
package model

import "time"

// 报警数据表
//
//
// columns and relationships of "vehicle_alarm_data"
type VehicleAlarmData struct {
	// GPS速度
	GpsSpeed *float64 `json:"GPS_speed"`
	// 处警ID
	AlarmDealID *string `json:"alarm_deal_id"`
	// 报警结束位置
	AlarmEndPosition *string `json:"alarm_end_position"`
	// 报警结束时间
	AlarmEndTime *time.Time `json:"alarm_end_time"`
	// 报警来源字典
	AlarmSource *string `json:"alarm_source"`
	// 报警开始时间
	AlarmStartTime *time.Time `json:"alarm_start_time"`
	// 报警类型字典
	AlarmType *string `json:"alarm_type"`
	// 进区域ID
	AreaID *string `json:"area_id"`
	// 空间数据类型point表示经度(longitude)和纬度(latitude)
	Coordinate *string `json:"coordinate"`
	// 持续时间
	Duration *string `json:"duration"`
	// 主键
	ID int64 `json:"id"`
	// 报警是否有效
	IsAlarmEffective *bool `json:"is_alarm_effective"`
	// 报警是否结束
	IsAlarmOver *bool `json:"is_alarm_over"`
	// 是否取消报警
	IsCancelAlarm *bool `json:"is_cancel_alarm"`
	// 是否解析
	IsResolve *bool `json:"is_resolve"`
	// 是否监管
	IsSupervise *bool `json:"is_supervise"`
	// 最新报警位置
	LatestAlarmPosition *int `json:"latest_alarm_position"`
	// 最新报警时间
	LatestAlarmTime *time.Time `json:"latest_alarm_time"`
	// 位置描述
	LocationDescription *string `json:"location_description"`
	// 最高速度
	MaximumSpeed *float64 `json:"maximum_speed"`
	// 地区
	Pid *string `json:"pid"`
	// 处理描述
	ProcessingDescription *string `json:"processing_description"`
	// 处警处理方式字典
	ProcessingMethod *string `json:"processing_method"`
	// 处警处理状态字典
	ProcessingStatus *string `json:"processing_status"`
	// 处理时间
	ProcessingTime *time.Time `json:"processing_time"`
	// 处理人
	Processor *string `json:"processor"`
	// 记录时间
	RecordTime *time.Time `json:"record_time"`
	// 道路等级字典
	RoadGrade *string `json:"road_grade"`
	// 道路名称
	RoadName *string `json:"road_name"`
	// 限速阀值
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	// 监管备注
	SupervisionNote *string `json:"supervision_note"`
	// 监管时间
	SupervisionTime *time.Time `json:"supervision_time"`
	// 监管人
	Supervisor *string `json:"supervisor"`
	// 行驶记录仪速度
	TachographSpeed *float64 `json:"tachograph_speed"`
	// 报警数据外部编码，由golang程序生成的xid，暴露到外部使用，联合主键
	VehicleAlarmDataID string `json:"vehicle_alarm_data_id"`
	// 车辆ID
	VehicleID *string `json:"vehicle_id"`
}

func (VehicleAlarmData) TableName() string {
	return "vehicle_alarm_data"
}
