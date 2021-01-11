package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileMainSubUnionPkLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileMainSub

// 数据库表名
func (t *EcdFileMainSub) TableName() string {
	return "ecd_file_main_sub"
}

// 主键列名
func (t *EcdFileMainSub) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *EcdFileMainSub) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *EcdFileMainSubPkLoader) NewLoader() *EcdFileMainSubPkLoader {
	return &EcdFileMainSubPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileMainSub, []error) {
			var rs []*EcdFileMainSub
			var m EcdFileMainSub
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *EcdFileMainSub) UnionPrimaryColumnName() string {
	return "file_main_sub_id"
}

// 获取联合主键
func (t *EcdFileMainSub) GetUnionPrimary() string {
	return t.FileMainSubID
}

// 新建联合主键dataloader
func (t *EcdFileMainSubUnionPkLoader) NewLoader() *EcdFileMainSubUnionPkLoader {
	return &EcdFileMainSubUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileMainSub, []error) {
			var rs []*EcdFileMainSub
			var m EcdFileMainSub
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
