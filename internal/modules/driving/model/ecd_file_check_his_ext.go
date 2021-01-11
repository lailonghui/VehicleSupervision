package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileCheckHisUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileCheckHis

// 数据库表名
func (t *EcdFileCheckHis) TableName() string {
	return "ecd_file_check_his"
}

// 主键列名
func (t *EcdFileCheckHis) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EcdFileCheckHis) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EcdFileCheckHisPkLoader) NewLoader() *EcdFileCheckHisPkLoader {
	return &EcdFileCheckHisPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileCheckHis, []error) {
			var rs []*EcdFileCheckHis
			var m EcdFileCheckHis
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EcdFileCheckHis) UnionPrimaryColumnName() string {
	return "his_id"
}

// 获取联合主键
func (t *EcdFileCheckHis) GetUnionPrimary() string {
	return t.HisID
}

// 新建联合主键dataloader
func (t *EcdFileCheckHisUnionPkLoader) NewLoader() *EcdFileCheckHisUnionPkLoader {
	return &EcdFileCheckHisUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileCheckHis, []error) {
			var rs []*EcdFileCheckHis
			var m EcdFileCheckHis
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
