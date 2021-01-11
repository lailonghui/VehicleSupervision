package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalModalUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalModal

// 数据库表名
func (t *TerminalModal) TableName() string {
	return "terminal_modal"
}

// 主键列名
func (t *TerminalModal) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalModal) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalModalPkLoader) NewLoader() *TerminalModalPkLoader {
	return &TerminalModalPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalModal, []error) {
			var rs []*TerminalModal
			var m TerminalModal
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalModal) UnionPrimaryColumnName() string {
	return "terminal_modal_id"
}

// 获取联合主键
func (t *TerminalModal) GetUnionPrimary() string {
	return t.TerminalModalID
}

// 新建联合主键dataloader
func (t *TerminalModalUnionPkLoader) NewLoader() *TerminalModalUnionPkLoader {
	return &TerminalModalUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalModal, []error) {
			var rs []*TerminalModal
			var m TerminalModal
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
