package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DriverBlacklistHisUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.DriverBlacklistHis

// 数据库表名
func (t *DriverBlacklistHis) TableName() string {
	return "driver_blacklist_his"
}

// 主键列名
func (t *DriverBlacklistHis) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *DriverBlacklistHis) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *DriverBlacklistHisPkLoader) NewLoader() *DriverBlacklistHisPkLoader {
	return &DriverBlacklistHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverBlacklistHis, []error) {
			var rs []*DriverBlacklistHis
			var m DriverBlacklistHis
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *DriverBlacklistHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 获取联合主键
func (t *DriverBlacklistHis) GetUnionPrimary() string {
	return t.HisID
}

// 新建联合主键dataloader
func (t *DriverBlacklistHisUnionPkLoader) NewLoader() *DriverBlacklistHisUnionPkLoader {
	return &DriverBlacklistHisUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverBlacklistHis, []error) {
			var rs []*DriverBlacklistHis
			var m DriverBlacklistHis
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
