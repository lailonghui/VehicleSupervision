package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileDistrictLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileDistrict

// 数据库表名
func (t EcdFileDistrict) TableName() string {
	return "ecd_file_district"
}

// 主键列名
func (t EcdFileDistrict) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EcdFileDistrict) UnionPrimaryColumnName() string {
	return "file_district_id"
}

// 新建dataloader
func (t *EcdFileDistrict) NewLoader() *EcdFileDistrictLoader {
	return &EcdFileDistrictLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileDistrict, []error) {
			var rs []*EcdFileDistrict

			db.DB.Model(&EcdFileDistrict{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
