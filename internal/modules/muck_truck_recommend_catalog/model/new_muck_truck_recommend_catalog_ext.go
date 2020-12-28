package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden NewMuckTruckRecommendCatalogLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.NewMuckTruckRecommendCatalog

func (t NewMuckTruckRecommendCatalog) TableName() string {
	return "new_muck_truck_recommend_catalog"
}

func (t *NewMuckTruckRecommendCatalog) NewLoader() *NewMuckTruckRecommendCatalogLoader {
	return &NewMuckTruckRecommendCatalogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*NewMuckTruckRecommendCatalog, []error) {
			var rs []*NewMuckTruckRecommendCatalog
			// TODO 按实际需要实现
			db.DB.Model(&NewMuckTruckRecommendCatalog{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
