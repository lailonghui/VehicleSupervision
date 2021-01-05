package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalCheckLastUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalCheckLast

// 数据库表名
func (t TerminalCheckLast) TableName() string {
	return "terminal_check_last"
}

// 主键列名
func (t TerminalCheckLast) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *TerminalCheckLast) NewPkLoader() *TerminalCheckLastPkLoader {
	return &TerminalCheckLastPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheckLast, []error) {
			var rs []*TerminalCheckLast
			db.DB.Model(&TerminalCheckLast{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t TerminalCheckLast) UnionPrimaryColumnName() string {
	return "terminal_check_last_id"
}

// 新建联合主键dataloader
func (t *TerminalCheckLast) NewUnionPkLoader() *TerminalCheckLastUnionPkLoader {
	return &TerminalCheckLastUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheckLast, []error) {
			var rs []*TerminalCheckLast
			db.DB.Model(&TerminalCheckLast{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
