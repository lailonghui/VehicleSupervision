package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamItemUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamItem

// 数据库表名
func (t *TerminalParamItem) TableName() string {
	return "terminal_param_item"
}

// 主键列名
func (t *TerminalParamItem) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalParamItem) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalParamItemPkLoader) NewLoader() *TerminalParamItemPkLoader {
	return &TerminalParamItemPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamItem, []error) {
			var rs []*TerminalParamItem
			var m TerminalParamItem
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalParamItem) UnionPrimaryColumnName() string {
	return "param_item_id"
}

// 获取联合主键
func (t *TerminalParamItem) GetUnionPrimary() string {
	return t.ParamItemID
}

// 新建联合主键dataloader
func (t *TerminalParamItemUnionPkLoader) NewLoader() *TerminalParamItemUnionPkLoader {
	return &TerminalParamItemUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamItem, []error) {
			var rs []*TerminalParamItem
			var m TerminalParamItem
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
