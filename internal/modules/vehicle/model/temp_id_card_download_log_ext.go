package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden TempIdCardDownloadLogLoader string *VehicleSupervision/internal/modules/vehicle/model.TempIdCardDownloadLog

func (t TempIdCardDownloadLog) TableName() string {
	return "TempIdCardDownloadLog"
}

func (t *TempIdCardDownloadLog) NewLoader() *TempIdCardDownloadLogLoader {
	return &TempIdCardDownloadLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TempIdCardDownloadLog, []error) {
			var rs []*TempIdCardDownloadLog
			// TODO 按实际需要实现
			db.DB.Model(&TempIdCardDownloadLog{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
