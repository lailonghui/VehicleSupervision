package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileMainUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileMain

// 数据库表名
func (t EcdFileMain) TableName() string {
	return "ecd_file_main"
}

// 主键列名
func (t EcdFileMain) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *EcdFileMain) NewPkLoader() *EcdFileMainPkLoader {
	return &EcdFileMainPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileMain, []error) {
			var rs []*EcdFileMain
			db.DB.Model(&EcdFileMain{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t EcdFileMain) UnionPrimaryColumnName() string {
	return "file_main_id"
}

// 新建联合主键dataloader
func (t *EcdFileMain) NewUnionPkLoader() *EcdFileMainUnionPkLoader {
	return &EcdFileMainUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileMain, []error) {
			var rs []*EcdFileMain
			db.DB.Model(&EcdFileMain{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
