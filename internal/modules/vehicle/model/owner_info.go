// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 车主信息表
type OwnerInfo struct {
	// 按指定方法生成                                          ( 主键                                 )
	ID int64 `json:"id"`
	// 车主信息外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                             )
	OwnerID string `json:"owner_id"`
	// 所在部门id                                              ( department 部门信息表            )
	DepartmentID *string `json:"department_id"`
	// 车主姓名
	Name *string `json:"name"`
	// 联系地址
	Address *string `json:"address"`
	// 固定电话
	Cellphone *string `json:"cellphone"`
	// 手机号码
	Telephone *string `json:"telephone"`
	// 邮箱地址
	Email *string `json:"email"`
	// 证件过期日期
	ExpiryDate *time.Time `json:"expiry_date"`
	// 身份证号
	IDNumber *string `json:"id_number"`
	// 备注
	Remarks *string `json:"remarks"`
	// 车主性别                                                ( 性别字典                         )
	Sex *int `json:"sex"`
	// 代理商                                                  ( enterprise_info表的enterprise_id )
	Agent *string `json:"agent"`
	// 运营商                                                  ( enterprise_info表的enterprise_id )
	Operator *string `json:"operator"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                                  ( system_user表的user_id           )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                                  ( system_user表的user_id           )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                                  ( system_user表的user_id           )
	DeletedBy *string `json:"deleted_by"`
}