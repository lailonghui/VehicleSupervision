package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamItemLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamItem

// 数据库表名
func (t TerminalParamItem) TableName() string {
	return "terminal_param_item"
}

// 主键列名
func (t TerminalParamItem) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalParamItem) UnionPrimaryColumnName() string {
	return "param_item_id"
}

// 新建dataloader
func (t *TerminalParamItem) NewLoader() *TerminalParamItemLoader {
	return &TerminalParamItemLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamItem, []error) {
			var rs []*TerminalParamItem

			db.DB.Model(&TerminalParamItem{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
