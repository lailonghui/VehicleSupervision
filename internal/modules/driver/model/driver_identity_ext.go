package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DriverIdentityLoader string *VehicleSupervision/internal/modules/driver/model.DriverIdentity

func (t DriverIdentity) TableName() string {
	return "driver_identity"
}

func (t *DriverIdentity) NewLoader() *DriverIdentityLoader {
	return &DriverIdentityLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverIdentity, []error) {
			var rs []*DriverIdentity
			// TODO 按实际需要实现
			db.DB.Model(&DriverIdentity{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
