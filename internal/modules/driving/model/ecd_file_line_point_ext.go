package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileLinePointLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileLinePoint

// 数据库表名
func (t EcdFileLinePoint) TableName() string {
	return "ecd_file_line_point"
}

// 主键列名
func (t EcdFileLinePoint) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EcdFileLinePoint) UnionPrimaryColumnName() string {
	return "line_point_id"
}

// 新建dataloader
func (t *EcdFileLinePoint) NewLoader() *EcdFileLinePointLoader {
	return &EcdFileLinePointLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileLinePoint, []error) {
			var rs []*EcdFileLinePoint

			db.DB.Model(&EcdFileLinePoint{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
