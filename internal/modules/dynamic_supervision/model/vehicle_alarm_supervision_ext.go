package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleAlarmSupervisionLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.VehicleAlarmSupervision

func (t VehicleAlarmSupervision) TableName() string {
	return "vehicle_alarm_supervision"
}

func (t *VehicleAlarmSupervision) NewLoader() *VehicleAlarmSupervisionLoader {
	return &VehicleAlarmSupervisionLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleAlarmSupervision, []error) {
			var rs []*VehicleAlarmSupervision
			// TODO 按实际需要实现
			db.DB.Model(&VehicleAlarmSupervision{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
