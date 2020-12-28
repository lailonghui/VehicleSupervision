package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DriverFingerprintLoader string *VehicleSupervision/internal/modules/driver/model.DriverFingerprint

func (t DriverFingerprint) TableName() string {
	return "driver_fingerprint"
}

func (t *DriverFingerprint) NewLoader() *DriverFingerprintLoader {
	return &DriverFingerprintLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverFingerprint, []error) {
			var rs []*DriverFingerprint
			// TODO 按实际需要实现
			db.DB.Model(&DriverFingerprint{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
