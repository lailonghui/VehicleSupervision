package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleInfoLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleInfo

func (t VehicleInfo) TableName() string {
	return "VehicleInfo"
}

func (t *VehicleInfo) NewLoader() *VehicleInfoLoader {
	return &VehicleInfoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleInfo, []error) {
			var rs []*VehicleInfo
			// TODO 按实际需要实现
			db.DB.Model(&VehicleInfo{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
