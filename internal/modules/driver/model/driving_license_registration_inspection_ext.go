package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DrivingLicenseRegistrationInspectionLoader string *VehicleSupervision/internal/modules/driver/model.DrivingLicenseRegistrationInspection

func (t DrivingLicenseRegistrationInspection) TableName() string {
	return "DrivingLicenseRegistrationInspection"
}

func (t *DrivingLicenseRegistrationInspection) NewLoader() *DrivingLicenseRegistrationInspectionLoader {
	return &DrivingLicenseRegistrationInspectionLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DrivingLicenseRegistrationInspection, []error) {
			var rs []*DrivingLicenseRegistrationInspection
			// TODO 按实际需要实现
			db.DB.Model(&DrivingLicenseRegistrationInspection{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
