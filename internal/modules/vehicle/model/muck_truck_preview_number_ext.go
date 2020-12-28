package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckPreviewNumberLoader string *VehicleSupervision/internal/modules/vehicle/model.MuckTruckPreviewNumber

func (t MuckTruckPreviewNumber) TableName() string {
	return "MuckTruckPreviewNumber"
}

func (t *MuckTruckPreviewNumber) NewLoader() *MuckTruckPreviewNumberLoader {
	return &MuckTruckPreviewNumberLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckPreviewNumber, []error) {
			var rs []*MuckTruckPreviewNumber
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckPreviewNumber{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
