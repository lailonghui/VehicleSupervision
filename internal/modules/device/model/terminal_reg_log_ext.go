package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalRegLogPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalRegLog

// 数据库表名
func (t TerminalRegLog) TableName() string {
	return "terminal_reg_log"
}

// 主键列名
func (t TerminalRegLog) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *TerminalRegLog) NewPkLoader() *TerminalRegLogPkLoader {
	return &TerminalRegLogPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalRegLog, []error) {
			var rs []*TerminalRegLog
			db.DB.Model(&TerminalRegLog{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
