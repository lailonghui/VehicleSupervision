package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseStateHisLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseStateHis

// 数据库表名
func (t EnterpriseStateHis) TableName() string {
	return "enterprise_state_his"
}

// 主键列名
func (t EnterpriseStateHis) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseStateHis) UnionPrimaryColumnName() string {
	return "state_his_id"
}

// 新建dataloader
func (t *EnterpriseStateHis) NewLoader() *EnterpriseStateHisLoader {
	return &EnterpriseStateHisLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseStateHis, []error) {
			var rs []*EnterpriseStateHis

			db.DB.Model(&EnterpriseStateHis{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
