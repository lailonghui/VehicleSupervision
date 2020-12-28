package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DriverFingerprintAssociationLoader string *VehicleSupervision/internal/modules/driver/model.DriverFingerprintAssociation

func (t DriverFingerprintAssociation) TableName() string {
	return "DriverFingerprintAssociation"
}

func (t *DriverFingerprintAssociation) NewLoader() *DriverFingerprintAssociationLoader {
	return &DriverFingerprintAssociationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverFingerprintAssociation, []error) {
			var rs []*DriverFingerprintAssociation
			// TODO 按实际需要实现
			db.DB.Model(&DriverFingerprintAssociation{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
