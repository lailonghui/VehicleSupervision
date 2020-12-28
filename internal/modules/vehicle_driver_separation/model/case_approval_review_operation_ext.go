package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden CaseApprovalReviewOperationLoader string *VehicleSupervision/internal/modules/vehicle_driver_separation/model.CaseApprovalReviewOperation

func (t CaseApprovalReviewOperation) TableName() string {
	return "CaseApprovalReviewOperation"
}

func (t *CaseApprovalReviewOperation) NewLoader() *CaseApprovalReviewOperationLoader {
	return &CaseApprovalReviewOperationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*CaseApprovalReviewOperation, []error) {
			var rs []*CaseApprovalReviewOperation
			// TODO 按实际需要实现
			db.DB.Model(&CaseApprovalReviewOperation{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
