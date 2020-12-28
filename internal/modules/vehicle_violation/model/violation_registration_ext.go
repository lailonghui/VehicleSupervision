package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden ViolationRegistrationLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.ViolationRegistration

func (t ViolationRegistration) TableName() string {
	return "ViolationRegistration"
}

func (t *ViolationRegistration) NewLoader() *ViolationRegistrationLoader {
	return &ViolationRegistrationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ViolationRegistration, []error) {
			var rs []*ViolationRegistration
			// TODO 按实际需要实现
			db.DB.Model(&ViolationRegistration{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
