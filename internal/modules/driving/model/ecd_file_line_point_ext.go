package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileLinePointUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileLinePoint

// 数据库表名
func (t *EcdFileLinePoint) TableName() string {
	return "ecd_file_line_point"
}

// 主键列名
func (t *EcdFileLinePoint) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EcdFileLinePoint) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EcdFileLinePointPkLoader) NewLoader() *EcdFileLinePointPkLoader {
	return &EcdFileLinePointPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileLinePoint, []error) {
			var rs []*EcdFileLinePoint
			var m EcdFileLinePoint
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EcdFileLinePoint) UnionPrimaryColumnName() string {
	return "line_point_id"
}

// 获取联合主键
func (t *EcdFileLinePoint) GetUnionPrimary() string {
	return t.LinePointID
}

// 新建联合主键dataloader
func (t *EcdFileLinePointUnionPkLoader) NewLoader() *EcdFileLinePointUnionPkLoader {
	return &EcdFileLinePointUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileLinePoint, []error) {
			var rs []*EcdFileLinePoint
			var m EcdFileLinePoint
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
