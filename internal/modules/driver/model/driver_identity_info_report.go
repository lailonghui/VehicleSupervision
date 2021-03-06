// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 驾驶员身份信息采集上报
type DriverIdentityInfoReport struct {
	// 按指定方法生成                                  ( 主键                       )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                   )
	DriverIdentityInfoReportID string `json:"driver_identity_info_report_id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// IC状态(从业资格证IC卡插入,从业资格证IC卡拔出)   ( 从业资格证IC卡字典     )
	IcCardStatus *int `json:"ic_card_status"`
	// 操作时间
	OperationTime *time.Time `json:"operation_time"`
	// 驾驶员姓名
	DriverName *string `json:"driver_name"`
	// 证件号码
	LicenseNumber *string `json:"license_number"`
	// 终端IMEI                                        ( 国际移动设备标识别码       )
	Imel *string `json:"imel"`
	// IC卡读取结果                                    ( IC卡读卡字典           )
	IcCardReadingResult *string `json:"ic_card_reading_result"`
	// 从业资格证编码
	OccupationalNumber *string `json:"occupational_number"`
	// 发证机构名称
	DriverLicenseName *string `json:"driver_license_name"`
	// 证件有效期
	LicenseExpireDate *time.Time `json:"license_expire_date"`
	// 登记时间
	RegistrationTime *time.Time `json:"registration_time"`
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
