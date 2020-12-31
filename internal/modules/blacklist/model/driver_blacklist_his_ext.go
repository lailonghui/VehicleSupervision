package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden DriverBlacklistHisLoader string *VehicleSupervision/internal/modules/blacklist/model.DriverBlacklistHis

// 数据库表名
func (t DriverBlacklistHis) TableName() string {
	return "driver_blacklist_his"
}

// 主键列名
func (t DriverBlacklistHis) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t DriverBlacklistHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 新建dataloader
func (t *DriverBlacklistHis) NewLoader() *DriverBlacklistHisLoader {
	return &DriverBlacklistHisLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DriverBlacklistHis, []error) {
			var rs []*DriverBlacklistHis

			db.DB.Model(&DriverBlacklistHis{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}