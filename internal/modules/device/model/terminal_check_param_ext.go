package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalCheckParamUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalCheckParam

// 数据库表名
func (t *TerminalCheckParam) TableName() string {
	return "terminal_check_param"
}

// 主键列名
func (t *TerminalCheckParam) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalCheckParam) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalCheckParamPkLoader) NewLoader() *TerminalCheckParamPkLoader {
	return &TerminalCheckParamPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheckParam, []error) {
			var rs []*TerminalCheckParam
			var m TerminalCheckParam
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalCheckParam) UnionPrimaryColumnName() string {
	return "terminal_check_param_id"
}

// 获取联合主键
func (t *TerminalCheckParam) GetUnionPrimary() string {
	return t.TerminalCheckParamID
}

// 新建联合主键dataloader
func (t *TerminalCheckParamUnionPkLoader) NewLoader() *TerminalCheckParamUnionPkLoader {
	return &TerminalCheckParamUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalCheckParam, []error) {
			var rs []*TerminalCheckParam
			var m TerminalCheckParam
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
