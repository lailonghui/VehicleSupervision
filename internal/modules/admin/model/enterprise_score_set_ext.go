package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseScoreSetUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseScoreSet

// 数据库表名
func (t EnterpriseScoreSet) TableName() string {
	return "enterprise_score_set"
}

// 主键列名
func (t EnterpriseScoreSet) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EnterpriseScoreSetPkLoader) NewLoader() *EnterpriseScoreSetPkLoader {
	return &EnterpriseScoreSetPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseScoreSet, []error) {
			var rs []*EnterpriseScoreSet
			var m EnterpriseScoreSet
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EnterpriseScoreSet) UnionPrimaryColumnName() string {
	return "score_set_id"
}

// 新建联合主键dataloader
func (t *EnterpriseScoreSetUnionPkLoader) NewLoader() *EnterpriseScoreSetUnionPkLoader {
	return &EnterpriseScoreSetUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseScoreSet, []error) {
			var rs []*EnterpriseScoreSet
			var m EnterpriseScoreSet
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
