package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalModalUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalModal

// 数据库表名
func (t TerminalModal) TableName() string {
	return "terminal_modal"
}

// 主键列名
func (t TerminalModal) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *TerminalModal) NewPkLoader() *TerminalModalPkLoader {
	return &TerminalModalPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalModal, []error) {
			var rs []*TerminalModal
			db.DB.Model(&TerminalModal{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t TerminalModal) UnionPrimaryColumnName() string {
	return "terminal_modal_id"
}

// 新建联合主键dataloader
func (t *TerminalModal) NewUnionPkLoader() *TerminalModalUnionPkLoader {
	return &TerminalModalUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalModal, []error) {
			var rs []*TerminalModal
			db.DB.Model(&TerminalModal{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
