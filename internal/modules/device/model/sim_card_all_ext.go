package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardAllPkLoader string *VehicleSupervision/internal/modules/device/model.SimCardAll

// 数据库表名
func (t *SimCardAll) TableName() string {
	return "sim_card_all"
}

// 主键列名
func (t *SimCardAll) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *SimCardAll) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *SimCardAllPkLoader) NewLoader() *SimCardAllPkLoader {
	return &SimCardAllPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardAll, []error) {
			var rs []*SimCardAll
			var m SimCardAll
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
