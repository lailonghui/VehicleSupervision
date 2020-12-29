package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalOperLogLoader string *VehicleSupervision/internal/modules/device/model.TerminalOperLog

// 数据库表名
func (t TerminalOperLog) TableName() string {
	return "terminal_oper_log"
}

// 主键列名
func (t TerminalOperLog) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalOperLog) UnionPrimaryColumnName() string {
	return "log_id"
}

// 新建dataloader
func (t *TerminalOperLog) NewLoader() *TerminalOperLogLoader {
	return &TerminalOperLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalOperLog, []error) {
			var rs []*TerminalOperLog

			db.DB.Model(&TerminalOperLog{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
