package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileCheckHisUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileCheckHis

// 数据库表名
func (t EcdFileCheckHis) TableName() string {
	return "ecd_file_check_his"
}

// 主键列名
func (t EcdFileCheckHis) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EcdFileCheckHis) NewPkLoader() *EcdFileCheckHisPkLoader {
	return &EcdFileCheckHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileCheckHis, []error) {
			var rs []*EcdFileCheckHis
			db.DB.Model(&EcdFileCheckHis{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EcdFileCheckHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 新建联合主键dataloader
func (t *EcdFileCheckHis) NewUnionPkLoader() *EcdFileCheckHisUnionPkLoader {
	return &EcdFileCheckHisUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileCheckHis, []error) {
			var rs []*EcdFileCheckHis
			db.DB.Model(&EcdFileCheckHis{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
