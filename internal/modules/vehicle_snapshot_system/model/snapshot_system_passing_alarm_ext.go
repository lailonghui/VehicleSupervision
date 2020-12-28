package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden SnapshotSystemPassingAlarmLoader string *VehicleSupervision/internal/modules/vehicle_snapshot_system/model.SnapshotSystemPassingAlarm

func (t SnapshotSystemPassingAlarm) TableName() string {
	return "snapshot_system_passing_alarm"
}

func (t *SnapshotSystemPassingAlarm) NewLoader() *SnapshotSystemPassingAlarmLoader {
	return &SnapshotSystemPassingAlarmLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SnapshotSystemPassingAlarm, []error) {
			var rs []*SnapshotSystemPassingAlarm
			// TODO 按实际需要实现
			db.DB.Model(&SnapshotSystemPassingAlarm{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
