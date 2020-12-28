package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckPurchaseIntentionLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.MuckTruckPurchaseIntention

func (t MuckTruckPurchaseIntention) TableName() string {
	return "muck_truck_purchase_intention"
}

func (t *MuckTruckPurchaseIntention) NewLoader() *MuckTruckPurchaseIntentionLoader {
	return &MuckTruckPurchaseIntentionLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckPurchaseIntention, []error) {
			var rs []*MuckTruckPurchaseIntention
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckPurchaseIntention{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
