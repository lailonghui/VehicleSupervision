package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VioCodewfdmLoader string *VehicleSupervision/internal/modules/vehicle_violation/model.VioCodewfdm

func (t VioCodewfdm) TableName() string {
	return "VioCodewfdm"
}

func (t *VioCodewfdm) NewLoader() *VioCodewfdmLoader {
	return &VioCodewfdmLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VioCodewfdm, []error) {
			var rs []*VioCodewfdm
			// TODO 按实际需要实现
			db.DB.Model(&VioCodewfdm{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
