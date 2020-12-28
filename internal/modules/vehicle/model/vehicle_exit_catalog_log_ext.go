package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleExitCatalogLogLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleExitCatalogLog

func (t VehicleExitCatalogLog) TableName() string {
	return "vehicle_exit_catalog_log"
}

func (t *VehicleExitCatalogLog) NewLoader() *VehicleExitCatalogLogLoader {
	return &VehicleExitCatalogLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleExitCatalogLog, []error) {
			var rs []*VehicleExitCatalogLog
			// TODO 按实际需要实现
			db.DB.Model(&VehicleExitCatalogLog{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
