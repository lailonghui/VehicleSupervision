package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DataDictionaryCategoryLoader string *VehicleSupervision/internal/modules/dictionary/model.DataDictionaryCategory

// 数据库表名
func (t DataDictionaryCategory) TableName() string {
	return "data_dictionary_category"
}

// 主键列名
func (t DataDictionaryCategory) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t DataDictionaryCategory) UnionPrimaryColumnName() string {
	return "dictionary_category_id"
}

// 新建dataloader
func (t *DataDictionaryCategory) NewLoader() *DataDictionaryCategoryLoader {
	return &DataDictionaryCategoryLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionaryCategory, []error) {
			var rs []*DataDictionaryCategory

			db.DB.Model(&DataDictionaryCategory{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
