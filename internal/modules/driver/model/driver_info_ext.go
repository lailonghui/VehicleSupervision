package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DriverInfoLoader string *VehicleSupervision/internal/modules/driver/model.DriverInfo

func (t DriverInfo) TableName() string {
	return "DriverInfo"
}

func (t *DriverInfo) NewLoader() *DriverInfoLoader {
	return &DriverInfoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverInfo, []error) {
			var rs []*DriverInfo
			// TODO 按实际需要实现
			db.DB.Model(&DriverInfo{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
