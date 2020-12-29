package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleBlacklistAlarmLoader string *VehicleSupervision/internal/modules/blacklist/model.VehicleBlacklistAlarm

// 数据库表名
func (t VehicleBlacklistAlarm) TableName() string {
	return "vehicle_blacklist_alarm"
}

// 主键列名
func (t VehicleBlacklistAlarm) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t VehicleBlacklistAlarm) UnionPrimaryColumnName() string {
	return "alarm_id"
}

// 新建dataloader
func (t *VehicleBlacklistAlarm) NewLoader() *VehicleBlacklistAlarmLoader {
	return &VehicleBlacklistAlarmLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleBlacklistAlarm, []error) {
			var rs []*VehicleBlacklistAlarm

			db.DB.Model(&VehicleBlacklistAlarm{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
