package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalBeidouValidUnionPkLoader string *VehicleSupervision/internal/modules/device/model.TerminalBeidouValid

// 数据库表名
func (t *TerminalBeidouValid) TableName() string {
	return "terminal_beidou_valid"
}

// 主键列名
func (t *TerminalBeidouValid) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *TerminalBeidouValid) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *TerminalBeidouValidPkLoader) NewLoader() *TerminalBeidouValidPkLoader {
	return &TerminalBeidouValidPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalBeidouValid, []error) {
			var rs []*TerminalBeidouValid
			var m TerminalBeidouValid
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *TerminalBeidouValid) UnionPrimaryColumnName() string {
	return "terminal_beidou_valid_id"
}

// 获取联合主键
func (t *TerminalBeidouValid) GetUnionPrimary() string {
	return t.TerminalBeidouValidID
}

// 新建联合主键dataloader
func (t *TerminalBeidouValidUnionPkLoader) NewLoader() *TerminalBeidouValidUnionPkLoader {
	return &TerminalBeidouValidUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalBeidouValid, []error) {
			var rs []*TerminalBeidouValid
			var m TerminalBeidouValid
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
