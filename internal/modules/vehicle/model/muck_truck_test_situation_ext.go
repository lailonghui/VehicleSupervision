package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

// go:generate go run github.com/vektah/dataloaden MuckTruckTestSituationLoader string *VehicleSupervision/internal/modules/vehicle/model.MuckTruckTestSituation

func (t MuckTruckTestSituation) TableName() string {
	return "muck_truck_test_situation"
}

func (t *MuckTruckTestSituation) NewLoader() *MuckTruckTestSituationLoader {
	return &MuckTruckTestSituationLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*MuckTruckTestSituation, []error) {
			var rs []*MuckTruckTestSituation
			// TODO 按实际需要实现
			db.DB.Model(&MuckTruckTestSituation{}).Where("id in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
