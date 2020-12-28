package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden DistrictAlarmContentPushLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.DistrictAlarmContentPush

func (t DistrictAlarmContentPush) TableName() string {
	return "district_alarm_content_push"
}

func (t *DistrictAlarmContentPush) NewLoader() *DistrictAlarmContentPushLoader {
	return &DistrictAlarmContentPushLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*DistrictAlarmContentPush, []error) {
			var rs []*DistrictAlarmContentPush
			// TODO 按实际需要实现
			db.DB.Model(&DistrictAlarmContentPush{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
