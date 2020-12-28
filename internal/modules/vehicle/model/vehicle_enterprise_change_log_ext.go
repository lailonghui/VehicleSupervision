package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleEnterpriseChangeLogLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleEnterpriseChangeLog

func (t VehicleEnterpriseChangeLog) TableName() string {
	return "vehicle_enterprise_change_log"
}

func (t *VehicleEnterpriseChangeLog) NewLoader() *VehicleEnterpriseChangeLogLoader {
	return &VehicleEnterpriseChangeLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleEnterpriseChangeLog, []error) {
			var rs []*VehicleEnterpriseChangeLog
			// TODO 按实际需要实现
			db.DB.Model(&VehicleEnterpriseChangeLog{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
