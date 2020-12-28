package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DynamicSupervisionLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.DynamicSupervision

func (t DynamicSupervision) TableName() string {
	return "DynamicSupervision"
}

func (t *DynamicSupervision) NewLoader() *DynamicSupervisionLoader {
	return &DynamicSupervisionLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DynamicSupervision, []error) {
			var rs []*DynamicSupervision
			// TODO 按实际需要实现
			db.DB.Model(&DynamicSupervision{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
