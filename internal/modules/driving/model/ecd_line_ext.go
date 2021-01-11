package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdLineUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdLine

// 数据库表名
func (t *EcdLine) TableName() string {
	return "ecd_line"
}

// 主键列名
func (t *EcdLine) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EcdLine) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EcdLinePkLoader) NewLoader() *EcdLinePkLoader {
	return &EcdLinePkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdLine, []error) {
			var rs []*EcdLine
			var m EcdLine
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EcdLine) UnionPrimaryColumnName() string {
	return "line_id"
}

// 获取联合主键
func (t *EcdLine) GetUnionPrimary() string {
	return t.LineID
}

// 新建联合主键dataloader
func (t *EcdLineUnionPkLoader) NewLoader() *EcdLineUnionPkLoader {
	return &EcdLineUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdLine, []error) {
			var rs []*EcdLine
			var m EcdLine
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
