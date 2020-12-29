package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SystemUserLoader string *VehicleSupervision/internal/modules/admin/model.SystemUser

// 数据库表名
func (t SystemUser) TableName() string {
	return "system_user"
}

// 主键列名
func (t SystemUser) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t SystemUser) UnionPrimaryColumnName() string {
	return "user_id"
}

// 新建dataloader
func (t *SystemUser) NewLoader() *SystemUserLoader {
	return &SystemUserLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SystemUser, []error) {
			var rs []*SystemUser

			db.DB.Model(&SystemUser{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
