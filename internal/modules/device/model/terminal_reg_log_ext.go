package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalRegLogPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalRegLog

// 数据库表名
func (t *TerminalRegLog) TableName() string {
	return "terminal_reg_log"
}

// 主键列名
func (t *TerminalRegLog) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalRegLog) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalRegLogPkLoader) NewLoader() *TerminalRegLogPkLoader {
	return &TerminalRegLogPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalRegLog, []error) {
			var rs []*TerminalRegLog
			var m TerminalRegLog
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
