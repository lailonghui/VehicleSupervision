package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DepartmentUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.Department

// 数据库表名
func (t *Department) TableName() string {
	return "department"
}

// 主键列名
func (t *Department) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *Department) GetPrimary() int64 {
	return t.ID
}

// 联合主键列名
func (t *Department) UnionPrimaryColumnName() string {
	return "department_id"
}

// 获取联合主键
func (t *Department) GetUnionPrimary() string {
	return t.DepartmentID
}


// 新建主键dataloader
func (t *DepartmentPkLoader) NewLoader() *DepartmentPkLoader {
	return &DepartmentPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Department, []error) {
			var rs []*Department
			var m Department
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 新建联合主键dataloader
func (t *DepartmentUnionPkLoader) NewLoader() *DepartmentUnionPkLoader {
	return &DepartmentUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Department, []error) {
			var rs []*Department
			var m Department
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
