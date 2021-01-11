package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden RideHailingDriverUnionPkLoader string *VehicleSupervision/internal/modules/ridehailing/model.RideHailingDriver

// 数据库表名
func (t *RideHailingDriver) TableName() string {
	return "ride_hailing_driver"
}

// 主键列名
func (t *RideHailingDriver) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *RideHailingDriver) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *RideHailingDriverPkLoader) NewLoader() *RideHailingDriverPkLoader {
	return &RideHailingDriverPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RideHailingDriver, []error) {
			var rs []*RideHailingDriver
			var m RideHailingDriver
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *RideHailingDriver) UnionPrimaryColumnName() string {
	return "ride_hailing_driver_id"
}

// 获取联合主键
func (t *RideHailingDriver) GetUnionPrimary() string {
	return t.RideHailingDriverID
}

// 新建联合主键dataloader
func (t *RideHailingDriverUnionPkLoader) NewLoader() *RideHailingDriverUnionPkLoader {
	return &RideHailingDriverUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RideHailingDriver, []error) {
			var rs []*RideHailingDriver
			var m RideHailingDriver
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
