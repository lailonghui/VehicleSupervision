package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ControlLineTimeLoader string *VehicleSupervision/internal/modules/driving/model.ControlLineTime

// 数据库表名
func (t ControlLineTime) TableName() string {
	return "control_line_time"
}

// 主键列名
func (t ControlLineTime) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t ControlLineTime) UnionPrimaryColumnName() string {
	return "control_line_time_id"
}

// 新建dataloader
func (t *ControlLineTime) NewLoader() *ControlLineTimeLoader {
	return &ControlLineTimeLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLineTime, []error) {
			var rs []*ControlLineTime

			db.DB.Model(&ControlLineTime{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
