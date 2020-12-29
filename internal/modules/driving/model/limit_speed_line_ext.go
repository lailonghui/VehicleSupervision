package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedLineLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedLine

// 数据库表名
func (t LimitSpeedLine) TableName() string {
	return "limit_speed_line"
}

// 主键列名
func (t LimitSpeedLine) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t LimitSpeedLine) UnionPrimaryColumnName() string {
	return "limit_speed_line_id"
}

// 新建dataloader
func (t *LimitSpeedLine) NewLoader() *LimitSpeedLineLoader {
	return &LimitSpeedLineLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedLine, []error) {
			var rs []*LimitSpeedLine

			db.DB.Model(&LimitSpeedLine{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
