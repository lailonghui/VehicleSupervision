package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DynamicSpotCheckDisposalLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.DynamicSpotCheckDisposal

func (t DynamicSpotCheckDisposal) TableName() string {
	return "DynamicSpotCheckDisposal"
}

func (t *DynamicSpotCheckDisposal) NewLoader() *DynamicSpotCheckDisposalLoader {
	return &DynamicSpotCheckDisposalLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DynamicSpotCheckDisposal, []error) {
			var rs []*DynamicSpotCheckDisposal
			// TODO 按实际需要实现
			db.DB.Model(&DynamicSpotCheckDisposal{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
