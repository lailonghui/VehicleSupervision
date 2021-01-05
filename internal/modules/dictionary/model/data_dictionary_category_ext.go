package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DataDictionaryCategoryUnionPkLoader string *VehicleSupervision/internal/modules/dictionary/model.DataDictionaryCategory

// 数据库表名
func (t DataDictionaryCategory) TableName() string {
	return "data_dictionary_category"
}

// 主键列名
func (t DataDictionaryCategory) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *DataDictionaryCategory) NewPkLoader() *DataDictionaryCategoryPkLoader {
	return &DataDictionaryCategoryPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionaryCategory, []error) {
			var rs []*DataDictionaryCategory
			db.DB.Model(&DataDictionaryCategory{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t DataDictionaryCategory) UnionPrimaryColumnName() string {
	return "dictionary_category_id"
}

// 新建联合主键dataloader
func (t *DataDictionaryCategory) NewUnionPkLoader() *DataDictionaryCategoryUnionPkLoader {
	return &DataDictionaryCategoryUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionaryCategory, []error) {
			var rs []*DataDictionaryCategory
			db.DB.Model(&DataDictionaryCategory{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
