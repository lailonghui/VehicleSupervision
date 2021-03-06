// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 违法信息登记表
type ViolationRegistration struct {
	// 按指定方法生成                                  ( 主键                           )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                       )
	ViolationRegistrationID string `json:"violation_registration_id"`
	// 违章车辆id                                      ( vehicle_info表的vehicle_id )
	VehicleID *string `json:"vehicle_id"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 码身份证号
	IDCardNum *string `json:"id_card_num"`
	// 姓名
	Name *string `json:"name"`
	// 地点
	Location *string `json:"location"`
	// 原因
	Cause *string `json:"cause"`
	// 违法时间
	IllegalTime *time.Time `json:"illegal_time"`
	// 操作人                                          ( system_user表的user_id     )
	Operator *string `json:"operator"`
	// 违法地区
	IllegalArea *string `json:"illegal_area"`
	// 违法代码
	IllegalCode *string `json:"illegal_code"`
	// 车辆所属地区
	VehicleArea *string `json:"vehicle_area"`
	// 车辆所属单位
	VehicleEnterprise *string `json:"vehicle_enterprise"`
	// 所在省                                          ( 省份表province_id          )
	ProvinceID *string `json:"province_id"`
	// 所在市                                          ( 城市表city_id              )
	CityID *string `json:"city_id"`
	// 所在县                                          ( 区域表district_id          )
	DistrictID *string `json:"district_id"`
	// 监管人
	Supervisor *string `json:"supervisor"`
	// 监管时间
	SupervisionTime *string `json:"supervision_time"`
	// 监管备注
	SepervisionRemarks *string `json:"sepervision_remarks"`
	// 是否监管
	IsSupervised *string `json:"is_supervised"`
	// 是否事故
	IsAccident *string `json:"is_accident"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id     )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id     )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id     )
	DeletedBy *string `json:"deleted_by"`
}
