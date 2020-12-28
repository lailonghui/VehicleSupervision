package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden TempIdentityCardDownloadLogLoader string *VehicleSupervision/internal/modules/vehicle/model.TempIdentityCardDownloadLog

func (t TempIdentityCardDownloadLog) TableName() string {
	return "temp_identity_card_download_log"
}

func (t *TempIdentityCardDownloadLog) NewLoader() *TempIdentityCardDownloadLogLoader {
	return &TempIdentityCardDownloadLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TempIdentityCardDownloadLog, []error) {
			var rs []*TempIdentityCardDownloadLog
			// TODO 按实际需要实现
			db.DB.Model(&TempIdentityCardDownloadLog{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
