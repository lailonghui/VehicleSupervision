package model

import (
	"VehicleSupervision/internal/db"
	"time"
)

//go:generate go run github.com/vektah/dataloaden SimCardStatusUnionPkLoader string *VehicleSupervision/internal/modules/device/model.SimCardStatus

// 数据库表名
func (t *SimCardStatus) TableName() string {
	return "sim_card_status"
}

// 主键列名
func (t *SimCardStatus) PrimaryColumnName() string {
	return "id"
}

// 获取主键
func (t *SimCardStatus) GetPrimary() int64 {
	return t.ID
}

// 新建主键dataloader
func (t *SimCardStatusPkLoader) NewLoader() *SimCardStatusPkLoader {
	return &SimCardStatusPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardStatus, []error) {
			var rs []*SimCardStatus
			var m SimCardStatus
			db.DB.Model(&m).Where(m.PrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}

// 联合主键列名
func (t *SimCardStatus) UnionPrimaryColumnName() string {
	return "sim_card_status_id"
}

// 获取联合主键
func (t *SimCardStatus) GetUnionPrimary() string {
	return t.SimCardStatusID
}

// 新建联合主键dataloader
func (t *SimCardStatusUnionPkLoader) NewLoader() *SimCardStatusUnionPkLoader {
	return &SimCardStatusUnionPkLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []string) ([]*SimCardStatus, []error) {
			var rs []*SimCardStatus
			var m SimCardStatus
			db.DB.Model(&m).Where(m.UnionPrimaryColumnName()+" in ?", keys).Find(&rs)
			return rs, nil
		},
	}
}
