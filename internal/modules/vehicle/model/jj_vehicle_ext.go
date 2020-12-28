package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden JjVehicleLoader string *VehicleSupervision/internal/modules/vehicle/model.JjVehicle

func (t JjVehicle) TableName() string {
	return "JjVehicle"
}

func (t *JjVehicle) NewLoader() *JjVehicleLoader {
	return &JjVehicleLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*JjVehicle, []error) {
			var rs []*JjVehicle
			// TODO 按实际需要实现
			db.DB.Model(&JjVehicle{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
