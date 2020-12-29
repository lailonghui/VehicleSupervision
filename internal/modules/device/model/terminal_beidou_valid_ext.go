package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden TerminalBeidouValidLoader string *VehicleSupervision/internal/modules/device/model.TerminalBeidouValid

// 数据库表名
func (t TerminalBeidouValid) TableName() string {
	return "terminal_beidou_valid"
}

// 主键列名
func (t TerminalBeidouValid) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t TerminalBeidouValid) UnionPrimaryColumnName() string {
	return "terminal_beidou_valid_id"
}

// 新建dataloader
func (t *TerminalBeidouValid) NewLoader() *TerminalBeidouValidLoader {
	return &TerminalBeidouValidLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*TerminalBeidouValid, []error) {
			var rs []*TerminalBeidouValid

			db.DB.Model(&TerminalBeidouValid{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
