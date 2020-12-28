package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleStateLatestLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleStateLatest

func (t VehicleStateLatest) TableName() string {
	return "vehicle_state_latest"
}

func (t *VehicleStateLatest) NewLoader() *VehicleStateLatestLoader {
	return &VehicleStateLatestLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleStateLatest, []error) {
			var rs []*VehicleStateLatest
			// TODO 按实际需要实现
			db.DB.Model(&VehicleStateLatest{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
