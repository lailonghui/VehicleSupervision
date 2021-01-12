package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden FingerprintUnionPkLoader string *VehicleSupervision/internal/modules/device/model.Fingerprint

// 数据库表名
func (t *Fingerprint) TableName() string {
	return "fingerprint"
}

// 主键列名
func (t *Fingerprint) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *Fingerprint) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *FingerprintPkLoader) NewLoader() *FingerprintPkLoader {
	return &FingerprintPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Fingerprint, []error) {
			var rs []*Fingerprint
			var m Fingerprint
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *Fingerprint) UnionPrimaryColumnName() string {
	return "fingerprint_id"
}

// 获取联合主键
func (t *Fingerprint) GetUnionPrimary() string {
	return t.FingerprintID
}

// 新建联合主键dataloader
func (t *FingerprintUnionPkLoader) NewLoader() *FingerprintUnionPkLoader {
	return &FingerprintUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Fingerprint, []error) {
			var rs []*Fingerprint
			var m Fingerprint
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
