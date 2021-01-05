package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseStateHisUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseStateHis

// 数据库表名
func (t EnterpriseStateHis) TableName() string {
	return "enterprise_state_his"
}

// 主键列名
func (t EnterpriseStateHis) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EnterpriseStateHisPkLoader) NewLoader() *EnterpriseStateHisPkLoader {
	return &EnterpriseStateHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseStateHis, []error) {
			var rs []*EnterpriseStateHis
			var m EnterpriseStateHis
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EnterpriseStateHis) UnionPrimaryColumnName() string {
	return "state_his_id"
}

// 新建联合主键dataloader
func (t *EnterpriseStateHisUnionPkLoader) NewLoader() *EnterpriseStateHisUnionPkLoader {
	return &EnterpriseStateHisUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseStateHis, []error) {
			var rs []*EnterpriseStateHis
			var m EnterpriseStateHis
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
