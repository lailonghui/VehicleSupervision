/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : regional_violation_register
*/
package model

import "time"

// 区域处理机关交通违法违规登记表
//
//
// columns and relationships of "regional_violation_register"
type RegionalViolationRegister struct {
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
	// 违法代码(VIO_CODEWFDM 违法描述字典表)
	IllegalCode *int `json:"illegal_code"`
	// 违法时间
	IllegalTime *time.Time `json:"illegal_time"`
	// 类型(1车辆2驾驶员),违法类型字典表
	IllegalType *int `json:"illegal_type"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 类型(false未登记true已登记)
	IsRegister *bool `json:"is_register"`
	// 操作员
	Operator *string `json:"operator"`
	// 处理机关
	ProcessingAgency *string `json:"processing_agency"`
	// 联合主键
	RegionalViolationRegisterID string `json:"regional_violation_register_id"`
	// 登记时间
	RegisterTime *time.Time `json:"register_time"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 违章车辆id
	VehicleID *string `json:"vehicle_id"`
	// 违法记录表ID(vehicle_violation_details的violation_detail_id)
	ViolationDetailID *string `json:"violation_detail_id"`
}

func (RegionalViolationRegister) TableName() string {
	return "regional_violation_register"
}
