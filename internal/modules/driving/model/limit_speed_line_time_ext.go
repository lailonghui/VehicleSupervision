package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedLineTimeUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedLineTime

// 数据库表名
func (t LimitSpeedLineTime) TableName() string {
	return "limit_speed_line_time"
}

// 主键列名
func (t LimitSpeedLineTime) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *LimitSpeedLineTime) NewPkLoader() *LimitSpeedLineTimePkLoader {
	return &LimitSpeedLineTimePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLineTime, []error) {
			var rs []*LimitSpeedLineTime
			db.DB.Model(&LimitSpeedLineTime{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t LimitSpeedLineTime) UnionPrimaryColumnName() string {
	return "limit_speed_line_time_id"
}

// 新建联合主键dataloader
func (t *LimitSpeedLineTime) NewUnionPkLoader() *LimitSpeedLineTimeUnionPkLoader {
	return &LimitSpeedLineTimeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLineTime, []error) {
			var rs []*LimitSpeedLineTime
			db.DB.Model(&LimitSpeedLineTime{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
