// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车辆报警监管
type VehicleAlarmSupervision struct {
	// 按指定方法生成                                  ( 主键                                                         )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                                                     )
	VehicleAlarmSupervisionID string `json:"vehicle_alarm_supervision_id"`
	// 监管单位id                                      (                                                              )
	SuperivisionAuthorityID *string `json:"superivision_authority_id"`
	// 监管类型                                        ( 地区机构监管，企业监管，执法机构监管，执法人员监管，地区抽查机构监管 )
	SuperivisionType *string `json:"superivision_type"`
	// 区域ID                                          (                                                              )
	AreaID *string `json:"area_id"`
	// 报警车辆数                                      (                                                              )
	VehicleAlarmNumber *int `json:"vehicle_alarm_number"`
	// 报警类型                                        (                                                              )
	AlarmType *string `json:"alarm_type"`
	// 报警次数                                        (                                                              )
	AlarmTimes *int `json:"alarm_times"`
	// 处置次数                                        (                                                              )
	DisposalTimes *int `json:"disposal_times"`
	// 报警处置率                                      (                                                              )
	AlarmDisposalRate *float64 `json:"alarm_disposal_rate"`
	// 县级提出的整改                                  (                                                              )
	DistrictRectification *string `json:"district_rectification"`
	// 市级提出的整改                                  (                                                              )
	CityRectification *string `json:"city_rectification"`
	// 省级提出的整改                                  (                                                              )
	ProvinceRectification *string `json:"province_rectification"`
	// 部级提出的整改                                  (                                                              )
	CountryRectification *string `json:"country_rectification"`
	// 登记时间                                        (                                                              )
	RegistrationTime *time.Time `json:"registration_time"`
	// 统计日期                                        (                                                              )
	StatisticsDate *string `json:"statistics_date"`
	// 应监管企业数                                    (                                                              )
	ShouldSupervisionEnterpriseNumber *int `json:"should_supervision_enterprise_number"`
	// 实监管企业数                                    (                                                              )
	ActualSupervisionEnterpriseNumber *int `json:"actual_supervision_enterprise_number"`
	// 监管率                                          (                                                              )
	SupervisionRate *float64 `json:"supervision_rate"`
	// 经营范围                                        ( 经营范围字典                                             )
	BusinessScope *int `json:"business_scope"`
	// 应检查车辆数                                    (                                                              )
	ShouldCheckVehicleNumber *int `json:"should_check_vehicle_number"`
	// 实检查车辆数                                    (                                                              )
	ActualCheckVechicleNumber *int `json:"actual_check_vechicle_number"`
	// 检查异常车辆数                                  (                                                              )
	CheckAbnormalVehicleNumber *int `json:"check_abnormal_vehicle_number"`
	// 异常处置数                                      (                                                              )
	AbnormalDisposalNumber *int `json:"abnormal_disposal_number"`
	// 异常处置率                                      (                                                              )
	AbnormalDisposalRate *float64 `json:"abnormal_disposal_rate"`
	// 创建时间                                        (                                                              )
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id                                   )
	CreatedBy string `json:"created_by"`
	// 修改时间                                        (                                                              )
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id                                   )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间                                        (                                                              )
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id                                   )
	DeletedBy *string `json:"deleted_by"`
}
