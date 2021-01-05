package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamTypeUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamType

// 数据库表名
func (t TerminalParamType) TableName() string {
	return "terminal_param_type"
}

// 主键列名
func (t TerminalParamType) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *TerminalParamType) NewPkLoader() *TerminalParamTypePkLoader {
	return &TerminalParamTypePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamType, []error) {
			var rs []*TerminalParamType
			db.DB.Model(&TerminalParamType{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t TerminalParamType) UnionPrimaryColumnName() string {
	return "param_type_id"
}

// 新建联合主键dataloader
func (t *TerminalParamType) NewUnionPkLoader() *TerminalParamTypeUnionPkLoader {
	return &TerminalParamTypeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamType, []error) {
			var rs []*TerminalParamType
			db.DB.Model(&TerminalParamType{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
