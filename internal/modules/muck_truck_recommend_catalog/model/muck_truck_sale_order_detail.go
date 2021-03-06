// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 新型渣土车销售订单明细
type MuckTruckSaleOrderDetail struct {
	// 按指定方法生成                                  ( 主键                                                         )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                                                     )
	MuckTruckSaleOrderDetailID string `json:"muck_truck_sale_order_detail_id"`
	// 订单号                                          ( muck_truck_sale_order 新型渣土车销售订单的muck_truck_sale_order_id )
	OrderID *string `json:"order_id"`
	// 车架号(后6位)                                   ( 车辆识别代号vin,如D02133                                     )
	VehicleIdentificationNumber *string `json:"vehicle_identification_number"`
	// 合格证
	Certificate *string `json:"certificate"`
	// 终端证明
	TerminalProof *string `json:"terminal_proof"`
	// 行驶证(照片)
	DriverLicensePic *string `json:"driver_license_pic"`
	// 车牌号码
	LicensePlateNumber *string `json:"license_plate_number"`
	// 车牌颜色                                        ( 车牌颜色字典                                             )
	LicensePlateColor *int `json:"license_plate_color"`
	// 号牌种类                                        ( 号牌种类字典                                             )
	LicensePlateType *int `json:"license_plate_type"`
	// 抵达泉州日期
	ArriveQzDate *time.Time `json:"arrive_qz_date"`
	// 初次登记日期
	FirstRegistrationDate *time.Time `json:"first_registration_date"`
	// 登记时间
	RegistrationDate *time.Time `json:"registration_date"`
	// 步骤(销售订单登记进度（2.到车登记 3.上牌登记）)
	Step *int `json:"step"`
	// 车辆照片
	VehiclePhoto *string `json:"vehicle_photo"`
	// 供应商销售预编号
	SellerPreviewNumber *string `json:"seller_preview_number"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id                                   )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id                                   )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id                                   )
	DeletedBy *string `json:"deleted_by"`
}
