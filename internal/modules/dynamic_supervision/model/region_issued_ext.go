package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden RegionIssuedLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.RegionIssued

func (t RegionIssued) TableName() string {
	return "region_issued"
}

func (t *RegionIssued) NewLoader() *RegionIssuedLoader {
	return &RegionIssuedLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RegionIssued, []error) {
			var rs []*RegionIssued
			// TODO 按实际需要实现
			db.DB.Model(&RegionIssued{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
