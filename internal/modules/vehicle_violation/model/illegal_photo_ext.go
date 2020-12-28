package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden IllegalPhotoLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.IllegalPhoto

func (t IllegalPhoto) TableName() string {
	return "illegal_photo"
}

func (t *IllegalPhoto) NewLoader() *IllegalPhotoLoader {
	return &IllegalPhotoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*IllegalPhoto, []error) {
			var rs []*IllegalPhoto
			// TODO 按实际需要实现
			db.DB.Model(&IllegalPhoto{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
