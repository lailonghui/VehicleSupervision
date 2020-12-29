package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalChangeLoader string *VehicleSupervision/internal/modules/device/model.TerminalChange

// 数据库表名
func (t TerminalChange) TableName() string {
	return "terminal_change"
}

// 主键列名
func (t TerminalChange) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalChange) UnionPrimaryColumnName() string {
	return "change_id"
}

// 新建dataloader
func (t *TerminalChange) NewLoader() *TerminalChangeLoader {
	return &TerminalChangeLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalChange, []error) {
			var rs []*TerminalChange

			db.DB.Model(&TerminalChange{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
