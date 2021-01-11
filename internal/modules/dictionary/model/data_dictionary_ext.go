package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DataDictionaryUnionPkLoader string *VehicleSupervision/internal/modules/dictionary/model.DataDictionary

// 数据库表名
func (t *DataDictionary) TableName() string {
	return "data_dictionary"
}

// 主键列名
func (t *DataDictionary) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *DataDictionary) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *DataDictionaryPkLoader) NewLoader() *DataDictionaryPkLoader {
	return &DataDictionaryPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionary, []error) {
			var rs []*DataDictionary
			var m DataDictionary
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *DataDictionary) UnionPrimaryColumnName() string {
	return "dictionary_id"
}

// 获取联合主键
func (t *DataDictionary) GetUnionPrimary() string {
	return t.DictionaryID
}

// 新建联合主键dataloader
func (t *DataDictionaryUnionPkLoader) NewLoader() *DataDictionaryUnionPkLoader {
	return &DataDictionaryUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionary, []error) {
			var rs []*DataDictionary
			var m DataDictionary
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
