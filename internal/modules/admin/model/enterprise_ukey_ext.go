package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseUkeyUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseUkey

// 数据库表名
func (t *EnterpriseUkey) TableName() string {
	return "enterprise_ukey"
}

// 主键列名
func (t *EnterpriseUkey) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EnterpriseUkey) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EnterpriseUkeyPkLoader) NewLoader() *EnterpriseUkeyPkLoader {
	return &EnterpriseUkeyPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseUkey, []error) {
			var rs []*EnterpriseUkey
			var m EnterpriseUkey
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EnterpriseUkey) UnionPrimaryColumnName() string {
	return "ukey_id"
}

// 获取联合主键
func (t *EnterpriseUkey) GetUnionPrimary() string {
	return t.UkeyID
}

// 新建联合主键dataloader
func (t *EnterpriseUkeyUnionPkLoader) NewLoader() *EnterpriseUkeyUnionPkLoader {
	return &EnterpriseUkeyUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseUkey, []error) {
			var rs []*EnterpriseUkey
			var m EnterpriseUkey
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
