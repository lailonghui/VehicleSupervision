package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalCheckParamLoader string *VehicleSupervision/internal/modules/device/model.TerminalCheckParam

// 数据库表名
func (t TerminalCheckParam) TableName() string {
	return "terminal_check_param"
}

// 主键列名
func (t TerminalCheckParam) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalCheckParam) UnionPrimaryColumnName() string {
	return "terminal_check_param_id"
}

// 新建dataloader
func (t *TerminalCheckParam) NewLoader() *TerminalCheckParamLoader {
	return &TerminalCheckParamLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheckParam, []error) {
			var rs []*TerminalCheckParam

			db.DB.Model(&TerminalCheckParam{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
