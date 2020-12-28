package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleOnlineTimeLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleOnlineTime

func (t VehicleOnlineTime) TableName() string {
	return "VehicleOnlineTime"
}

func (t *VehicleOnlineTime) NewLoader() *VehicleOnlineTimeLoader {
	return &VehicleOnlineTimeLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleOnlineTime, []error) {
			var rs []*VehicleOnlineTime
			// TODO 按实际需要实现
			db.DB.Model(&VehicleOnlineTime{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
