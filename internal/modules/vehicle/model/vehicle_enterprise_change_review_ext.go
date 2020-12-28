package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleEnterpriseChangeReviewLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleEnterpriseChangeReview

func (t VehicleEnterpriseChangeReview) TableName() string {
	return "vehicle_enterprise_change_review"
}

func (t *VehicleEnterpriseChangeReview) NewLoader() *VehicleEnterpriseChangeReviewLoader {
	return &VehicleEnterpriseChangeReviewLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleEnterpriseChangeReview, []error) {
			var rs []*VehicleEnterpriseChangeReview
			// TODO 按实际需要实现
			db.DB.Model(&VehicleEnterpriseChangeReview{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
