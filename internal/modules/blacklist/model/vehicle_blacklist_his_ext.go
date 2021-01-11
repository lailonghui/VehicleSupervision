package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleBlacklistHisUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.VehicleBlacklistHis

// 数据库表名
func (t *VehicleBlacklistHis) TableName() string {
	return "vehicle_blacklist_his"
}

// 主键列名
func (t *VehicleBlacklistHis) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *VehicleBlacklistHis) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *VehicleBlacklistHisPkLoader) NewLoader() *VehicleBlacklistHisPkLoader {
	return &VehicleBlacklistHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleBlacklistHis, []error) {
			var rs []*VehicleBlacklistHis
			var m VehicleBlacklistHis
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *VehicleBlacklistHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 获取联合主键
func (t *VehicleBlacklistHis) GetUnionPrimary() string {
	return t.HisID
}

// 新建联合主键dataloader
func (t *VehicleBlacklistHisUnionPkLoader) NewLoader() *VehicleBlacklistHisUnionPkLoader {
	return &VehicleBlacklistHisUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleBlacklistHis, []error) {
			var rs []*VehicleBlacklistHis
			var m VehicleBlacklistHis
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
