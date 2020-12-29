package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden RideHailingDriverVerifyLoader string *VehicleSupervision/internal/modules/ridehailing/model.RideHailingDriverVerify

// 数据库表名
func (t RideHailingDriverVerify) TableName() string {
	return "ride_hailing_driver_verify"
}

// 主键列名
func (t RideHailingDriverVerify) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t RideHailingDriverVerify) UnionPrimaryColumnName() string {
	return "ride_hailing_driver_verify_id"
}

// 新建dataloader
func (t *RideHailingDriverVerify) NewLoader() *RideHailingDriverVerifyLoader {
	return &RideHailingDriverVerifyLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*RideHailingDriverVerify, []error) {
			var rs []*RideHailingDriverVerify

			db.DB.Model(&RideHailingDriverVerify{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
