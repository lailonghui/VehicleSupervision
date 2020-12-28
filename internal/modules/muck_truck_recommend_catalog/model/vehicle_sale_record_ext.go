package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleSaleRecordLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.VehicleSaleRecord

func (t VehicleSaleRecord) TableName() string {
	return "vehicle_sale_record"
}

func (t *VehicleSaleRecord) NewLoader() *VehicleSaleRecordLoader {
	return &VehicleSaleRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleSaleRecord, []error) {
			var rs []*VehicleSaleRecord
			// TODO 按实际需要实现
			db.DB.Model(&VehicleSaleRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
