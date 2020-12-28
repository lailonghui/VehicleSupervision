package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden EnterpriseAlarmSendPoliceLoader string *VehicleSupervision/internal/modules/vehicle_alarm/model.EnterpriseAlarmSendPolice

func (t EnterpriseAlarmSendPolice) TableName() string {
	return "enterprise_alarm_send_police"
}

func (t *EnterpriseAlarmSendPolice) NewLoader() *EnterpriseAlarmSendPoliceLoader {
	return &EnterpriseAlarmSendPoliceLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EnterpriseAlarmSendPolice, []error) {
			var rs []*EnterpriseAlarmSendPolice
			// TODO 按实际需要实现
			db.DB.Model(&EnterpriseAlarmSendPolice{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
