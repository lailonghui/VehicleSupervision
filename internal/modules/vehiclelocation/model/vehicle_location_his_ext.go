package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleLocationHisPkLoader string *VehicleSupervision/internal/modules/vehiclelocation/model.VehicleLocationHis

// 数据库表名
func (t VehicleLocationHis) TableName() string {
	return "vehicle_location_his"
}

// 主键列名
func (t VehicleLocationHis) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *VehicleLocationHis) NewPkLoader() *VehicleLocationHisPkLoader {
	return &VehicleLocationHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleLocationHis, []error) {
			var rs []*VehicleLocationHis
			db.DB.Model(&VehicleLocationHis{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
