package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden RideHailingDriverVerifyUnionPkLoader string *VehicleSupervision/internal/modules/ridehailing/model.RideHailingDriverVerify

// 数据库表名
func (t *RideHailingDriverVerify) TableName() string {
	return "ride_hailing_driver_verify"
}

// 主键列名
func (t *RideHailingDriverVerify) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *RideHailingDriverVerify) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *RideHailingDriverVerifyPkLoader) NewLoader() *RideHailingDriverVerifyPkLoader {
	return &RideHailingDriverVerifyPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RideHailingDriverVerify, []error) {
			var rs []*RideHailingDriverVerify
			var m RideHailingDriverVerify
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *RideHailingDriverVerify) UnionPrimaryColumnName() string {
	return "ride_hailing_driver_verify_id"
}

// 获取联合主键
func (t *RideHailingDriverVerify) GetUnionPrimary() string {
	return t.RideHailingDriverVerifyID
}

// 新建联合主键dataloader
func (t *RideHailingDriverVerifyUnionPkLoader) NewLoader() *RideHailingDriverVerifyUnionPkLoader {
	return &RideHailingDriverVerifyUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RideHailingDriverVerify, []error) {
			var rs []*RideHailingDriverVerify
			var m RideHailingDriverVerify
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
