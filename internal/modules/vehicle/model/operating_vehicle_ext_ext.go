package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden OperatingVehicleExtLoader string *VehicleSupervision/internal/modules/vehicle/model.OperatingVehicleExt

func (t OperatingVehicleExt) TableName() string {
	return "OperatingVehicleExt"
}

func (t *OperatingVehicleExt) NewLoader() *OperatingVehicleExtLoader {
	return &OperatingVehicleExtLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*OperatingVehicleExt, []error) {
			var rs []*OperatingVehicleExt
			// TODO 按实际需要实现
			db.DB.Model(&OperatingVehicleExt{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
