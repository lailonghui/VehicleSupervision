package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden NewMuckTruckInfoLoader string *VehicleSupervision/internal/modules/muck_truck_recommend_catalog/model.NewMuckTruckInfo

func (t NewMuckTruckInfo) TableName() string {
	return "new_muck_truck_info"
}

func (t *NewMuckTruckInfo) NewLoader() *NewMuckTruckInfoLoader {
	return &NewMuckTruckInfoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*NewMuckTruckInfo, []error) {
			var rs []*NewMuckTruckInfo
			// TODO 按实际需要实现
			db.DB.Model(&NewMuckTruckInfo{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
