package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden EcdFileMainSubLoader string *VehicleSupervision/internal/modules/driving/model.EcdFileMainSub

// 数据库表名
func (t EcdFileMainSub) TableName() string {
	return "ecd_file_main_sub"
}

// 主键列名
func (t EcdFileMainSub) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t EcdFileMainSub) UnionPrimaryColumnName() string {
	return "file_main_sub_id"
}

// 新建dataloader
func (t *EcdFileMainSub) NewLoader() *EcdFileMainSubLoader {
	return &EcdFileMainSubLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*EcdFileMainSub, []error) {
			var rs []*EcdFileMainSub

			db.DB.Model(&EcdFileMainSub{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}