package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleLocationHisPkLoader string *VehicleSupervision/internal/modules/vehiclelocation/model.VehicleLocationHis

// 数据库表名
func (t *VehicleLocationHis) TableName() string {
	return "vehicle_location_his"
}

// 主键列名
func (t *VehicleLocationHis) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *VehicleLocationHis) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *VehicleLocationHisPkLoader) NewLoader() *VehicleLocationHisPkLoader {
	return &VehicleLocationHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleLocationHis, []error) {
			var rs []*VehicleLocationHis
			var m VehicleLocationHis
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
