package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden OfflineAlarmRegistrationLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.OfflineAlarmRegistration

func (t OfflineAlarmRegistration) TableName() string {
	return "offline_alarm_registration"
}

func (t *OfflineAlarmRegistration) NewLoader() *OfflineAlarmRegistrationLoader {
	return &OfflineAlarmRegistrationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*OfflineAlarmRegistration, []error) {
			var rs []*OfflineAlarmRegistration
			// TODO 按实际需要实现
			db.DB.Model(&OfflineAlarmRegistration{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
