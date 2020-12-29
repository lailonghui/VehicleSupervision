package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseContactLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseContact

// 数据库表名
func (t EnterpriseContact) TableName() string {
	return "enterprise_contact"
}

// 主键列名
func (t EnterpriseContact) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseContact) UnionPrimaryColumnName() string {
	return "contact_id"
}

// 新建dataloader
func (t *EnterpriseContact) NewLoader() *EnterpriseContactLoader {
	return &EnterpriseContactLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseContact, []error) {
			var rs []*EnterpriseContact

			db.DB.Model(&EnterpriseContact{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
