package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden ConstructionCameraLoader string *VehicleSupervision/internal/modules/vehicle_snapshot_system/model.ConstructionCamera

func (t ConstructionCamera) TableName() string {
	return "construction_camera"
}

func (t *ConstructionCamera) NewLoader() *ConstructionCameraLoader {
	return &ConstructionCameraLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ConstructionCamera, []error) {
			var rs []*ConstructionCamera
			// TODO 按实际需要实现
			db.DB.Model(&ConstructionCamera{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
