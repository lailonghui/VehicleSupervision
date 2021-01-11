package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileMainUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileMain

// 数据库表名
func (t *EcdFileMain) TableName() string {
	return "ecd_file_main"
}

// 主键列名
func (t *EcdFileMain) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EcdFileMain) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EcdFileMainPkLoader) NewLoader() *EcdFileMainPkLoader {
	return &EcdFileMainPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileMain, []error) {
			var rs []*EcdFileMain
			var m EcdFileMain
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EcdFileMain) UnionPrimaryColumnName() string {
	return "file_main_id"
}

// 获取联合主键
func (t *EcdFileMain) GetUnionPrimary() string {
	return t.FileMainID
}

// 新建联合主键dataloader
func (t *EcdFileMainUnionPkLoader) NewLoader() *EcdFileMainUnionPkLoader {
	return &EcdFileMainUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileMain, []error) {
			var rs []*EcdFileMain
			var m EcdFileMain
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
