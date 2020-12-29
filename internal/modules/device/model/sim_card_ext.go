package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardLoader string *VehicleSupervision/internal/modules/device/model.SimCard

// 数据库表名
func (t SimCard) TableName() string {
	return "sim_card"
}

// 主键列名
func (t SimCard) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t SimCard) UnionPrimaryColumnName() string {
	return "sim_card_id"
}

// 新建dataloader
func (t *SimCard) NewLoader() *SimCardLoader {
	return &SimCardLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCard, []error) {
			var rs []*SimCard

			db.DB.Model(&SimCard{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
