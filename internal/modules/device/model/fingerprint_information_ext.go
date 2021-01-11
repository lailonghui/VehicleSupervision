package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden FingerprintInformationUnionPkLoader string *VehicleSupervision/internal/modules/device/model.FingerprintInformation

// 数据库表名
func (t *FingerprintInformation) TableName() string {
	return "fingerprint_information"
}

// 主键列名
func (t *FingerprintInformation) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *FingerprintInformation) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *FingerprintInformationPkLoader) NewLoader() *FingerprintInformationPkLoader {
	return &FingerprintInformationPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*FingerprintInformation, []error) {
			var rs []*FingerprintInformation
			var m FingerprintInformation
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *FingerprintInformation) UnionPrimaryColumnName() string {
	return "fingerprint_information_id"
}

// 获取联合主键
func (t *FingerprintInformation) GetUnionPrimary() string {
	return t.FingerprintInformationID
}

// 新建联合主键dataloader
func (t *FingerprintInformationUnionPkLoader) NewLoader() *FingerprintInformationUnionPkLoader {
	return &FingerprintInformationUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*FingerprintInformation, []error) {
			var rs []*FingerprintInformation
			var m FingerprintInformation
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
