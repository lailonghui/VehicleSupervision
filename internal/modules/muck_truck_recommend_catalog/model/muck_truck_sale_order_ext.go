package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckSaleOrderLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.MuckTruckSaleOrder

func (t MuckTruckSaleOrder) TableName() string {
	return "muck_truck_sale_order"
}

func (t *MuckTruckSaleOrder) NewLoader() *MuckTruckSaleOrderLoader {
	return &MuckTruckSaleOrderLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckSaleOrder, []error) {
			var rs []*MuckTruckSaleOrder
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckSaleOrder{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
