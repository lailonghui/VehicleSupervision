package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckWorkerIdCardOrdersLoader string *VehicleSupervision/internal/modules/vehicle/model.MuckTruckWorkerIdCardOrders

func (t MuckTruckWorkerIdCardOrders) TableName() string {
	return "MuckTruckWorkerIdCardOrders"
}

func (t *MuckTruckWorkerIdCardOrders) NewLoader() *MuckTruckWorkerIdCardOrdersLoader {
	return &MuckTruckWorkerIdCardOrdersLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckWorkerIdCardOrders, []error) {
			var rs []*MuckTruckWorkerIdCardOrders
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckWorkerIdCardOrders{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
