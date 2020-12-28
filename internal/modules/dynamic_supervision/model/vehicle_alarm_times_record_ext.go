package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleAlarmTimesRecordLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.VehicleAlarmTimesRecord

func (t VehicleAlarmTimesRecord) TableName() string {
	return "VehicleAlarmTimesRecord"
}

func (t *VehicleAlarmTimesRecord) NewLoader() *VehicleAlarmTimesRecordLoader {
	return &VehicleAlarmTimesRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleAlarmTimesRecord, []error) {
			var rs []*VehicleAlarmTimesRecord
			// TODO 按实际需要实现
			db.DB.Model(&VehicleAlarmTimesRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
