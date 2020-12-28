package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleSupervisionPhotoLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleSupervisionPhoto

func (t VehicleSupervisionPhoto) TableName() string {
	return "VehicleSupervisionPhoto"
}

func (t *VehicleSupervisionPhoto) NewLoader() *VehicleSupervisionPhotoLoader {
	return &VehicleSupervisionPhotoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleSupervisionPhoto, []error) {
			var rs []*VehicleSupervisionPhoto
			// TODO 按实际需要实现
			db.DB.Model(&VehicleSupervisionPhoto{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
