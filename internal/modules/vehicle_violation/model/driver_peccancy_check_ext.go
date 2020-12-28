package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DriverPeccancyCheckLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.DriverPeccancyCheck

func (t DriverPeccancyCheck) TableName() string {
	return "DriverPeccancyCheck"
}

func (t *DriverPeccancyCheck) NewLoader() *DriverPeccancyCheckLoader {
	return &DriverPeccancyCheckLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverPeccancyCheck, []error) {
			var rs []*DriverPeccancyCheck
			// TODO 按实际需要实现
			db.DB.Model(&DriverPeccancyCheck{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
