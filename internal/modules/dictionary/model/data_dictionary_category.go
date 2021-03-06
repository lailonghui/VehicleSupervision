// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 数据字典类型
type DataDictionaryCategory struct {
	// ID
	ID int64 `json:"id"`
	// 字典类型ID
	DictionaryCategoryID string `json:"dictionary_category_id"`
	// 类型名称
	CategoryName string `json:"category_name"`
	// 类型编号
	CategoryCode string `json:"category_code"`
	// 备注
	Remarks *string `json:"remarks"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建人
	CreatedBy *string `json:"created_by"`
	// 修改时间
	UpdatedAt *time.Time `json:"updated_at"`
	// 修改人
	UpdatedBy *string `json:"updated_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 删除人
	DeletedBy *string `json:"deleted_by"`
	// 是否删除
	IsDeleted bool `json:"is_deleted"`
}
