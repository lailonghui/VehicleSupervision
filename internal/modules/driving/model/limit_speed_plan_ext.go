package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedPlanLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedPlan

// 数据库表名
func (t LimitSpeedPlan) TableName() string {
	return "limit_speed_plan"
}

// 主键列名
func (t LimitSpeedPlan) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t LimitSpeedPlan) UnionPrimaryColumnName() string {
	return "limit_speed_plan_id"
}

// 新建dataloader
func (t *LimitSpeedPlan) NewLoader() *LimitSpeedPlanLoader {
	return &LimitSpeedPlanLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlan, []error) {
			var rs []*LimitSpeedPlan

			db.DB.Model(&LimitSpeedPlan{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
