package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SystemUserUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.SystemUser

// 数据库表名
func (t SystemUser) TableName() string {
	return "system_user"
}

// 主键列名
func (t SystemUser) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *SystemUserPkLoader) NewLoader() *SystemUserPkLoader {
	return &SystemUserPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SystemUser, []error) {
			var rs []*SystemUser
			var m SystemUser
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t SystemUser) UnionPrimaryColumnName() string {
	return "user_id"
}

// 新建联合主键dataloader
func (t *SystemUserUnionPkLoader) NewLoader() *SystemUserUnionPkLoader {
	return &SystemUserUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SystemUser, []error) {
			var rs []*SystemUser
			var m SystemUser
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
