package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleExitCatalogReviewLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleExitCatalogReview

func (t VehicleExitCatalogReview) TableName() string {
	return "vehicle_exit_catalog_review"
}

func (t *VehicleExitCatalogReview) NewLoader() *VehicleExitCatalogReviewLoader {
	return &VehicleExitCatalogReviewLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleExitCatalogReview, []error) {
			var rs []*VehicleExitCatalogReview
			// TODO 按实际需要实现
			db.DB.Model(&VehicleExitCatalogReview{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
