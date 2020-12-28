package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden AlarmSupervisionPictureUploadLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.AlarmSupervisionPictureUpload

func (t AlarmSupervisionPictureUpload) TableName() string {
	return "AlarmSupervisionPictureUpload"
}

func (t *AlarmSupervisionPictureUpload) NewLoader() *AlarmSupervisionPictureUploadLoader {
	return &AlarmSupervisionPictureUploadLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*AlarmSupervisionPictureUpload, []error) {
			var rs []*AlarmSupervisionPictureUpload
			// TODO 按实际需要实现
			db.DB.Model(&AlarmSupervisionPictureUpload{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
