package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleSecurityCheckRecordLoader string *VehicleSupervision/internal/modules/vehicle/model.VehicleSecurityCheckRecord

func (t VehicleSecurityCheckRecord) TableName() string {
	return "VehicleSecurityCheckRecord"
}

func (t *VehicleSecurityCheckRecord) NewLoader() *VehicleSecurityCheckRecordLoader {
	return &VehicleSecurityCheckRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleSecurityCheckRecord, []error) {
			var rs []*VehicleSecurityCheckRecord
			// TODO 按实际需要实现
			db.DB.Model(&VehicleSecurityCheckRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
