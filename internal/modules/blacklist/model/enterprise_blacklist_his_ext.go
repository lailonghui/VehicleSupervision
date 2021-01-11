package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseBlacklistHisUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.EnterpriseBlacklistHis

// 数据库表名
func (t *EnterpriseBlacklistHis) TableName() string {
	return "enterprise_blacklist_his"
}

// 主键列名
func (t *EnterpriseBlacklistHis) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EnterpriseBlacklistHis) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EnterpriseBlacklistHisPkLoader) NewLoader() *EnterpriseBlacklistHisPkLoader {
	return &EnterpriseBlacklistHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseBlacklistHis, []error) {
			var rs []*EnterpriseBlacklistHis
			var m EnterpriseBlacklistHis
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EnterpriseBlacklistHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 获取联合主键
func (t *EnterpriseBlacklistHis) GetUnionPrimary() string {
	return t.HisID
}

// 新建联合主键dataloader
func (t *EnterpriseBlacklistHisUnionPkLoader) NewLoader() *EnterpriseBlacklistHisUnionPkLoader {
	return &EnterpriseBlacklistHisUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseBlacklistHis, []error) {
			var rs []*EnterpriseBlacklistHis
			var m EnterpriseBlacklistHis
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
