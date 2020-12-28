// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 销售商评分记录
type SellerRatingRecord struct {
	// 按指定方法生成                                  ( 主键                                 )
	ID int64 `json:"id"`
	// 外部编码，由golang程序生成的xid，暴露到外部使用 ( 联合主键                             )
	SellerRatingRecordID string `json:"seller_rating_record_id"`
	// 扣分对象                                        ( enterprise_info表的enterprise_id )
	DemeritObj *string `json:"demerit_obj"`
	// 扣分分值
	DemeritPoints *float64 `json:"demerit_points"`
	// 扣分原因
	DemeritReason *string `json:"demerit_reason"`
	// 操作人                                          ( system_user表的user_id           )
	Operator *string `json:"operator"`
	// 操作时间
	OperationTime *time.Time `json:"operation_time"`
	// 是否删除
	IsDeleted *bool `json:"is_deleted"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人                                          ( system_user表的user_id           )
	CreatedBy string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人                                          ( system_user表的user_id           )
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人                                          ( system_user表的user_id           )
	DeletedBy *string `json:"deleted_by"`
}
