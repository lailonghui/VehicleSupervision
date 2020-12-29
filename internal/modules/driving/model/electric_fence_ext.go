package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden ElectricFenceLoader string *VehicleSupervision/internal/modules/driving/model.ElectricFence

// 数据库表名
func (t ElectricFence) TableName() string {
	return "electric_fence"
}

// 主键列名
func (t ElectricFence) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t ElectricFence) UnionPrimaryColumnName() string {
	return "electric_fence_id"
}

// 新建dataloader
func (t *ElectricFence) NewLoader() *ElectricFenceLoader {
	return &ElectricFenceLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ElectricFence, []error) {
			var rs []*ElectricFence

			db.DB.Model(&ElectricFence{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
