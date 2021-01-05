package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalParamConfigUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalParamConfig

// 数据库表名
func (t TerminalParamConfig) TableName() string {
	return "terminal_param_config"
}

// 主键列名
func (t TerminalParamConfig) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *TerminalParamConfig) NewPkLoader() *TerminalParamConfigPkLoader {
	return &TerminalParamConfigPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamConfig, []error) {
			var rs []*TerminalParamConfig
			db.DB.Model(&TerminalParamConfig{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t TerminalParamConfig) UnionPrimaryColumnName() string {
	return "config_id"
}

// 新建联合主键dataloader
func (t *TerminalParamConfig) NewUnionPkLoader() *TerminalParamConfigUnionPkLoader {
	return &TerminalParamConfigUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalParamConfig, []error) {
			var rs []*TerminalParamConfig
			db.DB.Model(&TerminalParamConfig{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
