package model

import "time"

// 网约车驾驶员
//
//
// columns and relationships of "ride_hailing_driver"
type RideHailingDriver struct {
	// 出生日期
	Birthday *time.Time `json:"birthday"`
	// 审核状态
	CheckStation *int `json:"check_station"`
	// 记录时间
	CreateAt *time.Time `json:"create_at"`
	// 记录人员ID
	CreateBy *string `json:"create_by"`
	// 现居住地址
	CurrentAddress *string `json:"current_address"`
	// 驾驶员名称
	DriverName *string `json:"driver_name"`
	// 驾校ID
	DriverSchoolID *string `json:"driver_school_id"`
	// 有效期截止日期
	EndValidDate *time.Time `json:"end_valid_date"`
	// 初次领驾驶证日期
	FirstTimeReceivedDate *time.Time `json:"first_time_received_date"`
	// 手持身份证照片
	HandleIDPhoto *string `json:"handle_id_photo"`
	// ID
	ID int64 `json:"id"`
	// 身份证住址
	IDAddress *string `json:"id_address"`
	// 驾驶员身份证号
	IDNumber *string `json:"id_number"`
	// 身份证图片
	IdcardPhoto *string `json:"idcard_photo"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 是否旧驾驶员
	IsFormerDriver *bool `json:"is_former_driver"`
	// 民族
	Nation *string `json:"nation"`
	// 操作员id
	OperatorID *string `json:"operator_id"`
	// 联系电话
	PhoneNumber *string `json:"phone_number"`
	// 资格证号
	QualificationNumber *string `json:"qualification_number"`
	// 准假车型
	QuasiDrivingModels *int `json:"quasi_driving_models"`
	// 备注
	Remarks *string `json:"remarks"`
	// 网约车驾驶员ID
	RideHailingDriverID string `json:"ride_hailing_driver_id"`
	// 网约车驾驶员审核表ID
	RideHailingDriverVerifyID *string `json:"ride_hailing_driver_verify_id"`
	// 性别
	Sex *int `json:"sex"`
	// 发证机构
	SignGov *string `json:"sign_gov"`
	// 签字的照片
	SignnaturePhoto *string `json:"signnature_photo"`
	// 身份证-有效期起始日期
	StartValidDate *time.Time `json:"start_valid_date"`
	// 更新时间
	UpdateAt *time.Time `json:"update_at"`
	// 更新人员id
	UpdateBy *string `json:"update_by"`
	// 内网更新时间
	UpdateTimeIn *time.Time `json:"update_time_in"`
}

func (r RideHailingDriver) TableName() string {
	return "ride_hailing_driver"
}
