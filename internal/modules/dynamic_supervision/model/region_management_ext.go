package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden RegionManagementLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.RegionManagement

func (t RegionManagement) TableName() string {
	return "region_management"
}

func (t *RegionManagement) NewLoader() *RegionManagementLoader {
	return &RegionManagementLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RegionManagement, []error) {
			var rs []*RegionManagement
			// TODO 按实际需要实现
			db.DB.Model(&RegionManagement{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
