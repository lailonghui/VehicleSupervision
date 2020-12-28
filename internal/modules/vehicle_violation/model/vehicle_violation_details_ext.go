package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleViolationDetailsLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.VehicleViolationDetails

func (t VehicleViolationDetails) TableName() string {
	return "vehicle_violation_details"
}

func (t *VehicleViolationDetails) NewLoader() *VehicleViolationDetailsLoader {
	return &VehicleViolationDetailsLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleViolationDetails, []error) {
			var rs []*VehicleViolationDetails
			// TODO 按实际需要实现
			db.DB.Model(&VehicleViolationDetails{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
