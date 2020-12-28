package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden ProvinceUploadVehicleLoader string *VehicleSupervision/internal/modules/vehicle/model.ProvinceUploadVehicle

func (t ProvinceUploadVehicle) TableName() string {
	return "province_upload_vehicle"
}

func (t *ProvinceUploadVehicle) NewLoader() *ProvinceUploadVehicleLoader {
	return &ProvinceUploadVehicleLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ProvinceUploadVehicle, []error) {
			var rs []*ProvinceUploadVehicle
			// TODO 按实际需要实现
			db.DB.Model(&ProvinceUploadVehicle{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
