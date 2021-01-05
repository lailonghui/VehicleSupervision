package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedPlanUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedPlan

// 数据库表名
func (t LimitSpeedPlan) TableName() string {
	return "limit_speed_plan"
}

// 主键列名
func (t LimitSpeedPlan) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *LimitSpeedPlan) NewPkLoader() *LimitSpeedPlanPkLoader {
	return &LimitSpeedPlanPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlan, []error) {
			var rs []*LimitSpeedPlan
			db.DB.Model(&LimitSpeedPlan{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t LimitSpeedPlan) UnionPrimaryColumnName() string {
	return "limit_speed_plan_id"
}

// 新建联合主键dataloader
func (t *LimitSpeedPlan) NewUnionPkLoader() *LimitSpeedPlanUnionPkLoader {
	return &LimitSpeedPlanUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlan, []error) {
			var rs []*LimitSpeedPlan
			db.DB.Model(&LimitSpeedPlan{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
