package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden ConstructionInfoLoader string *VehicleSupervision/internal/modules/vehicle_snapshot_system/model.ConstructionInfo

func (t ConstructionInfo) TableName() string {
	return "ConstructionInfo"
}

func (t *ConstructionInfo) NewLoader() *ConstructionInfoLoader {
	return &ConstructionInfoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ConstructionInfo, []error) {
			var rs []*ConstructionInfo
			// TODO 按实际需要实现
			db.DB.Model(&ConstructionInfo{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
