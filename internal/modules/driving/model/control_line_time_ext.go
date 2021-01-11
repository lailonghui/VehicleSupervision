package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ControlLineTimeUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.ControlLineTime

// 数据库表名
func (t *ControlLineTime) TableName() string {
	return "control_line_time"
}

// 主键列名
func (t *ControlLineTime) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *ControlLineTime) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *ControlLineTimePkLoader) NewLoader() *ControlLineTimePkLoader {
	return &ControlLineTimePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLineTime, []error) {
			var rs []*ControlLineTime
			var m ControlLineTime
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *ControlLineTime) UnionPrimaryColumnName() string {
	return "control_line_time_id"
}

// 获取联合主键
func (t *ControlLineTime) GetUnionPrimary() string {
	return t.ControlLineTimeID
}

// 新建联合主键dataloader
func (t *ControlLineTimeUnionPkLoader) NewLoader() *ControlLineTimeUnionPkLoader {
	return &ControlLineTimeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ControlLineTime, []error) {
			var rs []*ControlLineTime
			var m ControlLineTime
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
