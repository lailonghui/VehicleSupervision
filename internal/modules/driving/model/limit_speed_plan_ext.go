package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedPlanUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedPlan

// 数据库表名
func (t *LimitSpeedPlan) TableName() string {
	return "limit_speed_plan"
}

// 主键列名
func (t *LimitSpeedPlan) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *LimitSpeedPlan) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *LimitSpeedPlanPkLoader) NewLoader() *LimitSpeedPlanPkLoader {
	return &LimitSpeedPlanPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlan, []error) {
			var rs []*LimitSpeedPlan
			var m LimitSpeedPlan
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *LimitSpeedPlan) UnionPrimaryColumnName() string {
	return "limit_speed_plan_id"
}

// 获取联合主键
func (t *LimitSpeedPlan) GetUnionPrimary() string {
	return t.LimitSpeedPlanID
}

// 新建联合主键dataloader
func (t *LimitSpeedPlanUnionPkLoader) NewLoader() *LimitSpeedPlanUnionPkLoader {
	return &LimitSpeedPlanUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlan, []error) {
			var rs []*LimitSpeedPlan
			var m LimitSpeedPlan
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
