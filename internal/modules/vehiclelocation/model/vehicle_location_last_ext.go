package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden VehicleLocationLastPkLoader string *VehicleSupervision/internal/modules/vehiclelocation/model.VehicleLocationLast

// 数据库表名
func (t *VehicleLocationLast) TableName() string {
	return "vehicle_location_last"
}

// 主键列名
func (t *VehicleLocationLast) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *VehicleLocationLast) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *VehicleLocationLastPkLoader) NewLoader() *VehicleLocationLastPkLoader {
	return &VehicleLocationLastPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*VehicleLocationLast, []error) {
			var rs []*VehicleLocationLast
			var m VehicleLocationLast
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
