package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseBlacklistHisLoader string *VehicleSupervision/internal/modules/blacklist/model.EnterpriseBlacklistHis

// 数据库表名
func (t EnterpriseBlacklistHis) TableName() string {
	return "enterprise_blacklist_his"
}

// 主键列名
func (t EnterpriseBlacklistHis) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseBlacklistHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 新建dataloader
func (t *EnterpriseBlacklistHis) NewLoader() *EnterpriseBlacklistHisLoader {
	return &EnterpriseBlacklistHisLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseBlacklistHis, []error) {
			var rs []*EnterpriseBlacklistHis

			db.DB.Model(&EnterpriseBlacklistHis{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
