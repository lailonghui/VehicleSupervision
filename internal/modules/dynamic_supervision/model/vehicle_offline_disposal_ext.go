package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden VehicleOfflineDisposalLoader string *VehicleSupervision/internal/modules/dynamic_supervision/model.VehicleOfflineDisposal

func (t VehicleOfflineDisposal) TableName() string {
	return "VehicleOfflineDisposal"
}

func (t *VehicleOfflineDisposal) NewLoader() *VehicleOfflineDisposalLoader {
	return &VehicleOfflineDisposalLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleOfflineDisposal, []error) {
			var rs []*VehicleOfflineDisposal
			// TODO 按实际需要实现
			db.DB.Model(&VehicleOfflineDisposal{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
