package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden RideHailingDriverUnionPkLoader string *VehicleSupervision/internal/modules/ridehailing/model.RideHailingDriver

// 数据库表名
func (t RideHailingDriver) TableName() string {
	return "ride_hailing_driver"
}

// 主键列名
func (t RideHailingDriver) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *RideHailingDriver) NewPkLoader() *RideHailingDriverPkLoader {
	return &RideHailingDriverPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RideHailingDriver, []error) {
			var rs []*RideHailingDriver
			db.DB.Model(&RideHailingDriver{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t RideHailingDriver) UnionPrimaryColumnName() string {
	return "ride_hailing_driver_id"
}

// 新建联合主键dataloader
func (t *RideHailingDriver) NewUnionPkLoader() *RideHailingDriverUnionPkLoader {
	return &RideHailingDriverUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RideHailingDriver, []error) {
			var rs []*RideHailingDriver
			db.DB.Model(&RideHailingDriver{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
