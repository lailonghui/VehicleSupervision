package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedLineTimeUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedLineTime

// 数据库表名
func (t *LimitSpeedLineTime) TableName() string {
	return "limit_speed_line_time"
}

// 主键列名
func (t *LimitSpeedLineTime) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *LimitSpeedLineTime) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *LimitSpeedLineTimePkLoader) NewLoader() *LimitSpeedLineTimePkLoader {
	return &LimitSpeedLineTimePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLineTime, []error) {
			var rs []*LimitSpeedLineTime
			var m LimitSpeedLineTime
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *LimitSpeedLineTime) UnionPrimaryColumnName() string {
	return "limit_speed_line_time_id"
}

// 获取联合主键
func (t *LimitSpeedLineTime) GetUnionPrimary() string {
	return t.LimitSpeedLineTimeID
}

// 新建联合主键dataloader
func (t *LimitSpeedLineTimeUnionPkLoader) NewLoader() *LimitSpeedLineTimeUnionPkLoader {
	return &LimitSpeedLineTimeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLineTime, []error) {
			var rs []*LimitSpeedLineTime
			var m LimitSpeedLineTime
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
