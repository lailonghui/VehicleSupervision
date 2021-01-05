package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ControlLineTimeUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.ControlLineTime

// 数据库表名
func (t ControlLineTime) TableName() string {
	return "control_line_time"
}

// 主键列名
func (t ControlLineTime) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *ControlLineTime) NewPkLoader() *ControlLineTimePkLoader {
	return &ControlLineTimePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLineTime, []error) {
			var rs []*ControlLineTime
			db.DB.Model(&ControlLineTime{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t ControlLineTime) UnionPrimaryColumnName() string {
	return "control_line_time_id"
}

// 新建联合主键dataloader
func (t *ControlLineTime) NewUnionPkLoader() *ControlLineTimeUnionPkLoader {
	return &ControlLineTimeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLineTime, []error) {
			var rs []*ControlLineTime
			db.DB.Model(&ControlLineTime{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
