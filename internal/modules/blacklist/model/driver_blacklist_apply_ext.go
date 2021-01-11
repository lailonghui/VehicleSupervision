package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DriverBlacklistApplyUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.DriverBlacklistApply

// 数据库表名
func (t *DriverBlacklistApply) TableName() string {
	return "driver_blacklist_apply"
}

// 主键列名
func (t *DriverBlacklistApply) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *DriverBlacklistApply) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *DriverBlacklistApplyPkLoader) NewLoader() *DriverBlacklistApplyPkLoader {
	return &DriverBlacklistApplyPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverBlacklistApply, []error) {
			var rs []*DriverBlacklistApply
			var m DriverBlacklistApply
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *DriverBlacklistApply) UnionPrimaryColumnName() string {
	return "apply_id"
}

// 获取联合主键
func (t *DriverBlacklistApply) GetUnionPrimary() string {
	return t.ApplyID
}

// 新建联合主键dataloader
func (t *DriverBlacklistApplyUnionPkLoader) NewLoader() *DriverBlacklistApplyUnionPkLoader {
	return &DriverBlacklistApplyUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverBlacklistApply, []error) {
			var rs []*DriverBlacklistApply
			var m DriverBlacklistApply
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
