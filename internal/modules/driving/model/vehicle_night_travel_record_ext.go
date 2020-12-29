package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleNightTravelRecordLoader string *VehicleSupervision/internal/modules/driving/model.VehicleNightTravelRecord

// 数据库表名
func (t VehicleNightTravelRecord) TableName() string {
	return "vehicle_night_travel_record"
}

// 主键列名
func (t VehicleNightTravelRecord) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t VehicleNightTravelRecord) UnionPrimaryColumnName() string {
	return "record_id"
}

// 新建dataloader
func (t *VehicleNightTravelRecord) NewLoader() *VehicleNightTravelRecordLoader {
	return &VehicleNightTravelRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleNightTravelRecord, []error) {
			var rs []*VehicleNightTravelRecord

			db.DB.Model(&VehicleNightTravelRecord{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
