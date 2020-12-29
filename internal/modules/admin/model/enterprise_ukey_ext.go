package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseUkeyLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseUkey

// 数据库表名
func (t EnterpriseUkey) TableName() string {
	return "enterprise_ukey"
}

// 主键列名
func (t EnterpriseUkey) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseUkey) UnionPrimaryColumnName() string {
	return "ukey_id"
}

// 新建dataloader
func (t *EnterpriseUkey) NewLoader() *EnterpriseUkeyLoader {
	return &EnterpriseUkeyLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseUkey, []error) {
			var rs []*EnterpriseUkey

			db.DB.Model(&EnterpriseUkey{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
