// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 营运车信息扩展表
type OperatingVehicleExt struct {
	// 按指定方法生成                           ( 主键                       )
	ID int64 `json:"id"`
	// vehicle_info 车辆信息表 的vehicle_id
	VehicleID string `json:"vehicle_id"`
	// 代理商
	Agent *string `json:"agent"`
	// 合同编号
	ContactNumber *string `json:"contact_number"`
	// 平台                                     ( 平台标识字典           )
	Platform *int `json:"platform"`
	// 行驶证所有人
	DrivingLicenseOwner *string `json:"driving_license_owner"`
	// 速度模式状态获得时间
	SpeedModeStatusTime *time.Time `json:"speed_mode_status_time"`
	// 速度模式状态
	SpeedModeStatus *string `json:"speed_mode_status"`
	// 报废原因
	ScrapReason *string `json:"scrap_reason"`
	// 报废时间
	ScrapTime *time.Time `json:"scrap_time"`
	// 报废时间审核
	ScrapTimeCheck *int `json:"scrap_time_check"`
	// 行政区域
	AdministrativeRegion *string `json:"administrative_region"`
	// 车牌照片,云储存系统返回的路径
	LicensePlatePhoto *string `json:"license_plate_photo"`
	// 其他照片,云储存系统返回的路径
	OtherPhoto *string `json:"other_photo"`
	// 编号
	SerialNumber *string `json:"serial_number"`
	// 是否北斗部标平台
	IsBeidou *bool `json:"is_beidou"`
	// 是否在运证系统
	IsInOperatingSystem *bool `json:"is_in_operating_system"`
	// 是否上传货运平台
	IsInUploadPlatform *bool `json:"is_in_upload_platform"`
	// 是否监管
	IsSupervise *bool `json:"is_supervise"`
	// 是否需要监管
	IsNeedSupervise *bool `json:"is_need_supervise"`
	// 是否功能测试通过
	IsFunctionOk *bool `json:"is_function_ok"`
	// 是否屏蔽
	IsBlock *bool `json:"is_block"`
	// 是否办理终端安装证明
	IsApplayTerminalInstallation *bool `json:"is_applay_terminal_installation"`
	// 第一次上线时间
	FirstOnlineTime *time.Time `json:"first_online_time"`
	// 最后一次绑定终端时间
	LastBindingTerminalTime *time.Time `json:"last_binding_terminal_time"`
	// 服务到期时间
	ServiceExpirationTime *time.Time `json:"service_expiration_time"`
	// 合同时间
	ContractTime *time.Time `json:"contract_time"`
	// 安装时间
	InstallationTime *time.Time `json:"installation_time"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                   ( system_user表的user_id )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                   ( system_user表的user_id )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                   ( system_user表的user_id )
	DeletedBy *string `json:"deleted_by"`
}