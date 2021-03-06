// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 新型渣土车信息表
type NewMuckTruckInfo struct {
	// 按指定方法生成                                  ( 主键                           )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                       )
	NewMuckTruckInfoID string `json:"new_muck_truck_info_id"`
	// 扣分车辆id                                      ( vehicle_info表的vehicle_id )
	VehicleID *string `json:"vehicle_id"`
	// 车牌号
	LicensePlateNumber *string `json:"license_plate_number"`
	// 登记注册时间
	RegistrationTime *time.Time `json:"registration_time"`
	// 排放标准
	EmissionStandard *string `json:"emission_standard"`
	// 是否U型货箱
	IsUShapedCargoBox *bool `json:"is_u_shaped_cargo_box"`
	// 长
	Length *string `json:"length"`
	// 宽
	Width *string `json:"width"`
	// 高
	Height *string `json:"height"`
	// 密封设备
	SealingDevice *string `json:"sealing_device"`
	// 顶盖高度
	TopCoverHeight *string `json:"top_cover_height"`
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
