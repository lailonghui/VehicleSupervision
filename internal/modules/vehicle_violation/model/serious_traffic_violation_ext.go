package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden SeriousTrafficViolationLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.SeriousTrafficViolation

func (t SeriousTrafficViolation) TableName() string {
	return "SeriousTrafficViolation"
}

func (t *SeriousTrafficViolation) NewLoader() *SeriousTrafficViolationLoader {
	return &SeriousTrafficViolationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SeriousTrafficViolation, []error) {
			var rs []*SeriousTrafficViolation
			// TODO 按实际需要实现
			db.DB.Model(&SeriousTrafficViolation{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
