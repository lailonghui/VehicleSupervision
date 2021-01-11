package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalTypesPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalTypes

// 数据库表名
func (t *TerminalTypes) TableName() string {
	return "terminal_types"
}

// 主键列名
func (t *TerminalTypes) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalTypes) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalTypesPkLoader) NewLoader() *TerminalTypesPkLoader {
	return &TerminalTypesPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalTypes, []error) {
			var rs []*TerminalTypes
			var m TerminalTypes
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
