package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleDriverBindingLoader string *VehicleSupervision/internal/modules/vehicle_driver_separation/model.VehicleDriverBinding

func (t VehicleDriverBinding) TableName() string {
	return "VehicleDriverBinding"
}

func (t *VehicleDriverBinding) NewLoader() *VehicleDriverBindingLoader {
	return &VehicleDriverBindingLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleDriverBinding, []error) {
			var rs []*VehicleDriverBinding
			// TODO 按实际需要实现
			db.DB.Model(&VehicleDriverBinding{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
