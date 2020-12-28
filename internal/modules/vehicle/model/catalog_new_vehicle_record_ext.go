package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden CatalogNewVehicleRecordLoader string *VehicleSupervision/internal/modules/vehicle/model.CatalogNewVehicleRecord

func (t CatalogNewVehicleRecord) TableName() string {
	return "CatalogNewVehicleRecord"
}

func (t *CatalogNewVehicleRecord) NewLoader() *CatalogNewVehicleRecordLoader {
	return &CatalogNewVehicleRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*CatalogNewVehicleRecord, []error) {
			var rs []*CatalogNewVehicleRecord
			// TODO 按实际需要实现
			db.DB.Model(&CatalogNewVehicleRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
