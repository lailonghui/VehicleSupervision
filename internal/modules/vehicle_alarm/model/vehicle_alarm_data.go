// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 报警数据表
type VehicleAlarmData struct {
	// 按指定方法生成                                          ( 主键                               )
	ID int64 `json:"id"`
	// 报警数据外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                           )
	VehicleAlarmDataID string `json:"vehicle_alarm_data_id"`
	// 车辆ID                                                  ( 引用vehicle_info表的vehicle_id )
	VehicleID *string `json:"vehicle_id"`
	// 报警类型                                                ( 报警类型字典                   )
	AlarmType *string `json:"alarm_type"`
	// 报警开始时间                                            (                                    )
	AlarmStartTime *time.Time `json:"alarm_start_time"`
	// 报警结束时间                                            (                                    )
	AlarmEndTime *time.Time `json:"alarm_end_time"`
	// 报警结束位置                                            (                                    )
	AlarmEndPosition *string `json:"alarm_end_position"`
	// 最新报警时间                                            (                                    )
	LatestAlarmTime *time.Time `json:"latest_alarm_time"`
	// 最新报警位置                                            (                                    )
	LatestAlarmPosition *int `json:"latest_alarm_position"`
	// 报警是否有效                                            (                                    )
	IsAlarmEffective *bool `json:"is_alarm_effective"`
	// 报警是否结束                                            (                                    )
	IsAlarmOver *bool `json:"is_alarm_over"`
	// 是否取消报警                                            (                                    )
	IsCancelAlarm *bool `json:"is_cancel_alarm"`
	// 报警来源:1.终端报警 2.平台报警                          ( 报警来源字典                   )
	AlarmSource *string `json:"alarm_source"`
	// 处理时间                                                (                                    )
	ProcessingTime *time.Time `json:"processing_time"`
	// 处理方式                                                ( 处警处理方式字典               )
	ProcessingMethod *string `json:"processing_method"`
	// 处理描述                                                (                                    )
	ProcessingDescription *string `json:"processing_description"`
	// 处理人                                                  ( system_user表的user_id         )
	Processor *string `json:"processor"`
	// 处理状态                                                ( 处警处理状态字典               )
	ProcessingStatus *string `json:"processing_status"`
	// 行驶记录仪速度                                          (                                    )
	TachographSpeed *float64 `json:"tachograph_speed"`
	// GPS速度                                                 (                                    )
	GpsSpeed *float64 `json:"gps_speed"`
	// 最高速度                                                (                                    )
	MaximumSpeed *float64 `json:"maximum_speed"`
	// 限速阀值                                                (                                    )
	SpeedLimitThreshold *float64 `json:"speed_limit_threshold"`
	// 空间数据类型point表示经度(longitude)和纬度(latitude)    (                                    )
	Coordinate *string `json:"coordinate"`
	// 位置描述                                                (                                    )
	LocationDescription *string `json:"location_description"`
	// 持续时间                                                (                                    )
	Duration *string `json:"duration"`
	// 道路等级                                                ( 道路等级字典                   )
	RoadGrade *string `json:"road_grade"`
	// 道路名称                                                (                                    )
	RoadName *string `json:"road_name"`
	// 进区域ID                                                (                                    )
	AreaID *string `json:"area_id"`
	// 处警ID                                                  (                                    )
	AlarmDealID *string `json:"alarm_deal_id"`
	// 地区                                                    (                                    )
	Pid *string `json:"pid"`
	// 记录时间                                                (                                    )
	RecordTime *time.Time `json:"record_time"`
	// 监管人                                                  ( system_user表的user_id         )
	Supervisor *string `json:"supervisor"`
	// 管理部门是否监管                                        (                                    )
	IsSupervise *bool `json:"is_supervise"`
	// 管理部门监管时间                                        (                                    )
	SupervisionTime *time.Time `json:"supervision_time"`
	// 监管备注                                                (                                    )
	SupervisionNote *string `json:"supervision_note"`
	// 是否解析                                                (                                    )
	IsResolve *bool `json:"is_resolve"`
	// 工地是否处理                                            (                                    )
	IsConstructionSiteHandle *bool `json:"is_construction_site_handle"`
	// 工地处理时间                                            (                                    )
	ConstructionSiteHandleTime *time.Time `json:"construction_site_handle_time"`
	// 创建时间                                                (                                    )
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                  ( system_user表的user_id         )
	CreatedBy string `json:"created_by"`
	// 修改时间                                                (                                    )
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                  ( system_user表的user_id         )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间                                                (                                    )
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                  ( system_user表的user_id         )
	DeletedBy *string `json:"deleted_by"`
}
