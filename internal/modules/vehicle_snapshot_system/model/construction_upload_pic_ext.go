package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden ConstructionUploadPicLoader string *VehicleSupervision/internal/modules/vehicle_snapshot_system/model.ConstructionUploadPic

func (t ConstructionUploadPic) TableName() string {
	return "construction_upload_pic"
}

func (t *ConstructionUploadPic) NewLoader() *ConstructionUploadPicLoader {
	return &ConstructionUploadPicLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*ConstructionUploadPic, []error) {
			var rs []*ConstructionUploadPic
			// TODO 按实际需要实现
			db.DB.Model(&ConstructionUploadPic{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
