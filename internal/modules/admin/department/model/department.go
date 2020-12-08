package model

import "time"

// 部门
//
//
// columns and relationships of "department"
type Department struct {
	// 创建时间
	CreateAt *time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 部门类型
	DepartmentCategory *int `json:"department_category"`
	// 部门编号
	DepartmentCode *string `json:"department_code"`
	// 部门ID
	DepartmentID string `json:"department_id"`
	// 部门名称
	DepartmentName *string `json:"department_name"`
	// 企业ID
	EnterpriseID string `json:"enterprise_id"`
	// ID
	ID int64 `json:"id"`
	// 排序
	InternalNumber *int `json:"internal_number"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
	// 上级部门ID
	SuperiorDepartmentID *string `json:"superior_department_id"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
}
