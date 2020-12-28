package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DriverIdentityInfoReportLoader string *VehicleSupervision/internal/modules/driver/model.DriverIdentityInfoReport

func (t DriverIdentityInfoReport) TableName() string {
	return "driver_identity_info_report"
}

func (t *DriverIdentityInfoReport) NewLoader() *DriverIdentityInfoReportLoader {
	return &DriverIdentityInfoReportLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverIdentityInfoReport, []error) {
			var rs []*DriverIdentityInfoReport
			// TODO 按实际需要实现
			db.DB.Model(&DriverIdentityInfoReport{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
