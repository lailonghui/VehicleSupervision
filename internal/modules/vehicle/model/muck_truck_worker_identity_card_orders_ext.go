package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckWorkerIdentityCardOrdersLoader string *VehicleSupervision/internal/modules/vehicle/model.MuckTruckWorkerIdentityCardOrders

func (t MuckTruckWorkerIdentityCardOrders) TableName() string {
	return "muck_truck_worker_identity_card_orders"
}

func (t *MuckTruckWorkerIdentityCardOrders) NewLoader() *MuckTruckWorkerIdentityCardOrdersLoader {
	return &MuckTruckWorkerIdentityCardOrdersLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckWorkerIdentityCardOrders, []error) {
			var rs []*MuckTruckWorkerIdentityCardOrders
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckWorkerIdentityCardOrders{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
