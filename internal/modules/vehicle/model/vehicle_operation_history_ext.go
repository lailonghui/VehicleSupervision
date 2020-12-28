package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleOperationHistoryLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleOperationHistory

func (t VehicleOperationHistory) TableName() string {
	return "vehicle_operation_history"
}

func (t *VehicleOperationHistory) NewLoader() *VehicleOperationHistoryLoader {
	return &VehicleOperationHistoryLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleOperationHistory, []error) {
			var rs []*VehicleOperationHistory
			// TODO 按实际需要实现
			db.DB.Model(&VehicleOperationHistory{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
