package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ElectricFenceEnteranceRecordUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.ElectricFenceEnteranceRecord

// 数据库表名
func (t *ElectricFenceEnteranceRecord) TableName() string {
	return "electric_fence_enterance_record"
}

// 主键列名
func (t *ElectricFenceEnteranceRecord) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *ElectricFenceEnteranceRecord) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *ElectricFenceEnteranceRecordPkLoader) NewLoader() *ElectricFenceEnteranceRecordPkLoader {
	return &ElectricFenceEnteranceRecordPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFenceEnteranceRecord, []error) {
			var rs []*ElectricFenceEnteranceRecord
			var m ElectricFenceEnteranceRecord
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *ElectricFenceEnteranceRecord) UnionPrimaryColumnName() string {
	return "record_id"
}

// 获取联合主键
func (t *ElectricFenceEnteranceRecord) GetUnionPrimary() string {
	return t.RecordID
}

// 新建联合主键dataloader
func (t *ElectricFenceEnteranceRecordUnionPkLoader) NewLoader() *ElectricFenceEnteranceRecordUnionPkLoader {
	return &ElectricFenceEnteranceRecordUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFenceEnteranceRecord, []error) {
			var rs []*ElectricFenceEnteranceRecord
			var m ElectricFenceEnteranceRecord
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
