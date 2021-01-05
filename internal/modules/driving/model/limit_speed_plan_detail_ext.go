package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedPlanDetailUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedPlanDetail

// 数据库表名
func (t LimitSpeedPlanDetail) TableName() string {
	return "limit_speed_plan_detail"
}

// 主键列名
func (t LimitSpeedPlanDetail) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *LimitSpeedPlanDetail) NewPkLoader() *LimitSpeedPlanDetailPkLoader {
	return &LimitSpeedPlanDetailPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlanDetail, []error) {
			var rs []*LimitSpeedPlanDetail
			db.DB.Model(&LimitSpeedPlanDetail{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t LimitSpeedPlanDetail) UnionPrimaryColumnName() string {
	return "detail_id"
}

// 新建联合主键dataloader
func (t *LimitSpeedPlanDetail) NewUnionPkLoader() *LimitSpeedPlanDetailUnionPkLoader {
	return &LimitSpeedPlanDetailUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlanDetail, []error) {
			var rs []*LimitSpeedPlanDetail
			db.DB.Model(&LimitSpeedPlanDetail{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
