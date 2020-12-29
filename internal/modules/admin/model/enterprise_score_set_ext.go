package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseScoreSetLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseScoreSet

// 数据库表名
func (t EnterpriseScoreSet) TableName() string {
	return "enterprise_score_set"
}

// 主键列名
func (t EnterpriseScoreSet) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseScoreSet) UnionPrimaryColumnName() string {
	return "score_set_id"
}

// 新建dataloader
func (t *EnterpriseScoreSet) NewLoader() *EnterpriseScoreSetLoader {
	return &EnterpriseScoreSetLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseScoreSet, []error) {
			var rs []*EnterpriseScoreSet

			db.DB.Model(&EnterpriseScoreSet{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
