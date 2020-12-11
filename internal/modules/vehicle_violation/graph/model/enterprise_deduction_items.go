/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : enterprise_deduction_items
*/
package model

import "time"

// 企业扣分事项表
//
//
// columns and relationships of "enterprise_deduction_items"
type EnterpriseDeductionItems struct {
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
	DemeritPoints *float64 `json:"demerit_points"`
	// 联合主键
	EnterpriseDeductionItemID string `json:"enterprise_deduction_item_id"`
	// 主键
	ID int64 `json:"id"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
}

func (EnterpriseDeductionItems) TableName() string {
	return "enterprise_deduction_items"
}
