/*
@Time : 2020/12/9 10:48
@Author : lai
@Description :
@File : muck_truck_preview_number
*/
package model

import "time"

// 渣土车车辆预编号表
//
//
// columns and relationships of "muck_truck_preview_number"
type MuckTruckPreviewNumber struct {
	// 确认状态
	ConfirmStatus *int `json:"confirm_status"`
	// 确认时间
	ConfirmTime *time.Time `json:"confirm_time"`
	// 确认人
	Confirmor *string `json:"confirmor"`
	// 联系人
	ContactPerson *string `json:"contact_person"`
	// 联系电话
	ContactPhone *string `json:"contact_phone"`
	// 创建人
	CreatedBy string `json:"create_by"`
	// 删除时间
	DeletedAt *time.Time `json:"delete_at"`
	// 删除人
	DeletedBy *string `json:"delete_by"`
	// 前车牌
	FrontLicensePlate *string `json:"front_license_plate"`
	// 主键
	ID int64 `json:"id"`
	// 初次登记日期
	InitialRegistrationDate *time.Time `json:"initial_registration_date"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 是否登记销售订单
	IsRegisterSaleOrder *bool `json:"is_register_sale_order"`
	// 是否自动审核
	IsReviewAutomatically *bool `json:"is_review_automatically"`
	// 制作中时间
	MarkingTime *time.Time `json:"marking_time"`
	// 原编号
	OriginalNumber *string `json:"original_number"`
	// 制作状态
	ProductionStatus *int `json:"production_status"`
	// 制作时间
	ProductionTime *time.Time `json:"production_time"`
	// 制作次数
	ProductionTimes *int `json:"production_times"`
	// 后车牌
	RearLicensePlate *string `json:"rear_license_plate"`
	// 登记时间
	RegistrationTime *time.Time `json:"registration_time"`
	// 备注
	Remarks *string `json:"remarks"`
	// 侧车牌
	SideLicensePlate *string `json:"side_license_plate"`
	// 预编号
	SvnNumber *string `json:"svn_number"`
	// 违法未处理数
	UnlawfulViolationNumber *int `json:"unlawful_violation_number"`
	// 修改时间
	UpdatedAt time.Time `json:"update_at"`
	// 修改人
	UpdatedBy *string `json:"update_by"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 工号牌颜色（green.绿色 yellow.黄色）
	WorkNumberPlateColor *string `json:"work_number_plate_color"`
}
