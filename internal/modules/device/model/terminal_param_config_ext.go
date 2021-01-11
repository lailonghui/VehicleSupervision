package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamConfigUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamConfig

// 数据库表名
func (t *TerminalParamConfig) TableName() string {
	return "terminal_param_config"
}

// 主键列名
func (t *TerminalParamConfig) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalParamConfig) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalParamConfigPkLoader) NewLoader() *TerminalParamConfigPkLoader {
	return &TerminalParamConfigPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamConfig, []error) {
			var rs []*TerminalParamConfig
			var m TerminalParamConfig
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalParamConfig) UnionPrimaryColumnName() string {
	return "config_id"
}

// 获取联合主键
func (t *TerminalParamConfig) GetUnionPrimary() string {
	return t.ConfigID
}

// 新建联合主键dataloader
func (t *TerminalParamConfigUnionPkLoader) NewLoader() *TerminalParamConfigUnionPkLoader {
	return &TerminalParamConfigUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamConfig, []error) {
			var rs []*TerminalParamConfig
			var m TerminalParamConfig
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
