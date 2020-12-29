package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalLoader string *VehicleSupervision/internal/modules/device/model.Terminal

// 数据库表名
func (t Terminal) TableName() string {
	return "terminal"
}

// 主键列名
func (t Terminal) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t Terminal) UnionPrimaryColumnName() string {
	return "terminal_id"
}

// 新建dataloader
func (t *Terminal) NewLoader() *TerminalLoader {
	return &TerminalLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Terminal, []error) {
			var rs []*Terminal

			db.DB.Model(&Terminal{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
