package model

import "time"

// 数据字典类型
//
//
// columns and relationships of "data_dictionary_category"
type DataDictionaryCategory struct {
	// 类型编号
	CategoryCode string `json:"category_code"`
	// 类型名称
	CategoryName string `json:"category_name"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 字典类型ID
	DictionaryCategoryID string `json:"dictionary_category_id"`
	// ID
	ID int64 `json:"id"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
	// 备注
	Remarks *string `json:"remarks"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
}

func (d DataDictionaryCategory) TableName() string {
	return "data_dictionary_category"
}
