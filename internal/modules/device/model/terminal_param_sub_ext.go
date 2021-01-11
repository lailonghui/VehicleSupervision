package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamSubUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamSub

// 数据库表名
func (t *TerminalParamSub) TableName() string {
	return "terminal_param_sub"
}

// 主键列名
func (t *TerminalParamSub) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalParamSub) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalParamSubPkLoader) NewLoader() *TerminalParamSubPkLoader {
	return &TerminalParamSubPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamSub, []error) {
			var rs []*TerminalParamSub
			var m TerminalParamSub
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalParamSub) UnionPrimaryColumnName() string {
	return "param_sub_id"
}

// 获取联合主键
func (t *TerminalParamSub) GetUnionPrimary() string {
	return t.ParamSubID
}

// 新建联合主键dataloader
func (t *TerminalParamSubUnionPkLoader) NewLoader() *TerminalParamSubUnionPkLoader {
	return &TerminalParamSubUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamSub, []error) {
			var rs []*TerminalParamSub
			var m TerminalParamSub
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
