/*
@Time : 2020/12/11 14:35
@Author : lai
@Description :
@File : dynamic_supervision_detail
*/
package model

import "time"

// 动态监管抽查明细表
//
//
// columns and relationships of "dynamic_supervision_detail"
type DynamicSupervisionDetail struct {
	// 卫星定位速度
	GpsSpeed *string `json:"GPS_speed"`
	// 受理人
	Assignee *string `json:"assignee"`
	// 经营范围字典
	BusinessScope *int `json:"business_scope"`
	// 空间数据类型point表示经纬度
	Coordinate *string `json:"coordinate"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 曲线情况（曲线完整/回传异常/零速度）
	Curve *int `json:"curve"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 处置措施
	DisposalMeasures *string `json:"disposal_measures"`
	// 是否在线处置措施
	DisposalMeasures1 *string `json:"disposal_measures1"`
	// 是否超速处置措施
	DisposalMeasures2 *string `json:"disposal_measures2"`
	// 曲线情况处置措施
	DisposalMeasures3 *string `json:"disposal_measures3"`
	// 客运疲劳驾驶处置措施
	DisposalMeasures4 *string `json:"disposal_measures4"`
	// 客运凌晨停运处置措施
	DisposalMeasures5 *string `json:"disposal_measures5"`
	// 行车记录仪数据处置措施
	DisposalMeasures6 *string `json:"disposal_measures6"`
	// 轨迹情况处置措施
	DisposalMeasures7 *string `json:"disposal_measures7"`
	// 处置结果
	DisposalResults *string `json:"disposal_results"`
	// 是否在线处置结果
	DisposalResults1 *string `json:"disposal_results1"`
	// 是否超速处置结果
	DisposalResults2 *string `json:"disposal_results2"`
	// 曲线情况处置结果
	DisposalResults3 *string `json:"disposal_results3"`
	// 客运疲劳驾驶处置结果
	DisposalResults4 *string `json:"disposal_results4"`
	// 客运疲劳驾驶处置结果
	DisposalResults5 *string `json:"disposal_results5"`
	// 行车记录仪数据处置结果
	DisposalResults6 *string `json:"disposal_results6"`
	// 轨迹情况处置结果
	DisposalResults7 *string `json:"disposal_results7"`
	// 驾驶员id
	DriverID *string `json:"driver_id"`
	// 车辆所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 设备情况（图像正常/无图像/摄像头 号损坏）
	Equipment *int `json:"equipment"`
	// 疲劳驾驶报警时间
	FatigueAlarmTime *time.Time `json:"fatigue_alarm_time"`
	// 反馈时间
	FeedbackTime *time.Time `json:"feedback_time"`
	// 主键
	ID int64 `json:"id"`
	// 是否被删除
	IsDelete *bool `json:"is_delete"`
	// 客运疲劳驾驶（是/否）
	IsFatigueDriving *bool `json:"is_fatigue_driving"`
	// 是否定位
	IsLocate *bool `json:"is_locate"`
	// 客运凌晨2-5时停运（是/否）
	IsMorningOutage *bool `json:"is_morning_outage"`
	// 是否在线（是/否）
	IsOnline *bool `json:"is_online"`
	// 是否发送
	IsSend *bool `json:"is_send"`
	// 是否超速（是/否）
	IsSpeeding *bool `json:"is_speeding"`
	// 行车记录仪数据（是否异常）
	IsTachographRecordNormal *bool `json:"is_tachograph_record_normal"`
	// 经纬度描述
	LatitudeLongitudeDescription *string `json:"latitude_longitude_description"`
	// 摄像头损坏号
	LensOn *string `json:"lens_on"`
	// 镜头位置（正/偏）
	LensPosition *int `json:"lens_position"`
	// 监管费到期时间
	MonitorEndTime *time.Time `json:"monitor_end_time"`
	// 监控平台显示位置
	MonitoringLocation *string `json:"monitoring_location"`
	// 监控平台时间
	MonitoringTime *time.Time `json:"monitoring_time"`
	// 其他违规
	OtherInfraction *string `json:"other_infraction"`
	// 轨迹其他情况
	Others *string `json:"others"`
	// 凌晨2点到5点停运报警时间
	OutageAlarmTime *time.Time `json:"outage_alarm_time"`
	// 备注
	Remarks *string `json:"remarks"`
	// 超速报警时间
	SpeedAlarmTime *time.Time `json:"speed_alarm_time"`
	// 超速速度
	SpeedingSpeed *string `json:"speeding_speed"`
	// 联合主键
	SupervisionDetailID string `json:"supervision_detail_id"`
	// 动态监管抽查主表dynamic_supervision的supervision_id
	SupervisionID *string `json:"supervision_id"`
	// 行车记录仪异常数据项
	TachographDataException *string `json:"tachograph_data_exception"`
	// 行车记录仪速度
	TachographSpeed *string `json:"tachograph_speed"`
	// 出租空/重车状态（空/重）
	TaxiState *int `json:"taxi_state"`
	// 轨迹情况（正常/漂移/其他）
	Trail *int `json:"trail"`
	// 处置时间
	TreatmentTime *time.Time `json:"treatment_time"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 车辆id
	VehicleID *string `json:"vehicle_id"`
}

func (DynamicSupervisionDetail) TableName() string {
	return "dynamic_supervision"
}
