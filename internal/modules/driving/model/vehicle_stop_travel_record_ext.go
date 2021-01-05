package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleStopTravelRecordUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.VehicleStopTravelRecord

// 数据库表名
func (t VehicleStopTravelRecord) TableName() string {
	return "vehicle_stop_travel_record"
}

// 主键列名
func (t VehicleStopTravelRecord) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *VehicleStopTravelRecord) NewPkLoader() *VehicleStopTravelRecordPkLoader {
	return &VehicleStopTravelRecordPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleStopTravelRecord, []error) {
			var rs []*VehicleStopTravelRecord
			db.DB.Model(&VehicleStopTravelRecord{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t VehicleStopTravelRecord) UnionPrimaryColumnName() string {
	return "record_id"
}

// 新建联合主键dataloader
func (t *VehicleStopTravelRecord) NewUnionPkLoader() *VehicleStopTravelRecordUnionPkLoader {
	return &VehicleStopTravelRecordUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleStopTravelRecord, []error) {
			var rs []*VehicleStopTravelRecord
			db.DB.Model(&VehicleStopTravelRecord{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
