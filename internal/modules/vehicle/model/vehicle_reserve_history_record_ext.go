package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleReserveHistoryRecordLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleReserveHistoryRecord

func (t VehicleReserveHistoryRecord) TableName() string {
	return "VehicleReserveHistoryRecord"
}

func (t *VehicleReserveHistoryRecord) NewLoader() *VehicleReserveHistoryRecordLoader {
	return &VehicleReserveHistoryRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleReserveHistoryRecord, []error) {
			var rs []*VehicleReserveHistoryRecord
			// TODO 按实际需要实现
			db.DB.Model(&VehicleReserveHistoryRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
