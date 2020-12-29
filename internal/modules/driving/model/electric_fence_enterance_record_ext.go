package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ElectricFenceEnteranceRecordLoader string *VehicleSupervision/internal/modules/driving/model.ElectricFenceEnteranceRecord

// 数据库表名
func (t ElectricFenceEnteranceRecord) TableName() string {
	return "electric_fence_enterance_record"
}

// 主键列名
func (t ElectricFenceEnteranceRecord) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t ElectricFenceEnteranceRecord) UnionPrimaryColumnName() string {
	return "record_id"
}

// 新建dataloader
func (t *ElectricFenceEnteranceRecord) NewLoader() *ElectricFenceEnteranceRecordLoader {
	return &ElectricFenceEnteranceRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFenceEnteranceRecord, []error) {
			var rs []*ElectricFenceEnteranceRecord

			db.DB.Model(&ElectricFenceEnteranceRecord{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
