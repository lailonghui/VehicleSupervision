package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalProveUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalProve

// 数据库表名
func (t *TerminalProve) TableName() string {
	return "terminal_prove"
}

// 主键列名
func (t *TerminalProve) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalProve) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalProvePkLoader) NewLoader() *TerminalProvePkLoader {
	return &TerminalProvePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalProve, []error) {
			var rs []*TerminalProve
			var m TerminalProve
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalProve) UnionPrimaryColumnName() string {
	return "prove_id"
}

// 获取联合主键
func (t *TerminalProve) GetUnionPrimary() string {
	return t.ProveID
}

// 新建联合主键dataloader
func (t *TerminalProveUnionPkLoader) NewLoader() *TerminalProveUnionPkLoader {
	return &TerminalProveUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalProve, []error) {
			var rs []*TerminalProve
			var m TerminalProve
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
