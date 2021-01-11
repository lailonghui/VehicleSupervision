package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalOperLogUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalOperLog

// 数据库表名
func (t *TerminalOperLog) TableName() string {
	return "terminal_oper_log"
}

// 主键列名
func (t *TerminalOperLog) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalOperLog) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalOperLogPkLoader) NewLoader() *TerminalOperLogPkLoader {
	return &TerminalOperLogPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalOperLog, []error) {
			var rs []*TerminalOperLog
			var m TerminalOperLog
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalOperLog) UnionPrimaryColumnName() string {
	return "log_id"
}

// 获取联合主键
func (t *TerminalOperLog) GetUnionPrimary() string {
	return t.LogID
}

// 新建联合主键dataloader
func (t *TerminalOperLogUnionPkLoader) NewLoader() *TerminalOperLogUnionPkLoader {
	return &TerminalOperLogUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalOperLog, []error) {
			var rs []*TerminalOperLog
			var m TerminalOperLog
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
