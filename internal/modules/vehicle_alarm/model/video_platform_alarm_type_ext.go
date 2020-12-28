package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VideoPlatformAlarmTypeLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.VideoPlatformAlarmType

func (t VideoPlatformAlarmType) TableName() string {
	return "VideoPlatformAlarmType"
}

func (t *VideoPlatformAlarmType) NewLoader() *VideoPlatformAlarmTypeLoader {
	return &VideoPlatformAlarmTypeLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VideoPlatformAlarmType, []error) {
			var rs []*VideoPlatformAlarmType
			// TODO 按实际需要实现
			db.DB.Model(&VideoPlatformAlarmType{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
