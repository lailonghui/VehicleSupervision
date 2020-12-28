package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckOnlineLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.MuckTruckOnline

func (t MuckTruckOnline) TableName() string {
	return "MuckTruckOnline"
}

func (t *MuckTruckOnline) NewLoader() *MuckTruckOnlineLoader {
	return &MuckTruckOnlineLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckOnline, []error) {
			var rs []*MuckTruckOnline
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckOnline{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
