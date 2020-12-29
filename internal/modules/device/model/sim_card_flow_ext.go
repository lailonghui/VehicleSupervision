package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardFlowLoader string *VehicleSupervision/internal/modules/device/model.SimCardFlow

// 数据库表名
func (t SimCardFlow) TableName() string {
	return "sim_card_flow"
}

// 主键列名
func (t SimCardFlow) PrimaryColumnName() string {
	return "id"
}

// 联合主键列名
func (t SimCardFlow) UnionPrimaryColumnName() string {
	return "sim_card_flow_id"
}

// 新建dataloader
func (t *SimCardFlow) NewLoader() *SimCardFlowLoader {
	return &SimCardFlowLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardFlow, []error) {
			var rs []*SimCardFlow

			db.DB.Model(&SimCardFlow{}).Where(t.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)

			return rs, nil
		},
	}
}
