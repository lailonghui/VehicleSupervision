package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DataDictionaryCategoryUnionPkLoader string *VehicleSupervision/internal/modules/dictionary/model.DataDictionaryCategory

// 数据库表名
func (t *DataDictionaryCategory) TableName() string {
	return "data_dictionary_category"
}

// 主键列名
func (t *DataDictionaryCategory) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *DataDictionaryCategory) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *DataDictionaryCategoryPkLoader) NewLoader() *DataDictionaryCategoryPkLoader {
	return &DataDictionaryCategoryPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionaryCategory, []error) {
			var rs []*DataDictionaryCategory
			var m DataDictionaryCategory
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *DataDictionaryCategory) UnionPrimaryColumnName() string {
	return "dictionary_category_id"
}

// 获取联合主键
func (t *DataDictionaryCategory) GetUnionPrimary() string {
	return t.DictionaryCategoryID
}

// 新建联合主键dataloader
func (t *DataDictionaryCategoryUnionPkLoader) NewLoader() *DataDictionaryCategoryUnionPkLoader {
	return &DataDictionaryCategoryUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionaryCategory, []error) {
			var rs []*DataDictionaryCategory
			var m DataDictionaryCategory
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
