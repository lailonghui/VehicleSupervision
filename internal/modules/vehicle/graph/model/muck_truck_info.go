/*
@Time : 2020/12/9 10:47
@Author : lai
@Description :
@File : muck_truck_info
*/
package model

import "time"

// 渣土车信息表
//
//
// columns and relationships of "muck_truck_info"
type MuckTruckInfo struct {
	// 地区ID
	AreaID *int64 `json:"area_id"`
	// 创建时间
	CreatedAt time.Time `json:"create_at"`
	// 创建人
	CreatedBy string `json:"create_by"`
	// 删除时间
	DeletedAt *time.Time `json:"delete_at"`
	// 删除人
	DeletedBy *string `json:"delete_by"`
	// 费用到期时间
	ExpiryDate *time.Time `json:"expiry_date"`
	// 伪IP
	FakeIP *string `json:"fake_ip"`
	// 是否预备库
	IsReserveLibrary *bool `json:"is_reserve_library"`
	// 装载类别
	LoadCategory *string `json:"load_category"`
	// 移动办卡地
	MobileCardLocation *string `json:"mobile_card_location"`
	// 主键
	MuckTruckID int64 `json:"muck_truck_id"`
	// 车辆类型（1.渣土车 2.混凝土车 3.砂石车）
	MuckTruckType *int `json:"muck_truck_type"`
	// 注册日期
	RegistrationDate *time.Time `json:"registration_date"`
	// 自编号
	SelfNumber *string `json:"self_number"`
	// 服务器ID??
	ServerID *int64 `json:"server_id"`
	// SIM卡号
	SimCardNumber *string `json:"sim_card_number"`
	// 修改时间
	UpdatedAt *time.Time `json:"update_at"`
	// 修改人
	UpdatedBy *string `json:"update_by"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
}
