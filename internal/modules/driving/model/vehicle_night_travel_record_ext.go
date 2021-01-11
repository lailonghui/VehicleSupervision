package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleNightTravelRecordUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.VehicleNightTravelRecord

// 数据库表名
func (t *VehicleNightTravelRecord) TableName() string {
	return "vehicle_night_travel_record"
}

// 主键列名
func (t *VehicleNightTravelRecord) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *VehicleNightTravelRecord) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *VehicleNightTravelRecordPkLoader) NewLoader() *VehicleNightTravelRecordPkLoader {
	return &VehicleNightTravelRecordPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleNightTravelRecord, []error) {
			var rs []*VehicleNightTravelRecord
			var m VehicleNightTravelRecord
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *VehicleNightTravelRecord) UnionPrimaryColumnName() string {
	return "record_id"
}

// 获取联合主键
func (t *VehicleNightTravelRecord) GetUnionPrimary() string {
	return t.RecordID
}

// 新建联合主键dataloader
func (t *VehicleNightTravelRecordUnionPkLoader) NewLoader() *VehicleNightTravelRecordUnionPkLoader {
	return &VehicleNightTravelRecordUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleNightTravelRecord, []error) {
			var rs []*VehicleNightTravelRecord
			var m VehicleNightTravelRecord
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
