package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseBlacklistAlarmLoader string *VehicleSupervision/internal/modules/blacklist/model.EnterpriseBlacklistAlarm

// 数据库表名
func (t EnterpriseBlacklistAlarm) TableName() string {
	return "enterprise_blacklist_alarm"
}

// 主键列名
func (t EnterpriseBlacklistAlarm) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EnterpriseBlacklistAlarm) UnionPrimaryColumnName() string {
	return "alarm_id"
}

// 新建dataloader
func (t *EnterpriseBlacklistAlarm) NewLoader() *EnterpriseBlacklistAlarmLoader {
	return &EnterpriseBlacklistAlarmLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseBlacklistAlarm, []error) {
			var rs []*EnterpriseBlacklistAlarm

			db.DB.Model(&EnterpriseBlacklistAlarm{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
