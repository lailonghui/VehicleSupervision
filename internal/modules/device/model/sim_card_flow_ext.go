package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardFlowUnionPkLoader string *VehicleSupervision/internal/modules/device/model.SimCardFlow

// 数据库表名
func (t *SimCardFlow) TableName() string {
	return "sim_card_flow"
}

// 主键列名
func (t *SimCardFlow) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *SimCardFlow) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *SimCardFlowPkLoader) NewLoader() *SimCardFlowPkLoader {
	return &SimCardFlowPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardFlow, []error) {
			var rs []*SimCardFlow
			var m SimCardFlow
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *SimCardFlow) UnionPrimaryColumnName() string {
	return "sim_card_flow_id"
}

// 获取联合主键
func (t *SimCardFlow) GetUnionPrimary() string {
	return t.SimCardFlowID
}

// 新建联合主键dataloader
func (t *SimCardFlowUnionPkLoader) NewLoader() *SimCardFlowUnionPkLoader {
	return &SimCardFlowUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardFlow, []error) {
			var rs []*SimCardFlow
			var m SimCardFlow
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
