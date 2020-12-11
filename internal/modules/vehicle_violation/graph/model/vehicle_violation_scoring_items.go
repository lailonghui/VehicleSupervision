/*
@Time : 2020/12/11 14:45
@Author : lai
@Description :
@File : vehicle_violation_scoring_items
*/
package model

import "time"

// 车辆违规计分项表
//
//
// columns and relationships of "vehicle_violation_scoring_items"
type VehicleViolationScoringItems struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 扣分事项类别(车辆评分扣分类别字典)
	DeductionCategory *int `json:"deduction_category"`
	// 扣分事项描述
	DeductionItemDescription *string `json:"deduction_item_description"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 扣分分值
	DemeritPoints *string `json:"demerit_points"`
	// 主键
	ID int64 `json:"id"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 联合主键
	ViolationScoringItemID string `json:"violation_scoring_item_id"`
}

func (VehicleViolationScoringItems) TableName() string {
	return "vehicle_violation_scoring_items"
}
