package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden FingerprintInformationLoader string *VehicleSupervision/internal/modules/device/model.FingerprintInformation

// 数据库表名
func (t FingerprintInformation) TableName() string {
	return "fingerprint_information"
}

// 主键列名
func (t FingerprintInformation) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t FingerprintInformation) UnionPrimaryColumnName() string {
	return "fingerprint_information_id"
}

// 新建dataloader
func (t *FingerprintInformation) NewLoader() *FingerprintInformationLoader {
	return &FingerprintInformationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*FingerprintInformation, []error) {
			var rs []*FingerprintInformation

			db.DB.Model(&FingerprintInformation{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
