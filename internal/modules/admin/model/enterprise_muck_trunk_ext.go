package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseMuckTrunkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseMuckTrunk

// 数据库表名
func (t EnterpriseMuckTrunk) TableName() string {
	return "enterprise_muck_trunk"
}

// 主键列名
func (t EnterpriseMuckTrunk) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseMuckTrunk) UnionPrimaryColumnName() string {
	return "enterprise_muck_trunk_id"
}

// 新建dataloader
func (t *EnterpriseMuckTrunk) NewLoader() *EnterpriseMuckTrunkLoader {
	return &EnterpriseMuckTrunkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseMuckTrunk, []error) {
			var rs []*EnterpriseMuckTrunk

			db.DB.Model(&EnterpriseMuckTrunk{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
