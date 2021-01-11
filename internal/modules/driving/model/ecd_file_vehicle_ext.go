package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileVehicleUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileVehicle

// 数据库表名
func (t *EcdFileVehicle) TableName() string {
	return "ecd_file_vehicle"
}

// 主键列名
func (t *EcdFileVehicle) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EcdFileVehicle) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EcdFileVehiclePkLoader) NewLoader() *EcdFileVehiclePkLoader {
	return &EcdFileVehiclePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileVehicle, []error) {
			var rs []*EcdFileVehicle
			var m EcdFileVehicle
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EcdFileVehicle) UnionPrimaryColumnName() string {
	return "file_vehicle_id"
}

// 获取联合主键
func (t *EcdFileVehicle) GetUnionPrimary() string {
	return t.FileVehicleID
}

// 新建联合主键dataloader
func (t *EcdFileVehicleUnionPkLoader) NewLoader() *EcdFileVehicleUnionPkLoader {
	return &EcdFileVehicleUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileVehicle, []error) {
			var rs []*EcdFileVehicle
			var m EcdFileVehicle
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
