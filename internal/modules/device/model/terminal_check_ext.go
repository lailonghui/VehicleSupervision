package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalCheckUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalCheck

// 数据库表名
func (t TerminalCheck) TableName() string {
	return "terminal_check"
}

// 主键列名
func (t TerminalCheck) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *TerminalCheck) NewPkLoader() *TerminalCheckPkLoader {
	return &TerminalCheckPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheck, []error) {
			var rs []*TerminalCheck
			db.DB.Model(&TerminalCheck{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t TerminalCheck) UnionPrimaryColumnName() string {
	return "terminal_check_id"
}

// 新建联合主键dataloader
func (t *TerminalCheck) NewUnionPkLoader() *TerminalCheckUnionPkLoader {
	return &TerminalCheckUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheck, []error) {
			var rs []*TerminalCheck
			db.DB.Model(&TerminalCheck{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
