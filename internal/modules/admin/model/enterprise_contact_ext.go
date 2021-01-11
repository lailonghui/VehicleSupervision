package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseContactUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseContact

// 数据库表名
func (t *EnterpriseContact) TableName() string {
	return "enterprise_contact"
}

// 主键列名
func (t *EnterpriseContact) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EnterpriseContact) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EnterpriseContactPkLoader) NewLoader() *EnterpriseContactPkLoader {
	return &EnterpriseContactPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseContact, []error) {
			var rs []*EnterpriseContact
			var m EnterpriseContact
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EnterpriseContact) UnionPrimaryColumnName() string {
	return "contact_id"
}

// 获取联合主键
func (t *EnterpriseContact) GetUnionPrimary() string {
	return t.ContactID
}

// 新建联合主键dataloader
func (t *EnterpriseContactUnionPkLoader) NewLoader() *EnterpriseContactUnionPkLoader {
	return &EnterpriseContactUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseContact, []error) {
			var rs []*EnterpriseContact
			var m EnterpriseContact
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
