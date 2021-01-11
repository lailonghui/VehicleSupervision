package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden FingerprintDriverUnionPkLoader string *VehicleSupervision/internal/modules/device/model.FingerprintDriver

// 数据库表名
func (t *FingerprintDriver) TableName() string {
	return "fingerprint_driver"
}

// 主键列名
func (t *FingerprintDriver) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *FingerprintDriver) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *FingerprintDriverPkLoader) NewLoader() *FingerprintDriverPkLoader {
	return &FingerprintDriverPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*FingerprintDriver, []error) {
			var rs []*FingerprintDriver
			var m FingerprintDriver
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *FingerprintDriver) UnionPrimaryColumnName() string {
	return "fingerprint_driver_id"
}

// 获取联合主键
func (t *FingerprintDriver) GetUnionPrimary() string {
	return t.FingerprintDriverID
}

// 新建联合主键dataloader
func (t *FingerprintDriverUnionPkLoader) NewLoader() *FingerprintDriverUnionPkLoader {
	return &FingerprintDriverUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*FingerprintDriver, []error) {
			var rs []*FingerprintDriver
			var m FingerprintDriver
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
