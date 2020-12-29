package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalFactoryLoader string *VehicleSupervision/internal/modules/device/model.TerminalFactory

// 数据库表名
func (t TerminalFactory) TableName() string {
	return "terminal_factory"
}

// 主键列名
func (t TerminalFactory) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalFactory) UnionPrimaryColumnName() string {
	return "factory_id"
}

// 新建dataloader
func (t *TerminalFactory) NewLoader() *TerminalFactoryLoader {
	return &TerminalFactoryLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalFactory, []error) {
			var rs []*TerminalFactory

			db.DB.Model(&TerminalFactory{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
