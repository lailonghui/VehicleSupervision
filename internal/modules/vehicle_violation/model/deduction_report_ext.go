package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DeductionReportLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.DeductionReport

func (t DeductionReport) TableName() string {
	return "deduction_report"
}

func (t *DeductionReport) NewLoader() *DeductionReportLoader {
	return &DeductionReportLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DeductionReport, []error) {
			var rs []*DeductionReport
			// TODO 按实际需要实现
			db.DB.Model(&DeductionReport{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
