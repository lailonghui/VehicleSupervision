package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamSubLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamSub

// 数据库表名
func (t TerminalParamSub) TableName() string {
	return "terminal_param_sub"
}

// 主键列名
func (t TerminalParamSub) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalParamSub) UnionPrimaryColumnName() string {
	return "param_sub_id"
}

// 新建dataloader
func (t *TerminalParamSub) NewLoader() *TerminalParamSubLoader {
	return &TerminalParamSubLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamSub, []error) {
			var rs []*TerminalParamSub

			db.DB.Model(&TerminalParamSub{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
