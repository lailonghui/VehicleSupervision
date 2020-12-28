package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckSaleOrderDetailLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.MuckTruckSaleOrderDetail

func (t MuckTruckSaleOrderDetail) TableName() string {
	return "muck_truck_sale_order_detail"
}

func (t *MuckTruckSaleOrderDetail) NewLoader() *MuckTruckSaleOrderDetailLoader {
	return &MuckTruckSaleOrderDetailLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckSaleOrderDetail, []error) {
			var rs []*MuckTruckSaleOrderDetail
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckSaleOrderDetail{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
