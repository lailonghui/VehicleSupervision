// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// 数据字典
type DataDictionary struct {
	// ID
	ID int64 `json:"id"`
	// 字典ID
	DictionaryID string `json:"dictionary_id"`
	// 字典类型ID
	DictionaryCategoryID string `json:"dictionary_category_id"`
	// 名称
	Name string `json:"name"`
	// 值
	Value int `json:"value"`
	// 备注
	Remarks *string `json:"remarks"`
	// 创建时间
	CreateAt time.Time `json:"create_at"`
	// 创建人
	CreateBy *string `json:"create_by"`
	// 修改时间
	UpdateAt *time.Time `json:"update_at"`
	// 修改人
	UpdateBy *string `json:"update_by"`
	// 删除时间
	DeleteAt *time.Time `json:"delete_at"`
	// 删除人
	DeleteBy *string `json:"delete_by"`
	// 是否删除
	IsDelete bool `json:"is_delete"`
}

// expression to compare columns of type Float. All fields are combined with logical 'AND'.
type FloatComparisonExp struct {
	Eq     *float64  `json:"_eq"`
	Gt     *float64  `json:"_gt"`
	Gte    *float64  `json:"_gte"`
	In     []float64 `json:"_in"`
	IsNull *bool     `json:"_is_null"`
	Lt     *float64  `json:"_lt"`
	Lte    *float64  `json:"_lte"`
	Neq    *float64  `json:"_neq"`
	Nin    []float64 `json:"_nin"`
}
