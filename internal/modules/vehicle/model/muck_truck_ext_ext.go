package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckExtLoader string *VehicleSupervision/internal/modules/vehicle/model.MuckTruckExt

func (t MuckTruckExt) TableName() string {
	return "muck_truck_ext"
}

func (t *MuckTruckExt) NewLoader() *MuckTruckExtLoader {
	return &MuckTruckExtLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckExt, []error) {
			var rs []*MuckTruckExt
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckExt{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
