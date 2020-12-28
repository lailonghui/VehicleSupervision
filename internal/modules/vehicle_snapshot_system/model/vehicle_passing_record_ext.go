package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehiclePassingRecordLoader string *VehicleSupervision/internal/modules/vehicle_snapshot_system/model.VehiclePassingRecord

func (t VehiclePassingRecord) TableName() string {
	return "VehiclePassingRecord"
}

func (t *VehiclePassingRecord) NewLoader() *VehiclePassingRecordLoader {
	return &VehiclePassingRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehiclePassingRecord, []error) {
			var rs []*VehiclePassingRecord
			// TODO 按实际需要实现
			db.DB.Model(&VehiclePassingRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
