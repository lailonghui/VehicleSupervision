package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DepartmentLoader string *VehicleSupervision/internal/modules/admin/model.Department

// 数据库表名
func (t Department) TableName() string {
	return "department"
}

// 主键列名
func (t Department) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t Department) UnionPrimaryColumnName() string {
	return "department_id"
}

// 新建dataloader
func (t *Department) NewLoader() *DepartmentLoader {
	return &DepartmentLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Department, []error) {
			var rs []*Department

			db.DB.Model(&Department{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
