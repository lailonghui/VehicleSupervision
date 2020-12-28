package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden SellerRatingRecordLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.SellerRatingRecord

func (t SellerRatingRecord) TableName() string {
	return "seller_rating_record"
}

func (t *SellerRatingRecord) NewLoader() *SellerRatingRecordLoader {
	return &SellerRatingRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SellerRatingRecord, []error) {
			var rs []*SellerRatingRecord
			// TODO 按实际需要实现
			db.DB.Model(&SellerRatingRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
