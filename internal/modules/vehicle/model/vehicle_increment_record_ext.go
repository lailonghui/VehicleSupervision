package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleIncrementRecordLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleIncrementRecord

func (t VehicleIncrementRecord) TableName() string {
	return "vehicle_increment_record"
}

func (t *VehicleIncrementRecord) NewLoader() *VehicleIncrementRecordLoader {
	return &VehicleIncrementRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleIncrementRecord, []error) {
			var rs []*VehicleIncrementRecord
			// TODO 按实际需要实现
			db.DB.Model(&VehicleIncrementRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
