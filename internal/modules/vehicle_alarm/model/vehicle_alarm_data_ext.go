package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleAlarmDataLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.VehicleAlarmData

func (t VehicleAlarmData) TableName() string {
	return "vehicle_alarm_data"
}

func (t *VehicleAlarmData) NewLoader() *VehicleAlarmDataLoader {
	return &VehicleAlarmDataLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleAlarmData, []error) {
			var rs []*VehicleAlarmData
			// TODO 按实际需要实现
			db.DB.Model(&VehicleAlarmData{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
