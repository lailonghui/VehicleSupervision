package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleViolationScoringItemsLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.VehicleViolationScoringItems

func (t VehicleViolationScoringItems) TableName() string {
	return "vehicle_violation_scoring_items"
}

func (t *VehicleViolationScoringItems) NewLoader() *VehicleViolationScoringItemsLoader {
	return &VehicleViolationScoringItemsLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleViolationScoringItems, []error) {
			var rs []*VehicleViolationScoringItems
			// TODO 按实际需要实现
			db.DB.Model(&VehicleViolationScoringItems{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
