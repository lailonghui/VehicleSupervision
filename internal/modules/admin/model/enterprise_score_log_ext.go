package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseScoreLogUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseScoreLog

// 数据库表名
func (t EnterpriseScoreLog) TableName() string {
	return "enterprise_score_log"
}

// 主键列名
func (t EnterpriseScoreLog) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EnterpriseScoreLogPkLoader) NewLoader() *EnterpriseScoreLogPkLoader {
	return &EnterpriseScoreLogPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseScoreLog, []error) {
			var rs []*EnterpriseScoreLog
			var m EnterpriseScoreLog
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EnterpriseScoreLog) UnionPrimaryColumnName() string {
	return "log_id"
}

// 新建联合主键dataloader
func (t *EnterpriseScoreLogUnionPkLoader) NewLoader() *EnterpriseScoreLogUnionPkLoader {
	return &EnterpriseScoreLogUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseScoreLog, []error) {
			var rs []*EnterpriseScoreLog
			var m EnterpriseScoreLog
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
