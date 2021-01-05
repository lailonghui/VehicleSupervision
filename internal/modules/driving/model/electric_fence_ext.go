package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ElectricFenceUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.ElectricFence

// 数据库表名
func (t ElectricFence) TableName() string {
	return "electric_fence"
}

// 主键列名
func (t ElectricFence) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *ElectricFence) NewPkLoader() *ElectricFencePkLoader {
	return &ElectricFencePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFence, []error) {
			var rs []*ElectricFence
			db.DB.Model(&ElectricFence{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t ElectricFence) UnionPrimaryColumnName() string {
	return "electric_fence_id"
}

// 新建联合主键dataloader
func (t *ElectricFence) NewUnionPkLoader() *ElectricFenceUnionPkLoader {
	return &ElectricFenceUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFence, []error) {
			var rs []*ElectricFence
			db.DB.Model(&ElectricFence{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
