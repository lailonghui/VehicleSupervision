package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden AppEnforcementLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.AppEnforcement

func (t AppEnforcement) TableName() string {
	return "AppEnforcement"
}

func (t *AppEnforcement) NewLoader() *AppEnforcementLoader {
	return &AppEnforcementLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*AppEnforcement, []error) {
			var rs []*AppEnforcement
			// TODO 按实际需要实现
			db.DB.Model(&AppEnforcement{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
