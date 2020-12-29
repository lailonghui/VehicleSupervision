package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileVehicleTimeLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileVehicleTime

// 数据库表名
func (t EcdFileVehicleTime) TableName() string {
	return "ecd_file_vehicle_time"
}

// 主键列名
func (t EcdFileVehicleTime) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EcdFileVehicleTime) UnionPrimaryColumnName() string {
	return "vehicle_time_id"
}

// 新建dataloader
func (t *EcdFileVehicleTime) NewLoader() *EcdFileVehicleTimeLoader {
	return &EcdFileVehicleTimeLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileVehicleTime, []error) {
			var rs []*EcdFileVehicleTime

			db.DB.Model(&EcdFileVehicleTime{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
