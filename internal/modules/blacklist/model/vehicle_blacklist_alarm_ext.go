package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleBlacklistAlarmUnionPkLoader string *VehicleSupervision/internal/modules/blacklist/model.VehicleBlacklistAlarm

// 数据库表名
func (t *VehicleBlacklistAlarm) TableName() string {
	return "vehicle_blacklist_alarm"
}

// 主键列名
func (t *VehicleBlacklistAlarm) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *VehicleBlacklistAlarm) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *VehicleBlacklistAlarmPkLoader) NewLoader() *VehicleBlacklistAlarmPkLoader {
	return &VehicleBlacklistAlarmPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleBlacklistAlarm, []error) {
			var rs []*VehicleBlacklistAlarm
			var m VehicleBlacklistAlarm
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *VehicleBlacklistAlarm) UnionPrimaryColumnName() string {
	return "alarm_id"
}

// 获取联合主键
func (t *VehicleBlacklistAlarm) GetUnionPrimary() string {
	return t.AlarmID
}

// 新建联合主键dataloader
func (t *VehicleBlacklistAlarmUnionPkLoader) NewLoader() *VehicleBlacklistAlarmUnionPkLoader {
	return &VehicleBlacklistAlarmUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleBlacklistAlarm, []error) {
			var rs []*VehicleBlacklistAlarm
			var m VehicleBlacklistAlarm
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
