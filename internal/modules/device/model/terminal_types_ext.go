package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalTypesLoader string *VehicleSupervision/internal/modules/device/model.TerminalTypes

// 数据库表名
func (t TerminalTypes) TableName() string {
	return "terminal_types"
}

// 主键列名
func (t TerminalTypes) PrimaryColumnName() string {
	return "id"
}

// 新建dataloader
func (t *TerminalTypes) NewLoader() *TerminalTypesLoader {
	return &TerminalTypesLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalTypes, []error) {
			var rs []*TerminalTypes

			db.DB.Model(&TerminalTypes{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
