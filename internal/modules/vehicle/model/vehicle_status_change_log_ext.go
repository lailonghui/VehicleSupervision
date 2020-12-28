package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleStatusChangeLogLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleStatusChangeLog

func (t VehicleStatusChangeLog) TableName() string {
	return "vehicle_status_change_log"
}

func (t *VehicleStatusChangeLog) NewLoader() *VehicleStatusChangeLogLoader {
	return &VehicleStatusChangeLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleStatusChangeLog, []error) {
			var rs []*VehicleStatusChangeLog
			// TODO 按实际需要实现
			db.DB.Model(&VehicleStatusChangeLog{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
