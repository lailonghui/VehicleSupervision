package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DriverBlacklistApplyUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.DriverBlacklistApply

// 数据库表名
func (t DriverBlacklistApply) TableName() string {
	return "driver_blacklist_apply"
}

// 主键列名
func (t DriverBlacklistApply) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *DriverBlacklistApply) NewPkLoader() *DriverBlacklistApplyPkLoader {
	return &DriverBlacklistApplyPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverBlacklistApply, []error) {
			var rs []*DriverBlacklistApply
			db.DB.Model(&DriverBlacklistApply{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t DriverBlacklistApply) UnionPrimaryColumnName() string {
	return "apply_id"
}

// 新建联合主键dataloader
func (t *DriverBlacklistApply) NewUnionPkLoader() *DriverBlacklistApplyUnionPkLoader {
	return &DriverBlacklistApplyUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverBlacklistApply, []error) {
			var rs []*DriverBlacklistApply
			db.DB.Model(&DriverBlacklistApply{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
