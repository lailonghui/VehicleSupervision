package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ElectricFenceEnteranceRecordUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.ElectricFenceEnteranceRecord

// 数据库表名
func (t ElectricFenceEnteranceRecord) TableName() string {
	return "electric_fence_enterance_record"
}

// 主键列名
func (t ElectricFenceEnteranceRecord) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *ElectricFenceEnteranceRecord) NewPkLoader() *ElectricFenceEnteranceRecordPkLoader {
	return &ElectricFenceEnteranceRecordPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFenceEnteranceRecord, []error) {
			var rs []*ElectricFenceEnteranceRecord
			db.DB.Model(&ElectricFenceEnteranceRecord{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t ElectricFenceEnteranceRecord) UnionPrimaryColumnName() string {
	return "record_id"
}

// 新建联合主键dataloader
func (t *ElectricFenceEnteranceRecord) NewUnionPkLoader() *ElectricFenceEnteranceRecordUnionPkLoader {
	return &ElectricFenceEnteranceRecordUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFenceEnteranceRecord, []error) {
			var rs []*ElectricFenceEnteranceRecord
			db.DB.Model(&ElectricFenceEnteranceRecord{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
