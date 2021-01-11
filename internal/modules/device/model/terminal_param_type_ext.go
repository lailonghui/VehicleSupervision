package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamTypeUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamType

// 数据库表名
func (t *TerminalParamType) TableName() string {
	return "terminal_param_type"
}

// 主键列名
func (t *TerminalParamType) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalParamType) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalParamTypePkLoader) NewLoader() *TerminalParamTypePkLoader {
	return &TerminalParamTypePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamType, []error) {
			var rs []*TerminalParamType
			var m TerminalParamType
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalParamType) UnionPrimaryColumnName() string {
	return "param_type_id"
}

// 获取联合主键
func (t *TerminalParamType) GetUnionPrimary() string {
	return t.ParamTypeID
}

// 新建联合主键dataloader
func (t *TerminalParamTypeUnionPkLoader) NewLoader() *TerminalParamTypeUnionPkLoader {
	return &TerminalParamTypeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamType, []error) {
			var rs []*TerminalParamType
			var m TerminalParamType
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
