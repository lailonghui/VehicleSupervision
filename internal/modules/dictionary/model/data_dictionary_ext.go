package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DataDictionaryLoader string *VehicleSupervision/internal/modules/dictionary/model.DataDictionary

// 数据库表名
func (t DataDictionary) TableName() string {
	return "data_dictionary"
}

// 主键列名
func (t DataDictionary) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t DataDictionary) UnionPrimaryColumnName() string {
	return "dictionary_id"
}

// 新建dataloader
func (t *DataDictionary) NewLoader() *DataDictionaryLoader {
	return &DataDictionaryLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DataDictionary, []error) {
			var rs []*DataDictionary

			db.DB.Model(&DataDictionary{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
