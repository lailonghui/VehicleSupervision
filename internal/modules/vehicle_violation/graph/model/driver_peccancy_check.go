/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : driver_peccancy_check
*/
package model

import "time"

// 驾驶员违法核实记录表
//
//
// columns and relationships of "driver_peccancy_check"
type DriverPeccancyCheck struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 违章驾驶员id
	DriverID *string `json:"driver_id"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
	// 主键
	ID int64 `json:"id"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 违章车辆id
	VehicleID *string `json:"vehicle_id"`
}

func (DriverPeccancyCheck) TableName() string {
	return "driver_peccancy_check"
}
