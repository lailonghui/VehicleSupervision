package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileVehicleTimeUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileVehicleTime

// 数据库表名
func (t *EcdFileVehicleTime) TableName() string {
	return "ecd_file_vehicle_time"
}

// 主键列名
func (t *EcdFileVehicleTime) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EcdFileVehicleTime) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EcdFileVehicleTimePkLoader) NewLoader() *EcdFileVehicleTimePkLoader {
	return &EcdFileVehicleTimePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileVehicleTime, []error) {
			var rs []*EcdFileVehicleTime
			var m EcdFileVehicleTime
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EcdFileVehicleTime) UnionPrimaryColumnName() string {
	return "vehicle_time_id"
}

// 获取联合主键
func (t *EcdFileVehicleTime) GetUnionPrimary() string {
	return t.VehicleTimeID
}

// 新建联合主键dataloader
func (t *EcdFileVehicleTimeUnionPkLoader) NewLoader() *EcdFileVehicleTimeUnionPkLoader {
	return &EcdFileVehicleTimeUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileVehicleTime, []error) {
			var rs []*EcdFileVehicleTime
			var m EcdFileVehicleTime
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
