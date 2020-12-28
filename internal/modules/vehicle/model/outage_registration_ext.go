package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden OutageRegistrationLoader string *VehicleSupervision/internal/modules/vehicle/model.OutageRegistration

func (t OutageRegistration) TableName() string {
	return "outage_registration"
}

func (t *OutageRegistration) NewLoader() *OutageRegistrationLoader {
	return &OutageRegistrationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*OutageRegistration, []error) {
			var rs []*OutageRegistration
			// TODO 按实际需要实现
			db.DB.Model(&OutageRegistration{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
