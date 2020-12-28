package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DisputeViolationRecordLogLoader string *VehicleSupervision/internal/modules/vehicle_driver_separation/model.DisputeViolationRecordLog

func (t DisputeViolationRecordLog) TableName() string {
	return "dispute_violation_record_log"
}

func (t *DisputeViolationRecordLog) NewLoader() *DisputeViolationRecordLogLoader {
	return &DisputeViolationRecordLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DisputeViolationRecordLog, []error) {
			var rs []*DisputeViolationRecordLog
			// TODO 按实际需要实现
			db.DB.Model(&DisputeViolationRecordLog{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
