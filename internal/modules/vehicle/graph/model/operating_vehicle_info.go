/*
@Time : 2020/12/9 10:46
@Author : lai
@Description :
@File : operating_vehicle_info
*/
package model

import "time"

// 营运车信息表
//
//
// columns and relationships of "operating_vehicle_info"
type OperatingVehicleInfo struct {
	// 行政区域
	AdministrativeRegion *string `json:"administrative_region"`
	// 代理商
	Agent *string `json:"agent"`
	// 合同编号
	ContactNumber *string `json:"contact_number"`
	// 合同时间
	ContractTime *time.Time `json:"contract_time"`
	// 创建时间
	CreatedAt time.Time `json:"create_at"`
	// 创建人
	CreatedBy string `json:"create_by"`
	// 删除时间
	DeletedAt *time.Time `json:"delete_at"`
	// 删除人
	DeletedBy *string `json:"delete_by"`
	// 行驶证所有人
	DrivingLicenseOwner *string `json:"driving_license_owner"`
	// 第一次上线时间
	FirstOnlineTime *time.Time `json:"first_online_time"`
	// 安装时间
	InstallationTime *time.Time `json:"installation_time"`
	// 是否北斗部标平台
	IsBd *bool `json:"is_BD"`
	// 是否激活
	IsActive *bool `json:"is_active"`
	// 是否办理终端安装证明
	IsApplayTerminalInstallation *bool `json:"is_applay_terminal_installation"`
	// 是否屏蔽
	IsBlock *bool `json:"is_block"`
	// 是否工程运输车
	IsEngineeringVehicle *bool `json:"is_engineering_vehicle"`
	// 是否功能测试通过
	IsFunctionOk *bool `json:"is_function_ok"`
	// 是否在运证系统
	IsInOperatingSystem *bool `json:"is_in_operating_system"`
	// 是否上传货运平台
	IsInUploadPlatform *bool `json:"is_in_upload_platform"`
	// 是否需要监管
	IsNeedSupervise *bool `json:"is_need_supervise"`
	// 是否监管
	IsSupervise *bool `json:"is_supervise"`
	// 最后一次绑定终端时间
	LastBindingTerminalTime *time.Time `json:"last_binding_terminal_time"`
	// 车牌照片,云储存系统返回的路径
	LicensePlatePhoto *string `json:"license_plate_photo"`
	// 渣土车类型,工程运输车车辆类型字典
	MuckTruckType *int `json:"muck_truck_type"`
	// 主键
	OperatingVehicleID int64 `json:"operating_vehicle_id"`
	// 其他照片,云储存系统返回的路径
	OtherPhoto *string `json:"other_photo"`
	// 平台标识字典
	Platform *int `json:"platform"`
	// 位置数据库ID
	PositionDbID *string `json:"position_db_id"`
	// 报废原因
	ScrapReason *string `json:"scrap_reason"`
	// 报废时间
	ScrapTime *time.Time `json:"scrap_time"`
	// 报废时间审核
	ScrapTimeCheck *int `json:"scrap_time_check"`
	// 编号
	SerialNumber *string `json:"serial_number"`
	// 服务到期时间
	ServiceExpirationTime *time.Time `json:"service_expiration_time"`
	// 速度模式状态
	SpeedModeStatus *string `json:"speed_mode_status"`
	// 速度模式状态获得时间
	SpeedModeStatusTime *time.Time `json:"speed_mode_status_time"`
	// 临时库,用法不明？
	TemporaryLibrary *int `json:"temporary_library"`
	// 临时传交通局,用法不明
	TemporaryTransportBureau *int `json:"temporary_transport_bureau"`
	// 修改时间
	UpdatedAt *time.Time `json:"update_at"`
	// 修改人
	UpdatedBy *string `json:"update_by"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
}

func (OperatingVehicleInfo) TableName() string {
	return "operating_vehicle_info"
}
