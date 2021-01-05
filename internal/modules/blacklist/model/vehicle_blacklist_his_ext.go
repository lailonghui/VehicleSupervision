package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleBlacklistHisUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.VehicleBlacklistHis

// 数据库表名
func (t VehicleBlacklistHis) TableName() string {
	return "vehicle_blacklist_his"
}

// 主键列名
func (t VehicleBlacklistHis) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *VehicleBlacklistHis) NewPkLoader() *VehicleBlacklistHisPkLoader {
	return &VehicleBlacklistHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleBlacklistHis, []error) {
			var rs []*VehicleBlacklistHis
			db.DB.Model(&VehicleBlacklistHis{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t VehicleBlacklistHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 新建联合主键dataloader
func (t *VehicleBlacklistHis) NewUnionPkLoader() *VehicleBlacklistHisUnionPkLoader {
	return &VehicleBlacklistHisUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleBlacklistHis, []error) {
			var rs []*VehicleBlacklistHis
			db.DB.Model(&VehicleBlacklistHis{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
