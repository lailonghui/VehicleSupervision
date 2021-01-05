package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseMuckTrunkUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseMuckTrunk

// 数据库表名
func (t EnterpriseMuckTrunk) TableName() string {
	return "enterprise_muck_trunk"
}

// 主键列名
func (t EnterpriseMuckTrunk) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EnterpriseMuckTrunkPkLoader) NewLoader() *EnterpriseMuckTrunkPkLoader {
	return &EnterpriseMuckTrunkPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseMuckTrunk, []error) {
			var rs []*EnterpriseMuckTrunk
			var m EnterpriseMuckTrunk
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EnterpriseMuckTrunk) UnionPrimaryColumnName() string {
	return "enterprise_muck_trunk_id"
}

// 新建联合主键dataloader
func (t *EnterpriseMuckTrunkUnionPkLoader) NewLoader() *EnterpriseMuckTrunkUnionPkLoader {
	return &EnterpriseMuckTrunkUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseMuckTrunk, []error) {
			var rs []*EnterpriseMuckTrunk
			var m EnterpriseMuckTrunk
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
