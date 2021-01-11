package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalUnionPkLoader string *VehicleSupervision/internal/modules/device/model.Terminal

// 数据库表名
func (t *Terminal) TableName() string {
	return "terminal"
}

// 主键列名
func (t *Terminal) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *Terminal) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalPkLoader) NewLoader() *TerminalPkLoader {
	return &TerminalPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Terminal, []error) {
			var rs []*Terminal
			var m Terminal
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *Terminal) UnionPrimaryColumnName() string {
	return "terminal_id"
}

// 获取联合主键
func (t *Terminal) GetUnionPrimary() string {
	return t.TerminalID
}

// 新建联合主键dataloader
func (t *TerminalUnionPkLoader) NewLoader() *TerminalUnionPkLoader {
	return &TerminalUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*Terminal, []error) {
			var rs []*Terminal
			var m Terminal
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
