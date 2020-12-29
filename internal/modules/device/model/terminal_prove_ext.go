package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalProveLoader string *VehicleSupervision/internal/modules/device/model.TerminalProve

// 数据库表名
func (t TerminalProve) TableName() string {
	return "terminal_prove"
}

// 主键列名
func (t TerminalProve) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalProve) UnionPrimaryColumnName() string {
	return "prove_id"
}

// 新建dataloader
func (t *TerminalProve) NewLoader() *TerminalProveLoader {
	return &TerminalProveLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalProve, []error) {
			var rs []*TerminalProve

			db.DB.Model(&TerminalProve{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
