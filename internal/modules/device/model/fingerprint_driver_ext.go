package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden FingerprintDriverUnionPkLoader string *VehicleSupervision/internal/modules/device/model.FingerprintDriver

// 数据库表名
func (t FingerprintDriver) TableName() string {
	return "fingerprint_driver"
}

// 主键列名
func (t FingerprintDriver) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *FingerprintDriver) NewPkLoader() *FingerprintDriverPkLoader {
	return &FingerprintDriverPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*FingerprintDriver, []error) {
			var rs []*FingerprintDriver
			db.DB.Model(&FingerprintDriver{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t FingerprintDriver) UnionPrimaryColumnName() string {
	return "fingerprint_driver_id"
}

// 新建联合主键dataloader
func (t *FingerprintDriver) NewUnionPkLoader() *FingerprintDriverUnionPkLoader {
	return &FingerprintDriverUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*FingerprintDriver, []error) {
			var rs []*FingerprintDriver
			db.DB.Model(&FingerprintDriver{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
