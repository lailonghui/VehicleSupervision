package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DynamicSupervisionDetailLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.DynamicSupervisionDetail

func (t DynamicSupervisionDetail) TableName() string {
	return "DynamicSupervisionDetail"
}

func (t *DynamicSupervisionDetail) NewLoader() *DynamicSupervisionDetailLoader {
	return &DynamicSupervisionDetailLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DynamicSupervisionDetail, []error) {
			var rs []*DynamicSupervisionDetail
			// TODO 按实际需要实现
			db.DB.Model(&DynamicSupervisionDetail{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
