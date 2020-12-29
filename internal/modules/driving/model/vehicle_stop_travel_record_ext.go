package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleStopTravelRecordLoader string *VehicleSupervision/internal/modules/driving/model.VehicleStopTravelRecord

// 数据库表名
func (t VehicleStopTravelRecord) TableName() string {
	return "vehicle_stop_travel_record"
}

// 主键列名
func (t VehicleStopTravelRecord) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t VehicleStopTravelRecord) UnionPrimaryColumnName() string {
	return "record_id"
}

// 新建dataloader
func (t *VehicleStopTravelRecord) NewLoader() *VehicleStopTravelRecordLoader {
	return &VehicleStopTravelRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleStopTravelRecord, []error) {
			var rs []*VehicleStopTravelRecord

			db.DB.Model(&VehicleStopTravelRecord{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
