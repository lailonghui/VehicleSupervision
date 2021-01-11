package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EnterpriseAlarmSendPoliceUnionPkLoader string *VehicleSupervision/internal/modules/admin/model.EnterpriseAlarmSendPolice

// 数据库表名
func (t *EnterpriseAlarmSendPolice) TableName() string {
	return "enterprise_alarm_send_police"
}

// 主键列名
func (t *EnterpriseAlarmSendPolice) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EnterpriseAlarmSendPolice) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EnterpriseAlarmSendPolicePkLoader) NewLoader() *EnterpriseAlarmSendPolicePkLoader {
	return &EnterpriseAlarmSendPolicePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseAlarmSendPolice, []error) {
			var rs []*EnterpriseAlarmSendPolice
			var m EnterpriseAlarmSendPolice
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EnterpriseAlarmSendPolice) UnionPrimaryColumnName() string {
	return "enterprise_alarm_send_police_id"
}

// 获取联合主键
func (t *EnterpriseAlarmSendPolice) GetUnionPrimary() string {
	return t.EnterpriseAlarmSendPoliceID
}

// 新建联合主键dataloader
func (t *EnterpriseAlarmSendPoliceUnionPkLoader) NewLoader() *EnterpriseAlarmSendPoliceUnionPkLoader {
	return &EnterpriseAlarmSendPoliceUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseAlarmSendPolice, []error) {
			var rs []*EnterpriseAlarmSendPolice
			var m EnterpriseAlarmSendPolice
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
