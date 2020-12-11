/*
@Time : 2020/12/11 14:42
@Author : lai
@Description :
@File : vehicle_violation_scoring_record
*/
package model

import "time"

// 车辆违规计分记录
//
//
// columns and relationships of "vehicle_violation_scoring_record"
type VehicleViolationScoringRecord struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 扣分分值(vehicle_violation_scoring_ items表的violation_scoring_item_id)
	DemeritPoints *float64 `json:"demerit_points"`
	// 主键
	ID int64 `json:"id"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 扣分车辆id
	VehicleID *string `json:"vehicle_id"`
	// 联合主键
	ViolationScoringID string `json:"violation_scoring_id"`
	// 扣分明细id
	ViolationScoringItemID *string `json:"violation_scoring_item_id"`
}

func (VehicleViolationScoringRecord) TableName() string {
	return "vehicle_violation_scoring_record"
}
