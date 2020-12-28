package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DisputeViolationRecordLoader string *VehicleSupervision/internal/modules/vehicle_driver_separation/model.DisputeViolationRecord

func (t DisputeViolationRecord) TableName() string {
	return "dispute_violation_record"
}

func (t *DisputeViolationRecord) NewLoader() *DisputeViolationRecordLoader {
	return &DisputeViolationRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DisputeViolationRecord, []error) {
			var rs []*DisputeViolationRecord
			// TODO 按实际需要实现
			db.DB.Model(&DisputeViolationRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
