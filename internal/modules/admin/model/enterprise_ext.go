package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.Enterprise

// 数据库表名
func (t *Enterprise) TableName() string {
	return "enterprise"
}

// 主键列名
func (t *Enterprise) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *Enterprise) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EnterprisePkLoader) NewLoader() *EnterprisePkLoader {
	return &EnterprisePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Enterprise, []error) {
			var rs []*Enterprise
			var m Enterprise
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *Enterprise) UnionPrimaryColumnName() string {
	return "enterprise_id"
}

// 获取联合主键
func (t *Enterprise) GetUnionPrimary() string {
	return t.EnterpriseID
}

// 新建联合主键dataloader
func (t *EnterpriseUnionPkLoader) NewLoader() *EnterpriseUnionPkLoader {
	return &EnterpriseUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Enterprise, []error) {
			var rs []*Enterprise
			var m Enterprise
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
