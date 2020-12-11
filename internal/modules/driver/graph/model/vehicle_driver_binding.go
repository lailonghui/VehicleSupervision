/*
@Time : 2020/12/11 10:54
@Author : lai
@Description :
@File : vehicle_driver_binding
*/
package model

import "time"

// 车辆驾驶员绑定表
//
//
// columns and relationships of "vehicle_driver_binding"
type VehicleDriverBinding struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 驾驶员id
	DriverID *string `json:"driver_id"`
	// 主键
	ID int64 `json:"id"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 租赁合同,,云储存系统返回的完整租赁合同的图片路径
	LeaseContract *string `json:"lease_contract"`
	// 备注
	Remarks *string `json:"remarks"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 车辆驾驶员绑定外部编码，由golang程序生成的xid，暴露到外部使用
	VehicleDriverBindingID string `json:"vehicle_driver_binding_id"`
	// 车辆id
	VehicleID *string `json:"vehicle_id"`
}

func (VehicleDriverBinding) TableName() string {
	return "vehicle_driver_binding"
}
