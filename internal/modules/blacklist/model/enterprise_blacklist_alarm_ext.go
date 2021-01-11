package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseBlacklistAlarmUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.EnterpriseBlacklistAlarm

// 数据库表名
func (t *EnterpriseBlacklistAlarm) TableName() string {
	return "enterprise_blacklist_alarm"
}

// 主键列名
func (t *EnterpriseBlacklistAlarm) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EnterpriseBlacklistAlarm) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EnterpriseBlacklistAlarmPkLoader) NewLoader() *EnterpriseBlacklistAlarmPkLoader {
	return &EnterpriseBlacklistAlarmPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseBlacklistAlarm, []error) {
			var rs []*EnterpriseBlacklistAlarm
			var m EnterpriseBlacklistAlarm
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EnterpriseBlacklistAlarm) UnionPrimaryColumnName() string {
	return "alarm_id"
}

// 获取联合主键
func (t *EnterpriseBlacklistAlarm) GetUnionPrimary() string {
	return t.AlarmID
}

// 新建联合主键dataloader
func (t *EnterpriseBlacklistAlarmUnionPkLoader) NewLoader() *EnterpriseBlacklistAlarmUnionPkLoader {
	return &EnterpriseBlacklistAlarmUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseBlacklistAlarm, []error) {
			var rs []*EnterpriseBlacklistAlarm
			var m EnterpriseBlacklistAlarm
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
