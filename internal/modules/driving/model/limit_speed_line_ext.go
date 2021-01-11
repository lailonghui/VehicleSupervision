package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedLineUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedLine

// 数据库表名
func (t *LimitSpeedLine) TableName() string {
	return "limit_speed_line"
}

// 主键列名
func (t *LimitSpeedLine) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *LimitSpeedLine) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *LimitSpeedLinePkLoader) NewLoader() *LimitSpeedLinePkLoader {
	return &LimitSpeedLinePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLine, []error) {
			var rs []*LimitSpeedLine
			var m LimitSpeedLine
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *LimitSpeedLine) UnionPrimaryColumnName() string {
	return "limit_speed_line_id"
}

// 获取联合主键
func (t *LimitSpeedLine) GetUnionPrimary() string {
	return t.LimitSpeedLineID
}

// 新建联合主键dataloader
func (t *LimitSpeedLineUnionPkLoader) NewLoader() *LimitSpeedLineUnionPkLoader {
	return &LimitSpeedLineUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLine, []error) {
			var rs []*LimitSpeedLine
			var m LimitSpeedLine
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
