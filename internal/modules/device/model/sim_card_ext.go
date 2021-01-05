package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardUnionPkLoader string *VehicleSupervision/internal/modules/device/model.SimCard

// 数据库表名
func (t SimCard) TableName() string {
	return "sim_card"
}

// 主键列名
func (t SimCard) PrimaryColumnName() string {
	return "id"
}

// 新建主键dataloader
func (t *SimCard) NewPkLoader() *SimCardPkLoader {
	return &SimCardPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCard, []error) {
			var rs []*SimCard
			db.DB.Model(&SimCard{}).Where(t.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t SimCard) UnionPrimaryColumnName() string {
	return "sim_card_id"
}

// 新建联合主键dataloader
func (t *SimCard) NewUnionPkLoader() *SimCardUnionPkLoader {
	return &SimCardUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCard, []error) {
			var rs []*SimCard
			db.DB.Model(&SimCard{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
