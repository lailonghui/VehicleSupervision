package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden FingerprintLoader string *VehicleSupervision/internal/modules/device/model.Fingerprint

// 数据库表名
func (t Fingerprint) TableName() string {
	return "fingerprint"
}

// 主键列名
func (t Fingerprint) PrimaryColumnName() string {
	return "id"
}

// 新建dataloader
func (t *Fingerprint) NewLoader() *FingerprintLoader {
	return &FingerprintLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Fingerprint, []error) {
			var rs []*Fingerprint

			db.DB.Model(&Fingerprint{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
