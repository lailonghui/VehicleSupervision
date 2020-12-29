package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedLineTimeLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedLineTime

// 数据库表名
func (t LimitSpeedLineTime) TableName() string {
	return "limit_speed_line_time"
}

// 主键列名
func (t LimitSpeedLineTime) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t LimitSpeedLineTime) UnionPrimaryColumnName() string {
	return "limit_speed_line_time_id"
}

// 新建dataloader
func (t *LimitSpeedLineTime) NewLoader() *LimitSpeedLineTimeLoader {
	return &LimitSpeedLineTimeLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLineTime, []error) {
			var rs []*LimitSpeedLineTime

			db.DB.Model(&LimitSpeedLineTime{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
