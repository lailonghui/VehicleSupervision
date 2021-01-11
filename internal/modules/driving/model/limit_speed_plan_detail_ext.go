package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden LimitSpeedPlanDetailUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.LimitSpeedPlanDetail

// 数据库表名
func (t *LimitSpeedPlanDetail) TableName() string {
	return "limit_speed_plan_detail"
}

// 主键列名
func (t *LimitSpeedPlanDetail) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *LimitSpeedPlanDetail) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *LimitSpeedPlanDetailPkLoader) NewLoader() *LimitSpeedPlanDetailPkLoader {
	return &LimitSpeedPlanDetailPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlanDetail, []error) {
			var rs []*LimitSpeedPlanDetail
			var m LimitSpeedPlanDetail
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *LimitSpeedPlanDetail) UnionPrimaryColumnName() string {
	return "detail_id"
}

// 获取联合主键
func (t *LimitSpeedPlanDetail) GetUnionPrimary() string {
	return t.DetailID
}

// 新建联合主键dataloader
func (t *LimitSpeedPlanDetailUnionPkLoader) NewLoader() *LimitSpeedPlanDetailUnionPkLoader {
	return &LimitSpeedPlanDetailUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*LimitSpeedPlanDetail, []error) {
			var rs []*LimitSpeedPlanDetail
			var m LimitSpeedPlanDetail
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
