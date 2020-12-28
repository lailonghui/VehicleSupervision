package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleViolationScoringRecordLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.VehicleViolationScoringRecord

func (t VehicleViolationScoringRecord) TableName() string {
	return "vehicle_violation_scoring_record"
}

func (t *VehicleViolationScoringRecord) NewLoader() *VehicleViolationScoringRecordLoader {
	return &VehicleViolationScoringRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleViolationScoringRecord, []error) {
			var rs []*VehicleViolationScoringRecord
			// TODO 按实际需要实现
			db.DB.Model(&VehicleViolationScoringRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
