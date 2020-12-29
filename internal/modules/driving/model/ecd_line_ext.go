package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdLineLoader string *VehicleSupervision/internal/modules/driving/model.EcdLine

// 数据库表名
func (t EcdLine) TableName() string {
	return "ecd_line"
}

// 主键列名
func (t EcdLine) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EcdLine) UnionPrimaryColumnName() string {
	return "line_id"
}

// 新建dataloader
func (t *EcdLine) NewLoader() *EcdLineLoader {
	return &EcdLineLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdLine, []error) {
			var rs []*EcdLine

			db.DB.Model(&EcdLine{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
