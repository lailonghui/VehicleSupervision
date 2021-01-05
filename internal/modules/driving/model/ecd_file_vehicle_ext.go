package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileVehicleUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileVehicle

// 数据库表名
func (t EcdFileVehicle) TableName() string {
	return "ecd_file_vehicle"
}

// 主键列名
func (t EcdFileVehicle) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EcdFileVehicle) NewPkLoader() *EcdFileVehiclePkLoader {
	return &EcdFileVehiclePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileVehicle, []error) {
			var rs []*EcdFileVehicle
			db.DB.Model(&EcdFileVehicle{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EcdFileVehicle) UnionPrimaryColumnName() string {
	return "file_vehicle_id"
}

// 新建联合主键dataloader
func (t *EcdFileVehicle) NewUnionPkLoader() *EcdFileVehicleUnionPkLoader {
	return &EcdFileVehicleUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileVehicle, []error) {
			var rs []*EcdFileVehicle
			db.DB.Model(&EcdFileVehicle{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
