package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden NewMuckTruckPhotoLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.NewMuckTruckPhoto

func (t NewMuckTruckPhoto) TableName() string {
	return "new_muck_truck_photo"
}

func (t *NewMuckTruckPhoto) NewLoader() *NewMuckTruckPhotoLoader {
	return &NewMuckTruckPhotoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*NewMuckTruckPhoto, []error) {
			var rs []*NewMuckTruckPhoto
			// TODO 按实际需要实现
			db.DB.Model(&NewMuckTruckPhoto{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
