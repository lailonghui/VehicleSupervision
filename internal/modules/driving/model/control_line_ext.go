package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ControlLineUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.ControlLine

// 数据库表名
func (t ControlLine) TableName() string {
	return "control_line"
}

// 主键列名
func (t ControlLine) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *ControlLine) NewPkLoader() *ControlLinePkLoader {
	return &ControlLinePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLine, []error) {
			var rs []*ControlLine
			db.DB.Model(&ControlLine{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t ControlLine) UnionPrimaryColumnName() string {
	return "control_line_id"
}

// 新建联合主键dataloader
func (t *ControlLine) NewUnionPkLoader() *ControlLineUnionPkLoader {
	return &ControlLineUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLine, []error) {
			var rs []*ControlLine
			db.DB.Model(&ControlLine{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
