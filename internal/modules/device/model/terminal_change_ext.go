package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalChangeUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalChange

// 数据库表名
func (t *TerminalChange) TableName() string {
	return "terminal_change"
}

// 主键列名
func (t *TerminalChange) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalChange) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalChangePkLoader) NewLoader() *TerminalChangePkLoader {
	return &TerminalChangePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalChange, []error) {
			var rs []*TerminalChange
			var m TerminalChange
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalChange) UnionPrimaryColumnName() string {
	return "change_id"
}

// 获取联合主键
func (t *TerminalChange) GetUnionPrimary() string {
	return t.ChangeID
}

// 新建联合主键dataloader
func (t *TerminalChangeUnionPkLoader) NewLoader() *TerminalChangeUnionPkLoader {
	return &TerminalChangeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalChange, []error) {
			var rs []*TerminalChange
			var m TerminalChange
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
