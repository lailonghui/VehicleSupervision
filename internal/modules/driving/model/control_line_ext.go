package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ControlLineLoader string *VehicleSupervision/internal/modules/driving/model.ControlLine

// 数据库表名
func (t ControlLine) TableName() string {
	return "control_line"
}

// 主键列名
func (t ControlLine) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t ControlLine) UnionPrimaryColumnName() string {
	return "control_line_id"
}

// 新建dataloader
func (t *ControlLine) NewLoader() *ControlLineLoader {
	return &ControlLineLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLine, []error) {
			var rs []*ControlLine

			db.DB.Model(&ControlLine{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
