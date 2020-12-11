/*
@Time : 2020/12/11 14:41
@Author : lai
@Description :
@File : enterprise_deduction_operation_record
*/
package model

import "time"

// 企业扣分操作记录表
//
//
// columns and relationships of "enterprise_deduction_operation_record"
type EnterpriseDeductionOperationRecord struct {
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy string `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 扣分分值
	DemeritPoints *float64 `json:"demerit_points"`
	// 企业扣分事项表id(enterprise_deduction_ items表的enterprise_deduction_ item_id)
	EnterpriseDeductionItemID *string `json:"enterprise_deduction_item_id"`
	// 联合主键
	EnterpriseDuductionOperationID string `json:"enterprise_duduction_operation_id"`
	// 所在企业id
	EnterpriseID *string `json:"enterprise_id"`
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
}

func (EnterpriseDeductionOperationRecord) TableName() string {
	return "enterprise_deduction_operation_record"
}
