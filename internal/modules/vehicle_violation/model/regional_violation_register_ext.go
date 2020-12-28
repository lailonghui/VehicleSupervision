package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden RegionalViolationRegisterLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.RegionalViolationRegister

func (t RegionalViolationRegister) TableName() string {
	return "regional_violation_register"
}

func (t *RegionalViolationRegister) NewLoader() *RegionalViolationRegisterLoader {
	return &RegionalViolationRegisterLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RegionalViolationRegister, []error) {
			var rs []*RegionalViolationRegister
			// TODO 按实际需要实现
			db.DB.Model(&RegionalViolationRegister{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
