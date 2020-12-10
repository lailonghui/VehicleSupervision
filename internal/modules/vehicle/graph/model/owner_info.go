/*
@Time : 2020/12/9 10:50
@Author : lai
@Description :
@File : owner_info
*/
package model

import "time"

// 车主信息表
//
//
// columns and relationships of "owner_info"
type OwnerInfo struct {
	// 联系地址
	Address *string `json:"address"`
	// 代理商
	Agent *string `json:"agent"`
	// 固定电话
	Cellphone *string `json:"cellphone"`
	// 创建时间
	CreatedAt time.Time `json:"create_at"`
	// 创建人
	CreatedBy string `json:"create_by"`
	// 删除时间
	DeletedAt *time.Time `json:"delete_at"`
	// 删除人
	DeletedBy *string `json:"delete_by"`
	// department 部门信息表
	DepartmentID *string `json:"department_id"`
	// 邮箱地址
	Email *string `json:"email"`
	// 证件过期日期
	ExpiryDate *time.Time `json:"expiry_date"`
	// 主键
	ID int64 `json:"id"`
	// 身份证号
	IDNumber *string `json:"id_number"`
	// 是否删除
	IsDelete *bool `json:"is_delete"`
	// 车主姓名
	Name *string `json:"name"`
	// 运营商
	Operator *string `json:"operator"`
	// 车主信息外部编码，由golang程序生成的xid，暴露到外部使用
	// 联合主键
	OwnerID string `json:"owner_id"`
	// 备注
	Remarks *string `json:"remarks"`
	// 车主性别(性别字典)
	Sex *int `json:"sex"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 修改时间
	UpdatedAt *time.Time `json:"update_at"`
	// 修改人
	UpdatedBy *string `json:"update_by"`
}
