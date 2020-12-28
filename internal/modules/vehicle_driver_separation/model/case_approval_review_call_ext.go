package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden CaseApprovalReviewCallLoader string *VehicleSupervision/internal/modules/vehicle_driver_separation/model.CaseApprovalReviewCall

func (t CaseApprovalReviewCall) TableName() string {
	return "CaseApprovalReviewCall"
}

func (t *CaseApprovalReviewCall) NewLoader() *CaseApprovalReviewCallLoader {
	return &CaseApprovalReviewCallLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*CaseApprovalReviewCall, []error) {
			var rs []*CaseApprovalReviewCall
			// TODO 按实际需要实现
			db.DB.Model(&CaseApprovalReviewCall{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
