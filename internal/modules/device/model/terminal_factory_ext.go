package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalFactoryUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalFactory

// 数据库表名
func (t *TerminalFactory) TableName() string {
	return "terminal_factory"
}

// 主键列名
func (t *TerminalFactory) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalFactory) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalFactoryPkLoader) NewLoader() *TerminalFactoryPkLoader {
	return &TerminalFactoryPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalFactory, []error) {
			var rs []*TerminalFactory
			var m TerminalFactory
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalFactory) UnionPrimaryColumnName() string {
	return "factory_id"
}

// 获取联合主键
func (t *TerminalFactory) GetUnionPrimary() string {
	return t.FactoryID
}

// 新建联合主键dataloader
func (t *TerminalFactoryUnionPkLoader) NewLoader() *TerminalFactoryUnionPkLoader {
	return &TerminalFactoryUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalFactory, []error) {
			var rs []*TerminalFactory
			var m TerminalFactory
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
