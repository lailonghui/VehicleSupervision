package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VoiceAlarmRecordLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.VoiceAlarmRecord

func (t VoiceAlarmRecord) TableName() string {
	return "VoiceAlarmRecord"
}

func (t *VoiceAlarmRecord) NewLoader() *VoiceAlarmRecordLoader {
	return &VoiceAlarmRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VoiceAlarmRecord, []error) {
			var rs []*VoiceAlarmRecord
			// TODO 按实际需要实现
			db.DB.Model(&VoiceAlarmRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
