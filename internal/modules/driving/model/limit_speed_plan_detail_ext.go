package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedPlanDetailLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedPlanDetail

// 数据库表名
func (t LimitSpeedPlanDetail) TableName() string {
	return "limit_speed_plan_detail"
}

// 主键列名
func (t LimitSpeedPlanDetail) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t LimitSpeedPlanDetail) UnionPrimaryColumnName() string {
	return "detail_id"
}

// 新建dataloader
func (t *LimitSpeedPlanDetail) NewLoader() *LimitSpeedPlanDetailLoader {
	return &LimitSpeedPlanDetailLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlanDetail, []error) {
			var rs []*LimitSpeedPlanDetail

			db.DB.Model(&LimitSpeedPlanDetail{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}