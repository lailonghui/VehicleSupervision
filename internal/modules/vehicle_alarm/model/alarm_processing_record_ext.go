package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden AlarmProcessingRecordLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.AlarmProcessingRecord

func (t AlarmProcessingRecord) TableName() string {
	return "alarm_processing_record"
}

func (t *AlarmProcessingRecord) NewLoader() *AlarmProcessingRecordLoader {
	return &AlarmProcessingRecordLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*AlarmProcessingRecord, []error) {
			var rs []*AlarmProcessingRecord
			// TODO 按实际需要实现
			db.DB.Model(&AlarmProcessingRecord{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
