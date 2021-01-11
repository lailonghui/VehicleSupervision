package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalCheckUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalCheck

// 数据库表名
func (t *TerminalCheck) TableName() string {
	return "terminal_check"
}

// 主键列名
func (t *TerminalCheck) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalCheck) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalCheckPkLoader) NewLoader() *TerminalCheckPkLoader {
	return &TerminalCheckPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheck, []error) {
			var rs []*TerminalCheck
			var m TerminalCheck
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalCheck) UnionPrimaryColumnName() string {
	return "terminal_check_id"
}

// 获取联合主键
func (t *TerminalCheck) GetUnionPrimary() string {
	return t.TerminalCheckID
}

// 新建联合主键dataloader
func (t *TerminalCheckUnionPkLoader) NewLoader() *TerminalCheckUnionPkLoader {
	return &TerminalCheckUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheck, []error) {
			var rs []*TerminalCheck
			var m TerminalCheck
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
