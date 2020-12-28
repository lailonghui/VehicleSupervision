package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DrivingLogInfoLoader string *VehicleSupervision/internal/modules/vehicle_driver_separation/model.DrivingLogInfo

func (t DrivingLogInfo) TableName() string {
	return "driving_log_info"
}

func (t *DrivingLogInfo) NewLoader() *DrivingLogInfoLoader {
	return &DrivingLogInfoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DrivingLogInfo, []error) {
			var rs []*DrivingLogInfo
			// TODO 按实际需要实现
			db.DB.Model(&DrivingLogInfo{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
