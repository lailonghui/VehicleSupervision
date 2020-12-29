package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseScoreLogLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseScoreLog

// 数据库表名
func (t EnterpriseScoreLog) TableName() string {
	return "enterprise_score_log"
}

// 主键列名
func (t EnterpriseScoreLog) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseScoreLog) UnionPrimaryColumnName() string {
	return "log_id"
}

// 新建dataloader
func (t *EnterpriseScoreLog) NewLoader() *EnterpriseScoreLogLoader {
	return &EnterpriseScoreLogLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseScoreLog, []error) {
			var rs []*EnterpriseScoreLog

			db.DB.Model(&EnterpriseScoreLog{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
