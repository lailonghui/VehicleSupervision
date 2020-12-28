package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden OfflineRegistrationRecordLoader string *VehicleSupervision/internal/modules/vehicle_snapshot_system/model.OfflineRegistrationRecord

func (t OfflineRegistrationRecord) TableName() string {
	return "OfflineRegistrationRecord"
}

func (t *OfflineRegistrationRecord) NewLoader() *OfflineRegistrationRecordLoader {
	return &OfflineRegistrationRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*OfflineRegistrationRecord, []error) {
			var rs []*OfflineRegistrationRecord
			// TODO 按实际需要实现
			db.DB.Model(&OfflineRegistrationRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
