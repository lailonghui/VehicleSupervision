package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DriverIdInfoReportLoader string *VehicleSupervision/internal/modules/driver/model.DriverIdInfoReport

func (t DriverIdInfoReport) TableName() string {
	return "DriverIdInfoReport"
}

func (t *DriverIdInfoReport) NewLoader() *DriverIdInfoReportLoader {
	return &DriverIdInfoReportLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverIdInfoReport, []error) {
			var rs []*DriverIdInfoReport
			// TODO 按实际需要实现
			db.DB.Model(&DriverIdInfoReport{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
