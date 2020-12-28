package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden SellerFilingLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.SellerFiling

func (t SellerFiling) TableName() string {
	return "seller_filing"
}

func (t *SellerFiling) NewLoader() *SellerFilingLoader {
	return &SellerFilingLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SellerFiling, []error) {
			var rs []*SellerFiling
			// TODO 按实际需要实现
			db.DB.Model(&SellerFiling{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
