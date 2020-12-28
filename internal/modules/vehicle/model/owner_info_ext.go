package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden OwnerInfoLoader string *VehicleSupervision/internal/modules/vehicle/model.OwnerInfo

func (t OwnerInfo) TableName() string {
	return "OwnerInfo"
}

func (t *OwnerInfo) NewLoader() *OwnerInfoLoader {
	return &OwnerInfoLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*OwnerInfo, []error) {
			var rs []*OwnerInfo
			// TODO 按实际需要实现
			db.DB.Model(&OwnerInfo{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
