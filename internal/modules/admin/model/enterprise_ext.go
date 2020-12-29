package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseLoader string *VehicleSupervision/internal/modules/admin/model.Enterprise

// 数据库表名
func (t Enterprise) TableName() string {
	return "enterprise"
}

// 主键列名
func (t Enterprise) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t Enterprise) UnionPrimaryColumnName() string {
	return "enterprise_id"
}

// 新建dataloader
func (t *Enterprise) NewLoader() *EnterpriseLoader {
	return &EnterpriseLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Enterprise, []error) {
			var rs []*Enterprise

			db.DB.Model(&Enterprise{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
